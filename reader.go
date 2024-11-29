package canoto

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"slices"
	"unsafe"
)

const (
	falseByte = 0
	trueByte  = 1
)

var (
	errOverflow      = errors.New("overflow")
	errPaddedZeroes  = errors.New("varint has padded zeroes")
	errInvalidBool   = errors.New("decoded bool is neither true nor false")
	errInvalidLength = errors.New("decoded length is invalid")
)

type Reader struct {
	b      []byte
	unsafe bool
}

func (r *Reader) Tag() (uint32, WireType, error) {
	val, err := r.Uint32()
	if err != nil {
		return 0, 0, err
	}

	wireType := WireType(val & wireTypeMask)
	if !wireType.IsValid() {
		return 0, 0, errInvalidWireType
	}

	return val >> wireTypeLength, wireType, err
}

func (r *Reader) Int32() (int32, error) {
	uVal64, err := r.Uint64()
	if err != nil {
		return 0, err
	}
	val64 := int64(uVal64)
	if val64 < math.MinInt32 || val64 > math.MaxInt32 {
		return 0, errOverflow
	}
	return int32(val64), nil
}

func (r *Reader) Int64() (int64, error) {
	uVal, err := r.Uint64()
	return int64(uVal), err
}

func (r *Reader) Uint32() (uint32, error) {
	val, err := r.Uint64()
	if err != nil {
		return 0, err
	}
	if val > math.MaxUint32 {
		return 0, errOverflow
	}
	return uint32(val), nil
}

func (r *Reader) Uint64() (uint64, error) {
	val, bytesRead := binary.Uvarint(r.b)
	switch {
	case bytesRead == 0:
		return 0, io.ErrUnexpectedEOF
	case bytesRead < 0:
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
		return val, nil
	}
}

func (r *Reader) Sint32() (int32, error) {
	uVal, err := r.Uint32()
	val := int32(uVal >> 1)
	if uVal&1 != 0 {
		val = ^val
	}
	return val, err
}

func (r *Reader) Sint64() (int64, error) {
	uVal, err := r.Uint64()
	val := int64(uVal >> 1)
	if uVal&1 != 0 {
		val = ^val
	}
	return val, err
}

func (r *Reader) Fixed32() (uint32, error) {
	if len(r.b) < SizeFixed32 {
		return 0, io.ErrUnexpectedEOF
	}

	val := binary.LittleEndian.Uint32(r.b)
	r.b = r.b[SizeFixed32:]
	return val, nil
}

func (r *Reader) Fixed64() (uint64, error) {
	if len(r.b) < SizeFixed64 {
		return 0, io.ErrUnexpectedEOF
	}

	val := binary.LittleEndian.Uint64(r.b)
	r.b = r.b[SizeFixed64:]
	return val, nil
}

func (r *Reader) Sfixed32() (int32, error) {
	val, err := r.Fixed32()
	return int32(val), err
}

func (r *Reader) Sfixed64() (int64, error) {
	val, err := r.Fixed64()
	return int64(val), err
}

func (r *Reader) Bool() (bool, error) {
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

func (r *Reader) String() (string, error) {
	length, err := r.Int32()
	if err != nil {
		return "", err
	}
	if length < 0 {
		return "", errInvalidLength
	}
	if length > int32(len(r.b)) {
		return "", io.ErrUnexpectedEOF
	}

	bytes := r.b[:length]
	r.b = r.b[length:]
	if r.unsafe {
		return unsafeString(bytes), nil
	}
	return string(bytes), nil
}

func (r *Reader) Bytes() ([]byte, error) {
	length, err := r.Int32()
	if err != nil {
		return nil, err
	}
	if length < 0 {
		return nil, errInvalidLength
	}
	if length > int32(len(r.b)) {
		return nil, io.ErrUnexpectedEOF
	}

	bytes := r.b[:length]
	r.b = r.b[length:]
	if r.unsafe {
		return bytes, nil
	}
	return slices.Clone(bytes), nil
}

// unsafeString converts a []byte to an unsafe string.
//
// Invariant: The input []byte must not be modified.
func unsafeString(b []byte) string {
	// avoid copying during the conversion
	return unsafe.String(unsafe.SliceData(b), len(b))
}
