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

	falseByte = 0
	trueByte  = 1
)

var (
	ErrInvalidFieldOrder = errors.New("invalid field order")
	ErrUnknownField      = errors.New("unknown field")

	errOverflow      = errors.New("overflow")
	errPaddedZeroes  = errors.New("varint has padded zeroes")
	errInvalidBool   = errors.New("decoded bool is neither true nor false")
	errZeroValue     = errors.New("zero value")
	errInvalidLength = errors.New("decoded length is invalid")
	errStringNotUTF8 = errors.New("decoded string is not UTF-8")
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

func Append(w *Writer, b []byte) {
	w.B = append(w.B, b...)
}

func Tag(fieldNumber uint32, wireType WireType) []byte {
	w := Writer{}
	AppendInt(&w, uint64(fieldNumber<<wireTypeLength|uint32(wireType)))
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
	return (bits.Len64(uint64(v)) + 6) / 7
}

func ReadInt[T Int](r *Reader) (T, error) {
	val, bytesRead := binary.Uvarint(r.B)
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
	case r.B[bytesRead-1] == 0x00:
		return 0, errPaddedZeroes
	default:
		r.B = r.B[bytesRead:]
		return T(val), nil
	}
}

func AppendInt[T Int](w *Writer, v T) {
	w.B = binary.AppendUvarint(w.B, uint64(v))
}

func SizeSint[T Sint](v T) int {
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
	if val == 0 {
		return 0, errZeroValue
	}

	r.B = r.B[SizeFint32:]
	return T(val), nil
}

func AppendFint32[T Int32](w *Writer, v T) {
	var bytes [SizeFint32]byte
	binary.LittleEndian.PutUint32(bytes[:], uint32(v))
	w.B = append(w.B, bytes[:]...)
}

func ReadFint64[T Int64](r *Reader) (T, error) {
	if len(r.B) < SizeFint64 {
		return 0, io.ErrUnexpectedEOF
	}

	val := binary.LittleEndian.Uint64(r.B)
	if val == 0 {
		return 0, errZeroValue
	}

	r.B = r.B[SizeFint64:]
	return T(val), nil
}

func AppendFint64[T Int64](w *Writer, v T) {
	var bytes [SizeFint64]byte
	binary.LittleEndian.PutUint64(bytes[:], uint64(v))
	w.B = append(w.B, bytes[:]...)
}

func ReadTrue(r *Reader) error {
	switch {
	case len(r.B) < SizeBool:
		return io.ErrUnexpectedEOF
	case r.B[0] != trueByte:
		return errInvalidBool
	default:
		r.B = r.B[SizeBool:]
		return nil
	}
}

func AppendTrue(w *Writer) {
	w.B = append(w.B, trueByte)
}

func SizeBytes[T Bytes](v T) int {
	return SizeInt(int64(len(v))) + len(v)
}

func ReadString(r *Reader) (string, error) {
	length, err := ReadInt[int32](r)
	if err != nil {
		return "", err
	}
	if length < 0 {
		return "", errInvalidLength
	}
	if length > int32(len(r.B)) {
		return "", io.ErrUnexpectedEOF
	}

	bytes := r.B[:length]
	if !utf8.Valid(bytes) {
		return "", errStringNotUTF8
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
		return nil, errInvalidLength
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
