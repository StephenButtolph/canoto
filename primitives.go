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
	// SizeFint32 is the size of a 32-bit fixed size integer in bytes.
	SizeFint32 = 4
	// SizeFint64 is the size of a 64-bit fixed size integer in bytes.
	SizeFint64 = 8
	// SizeBool is the size of a boolean in bytes.
	SizeBool = 1

	falseByte        = 0
	trueByte         = 1
	continuationMask = 0x80
)

var (
	ErrInvalidFieldOrder  = errors.New("invalid field order")
	ErrUnexpectedWireType = errors.New("unexpected wire type")
	ErrDuplicateOneOf     = errors.New("duplicate oneof field")
	ErrInvalidLength      = errors.New("decoded length is invalid")
	ErrZeroValue          = errors.New("zero value")
	ErrUnknownField       = errors.New("unknown field")
	ErrPaddedZeroes       = errors.New("padded zeroes")

	ErrOverflow        = errors.New("overflow")
	ErrInvalidWireType = errors.New("invalid wire type")
	ErrInvalidBool     = errors.New("decoded bool is neither true nor false")
	ErrStringNotUTF8   = errors.New("decoded string is not UTF-8")
)

type (
	Sint interface {
		~int8 | ~int16 | ~int32 | ~int64
	}
	Uint interface {
		~uint8 | ~uint16 | ~uint32 | ~uint64
	}
	Int   interface{ Sint | Uint }
	Int32 interface{ ~int32 | ~uint32 }
	Int64 interface{ ~int64 | ~uint64 }
	Bytes interface{ ~string | ~[]byte }

	// Reader contains all the state needed to unmarshal a Canoto type.
	//
	// The functions in this package are not methods on the Reader type to
	// enable the usage of generics.
	Reader struct {
		B      []byte
		Unsafe bool
	}

	// Writer contains all the state needed to marshal a Canoto type.
	//
	// The functions in this package are not methods on the Writer type to
	// enable the usage of generics.
	Writer struct {
		B []byte
	}
)

// HasNext returns true if there are more bytes to read.
func HasNext(r *Reader) bool {
	return len(r.B) > 0
}

// Append writes unprefixed bytes to the writer.
func Append[T Bytes](w *Writer, v T) {
	w.B = append(w.B, v...)
}

// Tag calculates the tag for a field number and wire type.
//
// This function should not typically be used during marshaling, as tags can be
// precomputed.
func Tag(fieldNumber uint32, wireType WireType) []byte {
	w := Writer{}
	AppendInt(&w, fieldNumber<<wireTypeLength|uint32(wireType))
	return w.B
}

// ReadTag reads the next field number and wire type from the reader.
func ReadTag(r *Reader) (uint32, WireType, error) {
	var val uint32
	if err := ReadInt(r, &val); err != nil {
		return 0, 0, err
	}

	wireType := WireType(val & wireTypeMask)
	if !wireType.IsValid() {
		return 0, 0, ErrInvalidWireType
	}

	return val >> wireTypeLength, wireType, nil
}

// SizeInt calculates the size of an integer when encoded as a varint.
func SizeInt[T Int](v T) int {
	if v == 0 {
		return 1
	}
	return (bits.Len64(uint64(v)) + 6) / 7
}

// CountInts counts the number of varints that are encoded in bytes.
func CountInts(bytes []byte) int {
	var count int
	for _, b := range bytes {
		if b < continuationMask {
			count++
		}
	}
	return count
}

// ReadInt reads a varint encoded integer from the reader.
func ReadInt[T Int](r *Reader, v *T) error {
	val, bytesRead := binary.Uvarint(r.B)
	switch {
	case bytesRead == 0:
		return io.ErrUnexpectedEOF
	case bytesRead < 0 || uint64(T(val)) != val:
		return ErrOverflow
	// To ensure decoding is canonical, we check for padded zeroes in the
	// varint.
	// The last byte of the varint includes the most significant bits.
	// If the last byte is 0, then the number should have been encoded more
	// efficiently by removing this zero.
	case bytesRead > 1 && r.B[bytesRead-1] == 0x00:
		return ErrPaddedZeroes
	default:
		r.B = r.B[bytesRead:]
		*v = T(val)
		return nil
	}
}

// AppendInt writes an integer to the writer as a varint.
func AppendInt[T Int](w *Writer, v T) {
	w.B = binary.AppendUvarint(w.B, uint64(v))
}

// SizeSint calculates the size of an integer when zigzag encoded as a varint.
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

// ReadSint reads a zigzag encoded integer from the reader.
func ReadSint[T Sint](r *Reader, v *T) error {
	var largeVal uint64
	if err := ReadInt(r, &largeVal); err != nil {
		return err
	}

	uVal := largeVal >> 1
	val := T(uVal)
	// If T is an int32, it's possible that some bits were truncated during the
	// cast. In this case, casting back to uint64 would result in a different
	// value.
	if uint64(val) != uVal {
		return ErrOverflow
	}

	if largeVal&1 != 0 {
		val = ^val
	}
	*v = val
	return nil
}

// AppendSint writes an integer to the writer as a zigzag encoded varint.
func AppendSint[T Sint](w *Writer, v T) {
	if v >= 0 {
		w.B = binary.AppendUvarint(w.B, uint64(v)<<1)
	} else {
		w.B = binary.AppendUvarint(w.B, ^uint64(v)<<1|1)
	}
}

// ReadFint32 reads a 32-bit fixed size integer from the reader.
func ReadFint32[T Int32](r *Reader, v *T) error {
	if len(r.B) < SizeFint32 {
		return io.ErrUnexpectedEOF
	}

	*v = T(binary.LittleEndian.Uint32(r.B))
	r.B = r.B[SizeFint32:]
	return nil
}

// AppendFint32 writes a 32-bit fixed size integer to the writer.
func AppendFint32[T Int32](w *Writer, v T) {
	w.B = binary.LittleEndian.AppendUint32(w.B, uint32(v))
}

// ReadFint64 reads a 64-bit fixed size integer from the reader.
func ReadFint64[T Int64](r *Reader, v *T) error {
	if len(r.B) < SizeFint64 {
		return io.ErrUnexpectedEOF
	}

	*v = T(binary.LittleEndian.Uint64(r.B))
	r.B = r.B[SizeFint64:]
	return nil
}

// AppendFint64 writes a 64-bit fixed size integer to the writer.
func AppendFint64[T Int64](w *Writer, v T) {
	w.B = binary.LittleEndian.AppendUint64(w.B, uint64(v))
}

// ReadBool reads a boolean from the reader.
func ReadBool[T ~bool](r *Reader, v *T) error {
	switch {
	case len(r.B) < SizeBool:
		return io.ErrUnexpectedEOF
	case r.B[0] > trueByte:
		return ErrInvalidBool
	default:
		*v = r.B[0] == trueByte
		r.B = r.B[SizeBool:]
		return nil
	}
}

// AppendBool writes a boolean to the writer.
func AppendBool[T ~bool](w *Writer, b T) {
	if b {
		w.B = append(w.B, trueByte)
	} else {
		w.B = append(w.B, falseByte)
	}
}

// SizeBytes calculates the size the length-prefixed bytes would take if
// written.
func SizeBytes[T Bytes](v T) int {
	return SizeInt(int64(len(v))) + len(v)
}

// CountBytes counts the consecutive number of length-prefixed fields with the
// given tag.
func CountBytes(bytes []byte, tag string) (int, error) {
	var (
		r     = Reader{B: bytes}
		count = 0
	)
	for HasPrefix(r.B, tag) {
		r.B = r.B[len(tag):]
		var length int64
		if err := ReadInt(&r, &length); err != nil {
			return 0, err
		}
		if length < 0 {
			return 0, ErrInvalidLength
		}
		if length > int64(len(r.B)) {
			return 0, io.ErrUnexpectedEOF
		}
		r.B = r.B[length:]
		count++
	}
	return count, nil
}

// HasPrefix returns true if the bytes start with the given prefix.
func HasPrefix(bytes []byte, prefix string) bool {
	return len(bytes) >= len(prefix) && string(bytes[:len(prefix)]) == prefix
}

// ReadString reads a string from the reader. The string is verified to be valid
// UTF-8.
func ReadString[T ~string](r *Reader, v *T) error {
	var length int64
	if err := ReadInt[int64](r, &length); err != nil {
		return err
	}
	if length < 0 {
		return ErrInvalidLength
	}
	if length > int64(len(r.B)) {
		return io.ErrUnexpectedEOF
	}

	bytes := r.B[:length]
	if !utf8.Valid(bytes) {
		return ErrStringNotUTF8
	}

	r.B = r.B[length:]
	if r.Unsafe {
		*v = T(unsafeString(bytes))
	} else {
		*v = T(bytes)
	}
	return nil
}

// ReadBytes reads a byte slice from the reader.
func ReadBytes[T ~[]byte](r *Reader, v *T) error {
	var length int64
	if err := ReadInt[int64](r, &length); err != nil {
		return err
	}
	if length < 0 {
		return ErrInvalidLength
	}
	if length > int64(len(r.B)) {
		return io.ErrUnexpectedEOF
	}

	bytes := r.B[:length]
	r.B = r.B[length:]
	if !r.Unsafe {
		bytes = slices.Clone(bytes)
	}
	*v = T(bytes)
	return nil
}

// AppendBytes writes a length-prefixed byte slice to the writer.
func AppendBytes[T Bytes](w *Writer, v T) {
	AppendInt(w, int64(len(v)))
	w.B = append(w.B, v...)
}

// MakeSlice creates a new slice with the given length. It is equivalent to
// `make([]T, length)`.
//
// This function is useful to use in auto-generated code, when the type of a
// variable is unknown. For example, if we have a variable `v` which we know to
// be a slice, but we do not know the type of the elements, we can use this
// function to leverage golang's type inference to create the new slice.
func MakeSlice[T any](_ []T, length int) []T {
	return make([]T, length)
}

// IsZero returns true if the value is the zero value for its type.
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}

// unsafeString converts a []byte to an unsafe string.
//
// Invariant: The input []byte must not be modified.
func unsafeString(b []byte) string {
	// avoid copying during the conversion
	return unsafe.String(unsafe.SliceData(b), len(b))
}
