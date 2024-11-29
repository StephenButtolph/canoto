package canoto

import (
	"encoding/binary"
	"errors"
	"io"
	"math/bits"
	"slices"
)

const (
	SizeFint32 = 4
	SizeFint64 = 8
	SizeBool   = 1

	falseByte = 0
	trueByte  = 1
)

var (
	errOverflow      = errors.New("overflow")
	errPaddedZeroes  = errors.New("varint has padded zeroes")
	errInvalidBool   = errors.New("decoded bool is neither true nor false")
	errInvalidLength = errors.New("decoded length is invalid")
)

type (
	Sint  interface{ ~int32 | ~int64 }
	Uint  interface{ ~uint32 | ~uint64 }
	Int   interface{ Sint | Uint }
	Int32 interface{ ~int32 | uint32 }
	Int64 interface{ ~int64 | uint64 }
	Bytes interface{ ~string | ~[]byte }

	Reader struct {
		b      []byte
		unsafe bool
	}
	Writer struct {
		b []byte
	}
)

func SizeTag(fieldNumber uint32, wireType WireType) int {
	return SizeInt(tagToUint64(fieldNumber, wireType))
}

func ReadTag(r *Reader) (uint32, WireType, error) {
	val, err := ReadInt[uint32](r)
	if err != nil {
		return 0, 0, err
	}

	wireType := WireType(val & wireTypeMask)
	if !wireType.IsValid() {
		return 0, 0, errInvalidWireType
	}

	return val >> wireTypeLength, wireType, err
}

func AppendTag(w *Writer, fieldNumber uint32, wireType WireType) {
	AppendInt(w, tagToUint64(fieldNumber, wireType))
}

func SizeInt[T Int](v T) int {
	if v == 0 {
		return 1
	}
	return (bits.Len64(uint64(v)) + 6) / 7
}

func ReadInt[T Int](r *Reader) (T, error) {
	val, bytesRead := binary.Uvarint(r.b)
	switch {
	case bytesRead == 0:
		return 0, io.ErrUnexpectedEOF
	case bytesRead < 0 || uint64(T(val)) != val:
		return 0, errOverflow
	// To ensure decoding is canonical, we check for padded zeroes in the
	// varint.
	// The last byte of the varint includes the most significant bits.
	// If the last byte is 0, then the number should have been encoded more
	// efficiently by removing this zero.
	case bytesRead > 1 && r.b[bytesRead-1] == 0x00:
		return 0, errPaddedZeroes
	default:
		r.b = r.b[bytesRead:]
		return T(val), nil
	}
}

func AppendInt[T Int](w *Writer, v T) {
	w.b = binary.AppendUvarint(w.b, uint64(v))
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
	if uint64(T(uVal)) != uVal {
		return 0, errOverflow
	}

	val := T(uVal)
	if largeVal&1 != 0 {
		val = ^val
	}
	return val, err
}

func AppendSint[T Sint](w *Writer, v T) {
	if v >= 0 {
		w.b = binary.AppendUvarint(w.b, uint64(v)<<1)
	} else {
		w.b = binary.AppendUvarint(w.b, ^uint64(v)<<1|1)
	}
}

func ReadFint32[T Int32](r *Reader) (T, error) {
	if len(r.b) < SizeFint32 {
		return 0, io.ErrUnexpectedEOF
	}

	val := binary.LittleEndian.Uint32(r.b)
	r.b = r.b[SizeFint32:]
	return T(val), nil
}

func AppendFint32[T Int32](w *Writer, v T) {
	var bytes [SizeFint32]byte
	binary.LittleEndian.PutUint32(bytes[:], uint32(v))
	w.b = append(w.b, bytes[:]...)
}

func ReadFint64[T Int64](r *Reader) (T, error) {
	if len(r.b) < SizeFint64 {
		return 0, io.ErrUnexpectedEOF
	}

	val := binary.LittleEndian.Uint64(r.b)
	r.b = r.b[SizeFint64:]
	return T(val), nil
}

func AppendFint64[T Int64](w *Writer, v T) {
	var bytes [SizeFint64]byte
	binary.LittleEndian.PutUint64(bytes[:], uint64(v))
	w.b = append(w.b, bytes[:]...)
}

func ReadBool(r *Reader) (bool, error) {
	if len(r.b) < SizeBool {
		return false, io.ErrUnexpectedEOF
	}

	boolByte := r.b[0]
	if boolByte > trueByte {
		return false, errInvalidBool
	}

	r.b = r.b[SizeBool:]
	return boolByte == trueByte, nil
}

func AppendBool(w *Writer, v bool) {
	if v {
		w.b = append(w.b, trueByte)
	} else {
		w.b = append(w.b, falseByte)
	}
}

func SizeBytes[T Bytes](v T) int {
	return SizeInt(int64(len(v))) + len(v)
}

func ReadBytes[T Bytes](r *Reader) (val T, err error) {
	length, err := ReadInt[int32](r)
	if err != nil {
		return val, err
	}
	if length < 0 {
		return val, errInvalidLength
	}
	if length > int32(len(r.b)) {
		return val, io.ErrUnexpectedEOF
	}

	bytes := r.b[:length]
	r.b = r.b[length:]
	if r.unsafe {
		return T(bytes), nil
	}
	return T(slices.Clone(bytes)), nil
}

func AppendBytes[T Bytes](w *Writer, v T) {
	AppendInt(w, int64(len(v)))
	w.b = append(w.b, v...)
}

func tagToUint64(fieldNumber uint32, wireType WireType) uint64 {
	return uint64(fieldNumber<<wireTypeLength | uint32(wireType))
}
