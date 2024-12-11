package canoto

import (
	"encoding/binary"
	"errors"
	"io"
	"math/bits"
	"slices"
	"unicode/utf8"
	"unsafe"
)

const (
	SizeFint32 = 4
	SizeFint64 = 8
	SizeBool   = 1

	falseByte        = 0
	trueByte         = 1
	continuationMask = 0x80
)

var (
	ErrInvalidFieldOrder = errors.New("invalid field order")
	ErrZeroValue         = errors.New("zero value")
	ErrUnknownField      = errors.New("unknown field")

	ErrOverflow      = errors.New("overflow")
	ErrPaddedZeroes  = errors.New("varint has padded zeroes")
	ErrInvalidBool   = errors.New("decoded bool is neither true nor false")
	ErrInvalidLength = errors.New("decoded length is invalid")
	ErrStringNotUTF8 = errors.New("decoded string is not UTF-8")
)

type (
	Sint  interface{ ~int32 | ~int64 }
	Uint  interface{ ~uint32 | ~uint64 }
	Int   interface{ Sint | Uint }
	Int32 interface{ ~int32 | uint32 }
	Int64 interface{ ~int64 | uint64 }
	Bytes interface{ ~string | ~[]byte }

	Reader struct {
		B      []byte
		Unsafe bool
	}
	Writer struct {
		B []byte
	}
)

func HasNext(r *Reader) bool {
	return len(r.B) > 0
}

func Append[T Bytes](w *Writer, v T) {
	w.B = append(w.B, v...)
}

func Tag(fieldNumber uint32, wireType WireType) []byte {
	w := Writer{}
	AppendInt(&w, fieldNumber<<wireTypeLength|uint32(wireType))
	return w.B
}

func ReadTag(r *Reader) (uint32, WireType, error) {
	val, err := ReadInt[uint32](r)
	if err != nil {
		return 0, 0, err
	}

	wireType := WireType(val & wireTypeMask)
	if !wireType.IsValid() {
		return 0, 0, ErrInvalidWireType
	}

	return val >> wireTypeLength, wireType, err
}

func SizeInt[T Int](v T) int {
	if v == 0 {
		return 1
	}
	return (bits.Len64(uint64(v)) + 6) / 7
}

func CountInts(bytes []byte) int {
	var count int
	for _, b := range bytes {
		if b < continuationMask {
			count++
		}
	}
	return count
}

func ReadInt[T Int](r *Reader) (T, error) {
	val, bytesRead := binary.Uvarint(r.B)
	switch {
	case bytesRead == 0:
		return 0, io.ErrUnexpectedEOF
	case bytesRead < 0 || uint64(T(val)) != val:
		return 0, ErrOverflow
	// To ensure decoding is canonical, we check for padded zeroes in the
	// varint.
	// The last byte of the varint includes the most significant bits.
	// If the last byte is 0, then the number should have been encoded more
	// efficiently by removing this zero.
	case bytesRead > 1 && r.B[bytesRead-1] == 0x00:
		return 0, ErrPaddedZeroes
	default:
		r.B = r.B[bytesRead:]
		return T(val), nil
	}
}

func AppendInt[T Int](w *Writer, v T) {
	w.B = binary.AppendUvarint(w.B, uint64(v))
}

func SizeSint[T Sint](v T) int {
	if v == 0 {
		return 1
	}

	var uv uint64
	if v > 0 {
		uv = uint64(v) << 1
	} else {
		uv = ^uint64(v)<<1 | 1
	}
	return (bits.Len64(uv) + 6) / 7
}

func ReadSint[T Sint](r *Reader) (T, error) {
	largeVal, err := ReadInt[uint64](r)
	if err != nil {
		return 0, err
	}

	uVal := largeVal >> 1
	val := T(uVal)
	// If T is an int32, it's possible that some bits were truncated during the
	// cast. In this case, casting back to uint64 would result in a different
	// value.
	if uint64(val) != uVal {
		return 0, ErrOverflow
	}

	if largeVal&1 != 0 {
		val = ^val
	}
	return val, err
}

func AppendSint[T Sint](w *Writer, v T) {
	if v >= 0 {
		w.B = binary.AppendUvarint(w.B, uint64(v)<<1)
	} else {
		w.B = binary.AppendUvarint(w.B, ^uint64(v)<<1|1)
	}
}

func ReadFint32[T Int32](r *Reader) (T, error) {
	if len(r.B) < SizeFint32 {
		return 0, io.ErrUnexpectedEOF
	}

	val := binary.LittleEndian.Uint32(r.B)
	r.B = r.B[SizeFint32:]
	return T(val), nil
}

func AppendFint32[T Int32](w *Writer, v T) {
	w.B = binary.LittleEndian.AppendUint32(w.B, uint32(v))
}

func ReadFint64[T Int64](r *Reader) (T, error) {
	if len(r.B) < SizeFint64 {
		return 0, io.ErrUnexpectedEOF
	}

	val := binary.LittleEndian.Uint64(r.B)
	r.B = r.B[SizeFint64:]
	return T(val), nil
}

func AppendFint64[T Int64](w *Writer, v T) {
	w.B = binary.LittleEndian.AppendUint64(w.B, uint64(v))
}

func ReadBool(r *Reader) (bool, error) {
	switch {
	case len(r.B) < SizeBool:
		return false, io.ErrUnexpectedEOF
	case r.B[0] > trueByte:
		return false, ErrInvalidBool
	default:
		isTrue := r.B[0] == trueByte
		r.B = r.B[SizeBool:]
		return isTrue, nil
	}
}

func AppendBool(w *Writer, b bool) {
	if b {
		w.B = append(w.B, trueByte)
	} else {
		w.B = append(w.B, falseByte)
	}
}

func SizeBytes[T Bytes](v T) int {
	return SizeInt(int64(len(v))) + len(v)
}

func CountBytes(bytes []byte, tag string) (int, error) {
	var (
		tagLen = len(tag)
		r      = Reader{B: bytes}
		count  = 0
	)
	for len(r.B) >= len(tag) && string(r.B[:tagLen]) == tag {
		r.B = r.B[tagLen:]
		length, err := ReadInt[int32](&r)
		if err != nil {
			return 0, err
		}
		if length < 0 {
			return 0, ErrInvalidLength
		}
		if length > int32(len(r.B)) {
			return 0, io.ErrUnexpectedEOF
		}
		r.B = r.B[length:]
		count++
	}
	return count, nil
}

func ReadString(r *Reader) (string, error) {
	length, err := ReadInt[int32](r)
	if err != nil {
		return "", err
	}
	if length < 0 {
		return "", ErrInvalidLength
	}
	if length > int32(len(r.B)) {
		return "", io.ErrUnexpectedEOF
	}

	bytes := r.B[:length]
	if !utf8.Valid(bytes) {
		return "", ErrStringNotUTF8
	}

	r.B = r.B[length:]
	if r.Unsafe {
		return unsafeString(bytes), nil
	}
	return string(bytes), nil
}

func ReadBytes(r *Reader) ([]byte, error) {
	length, err := ReadInt[int32](r)
	if err != nil {
		return nil, err
	}
	if length < 0 {
		return nil, ErrInvalidLength
	}
	if length > int32(len(r.B)) {
		return nil, io.ErrUnexpectedEOF
	}

	bytes := r.B[:length]
	r.B = r.B[length:]
	if r.Unsafe {
		return bytes, nil
	}
	return slices.Clone(bytes), nil
}

func AppendBytes[T Bytes](w *Writer, v T) {
	AppendInt(w, int64(len(v)))
	w.B = append(w.B, v...)
}

// unsafeString converts a []byte to an unsafe string.
//
// Invariant: The input []byte must not be modified.
func unsafeString(b []byte) string {
	// avoid copying during the conversion
	return unsafe.String(unsafe.SliceData(b), len(b))
}
