// Code generated by Canoto. DO NOT EDIT.

package examples

import (
	"io"
	"sync/atomic"
	"unicode/utf8"

	"github.com/StephenButtolph/canoto"
)

// Ensure that unused imports do not error
var (
	_ = io.ErrUnexpectedEOF
	_ = utf8.ValidString
)

const (
	canoto__LargestFieldNumber__Int32__tag = "\xf8\xff\xff\xff\x0f" // canoto.Tag(536870911, canoto.Varint)
)

// Ensure that the generated methods correctly implement the interface
var _ canoto.Message = (*LargestFieldNumber)(nil)

type canotoData_LargestFieldNumber struct {
	size atomic.Int64
}

// UnmarshalCanoto unmarshals a Canoto-encoded byte slice into the struct.
//
// The struct is not cleared before unmarshaling, any fields not present in the
// bytes will retain their previous values.
func (c *LargestFieldNumber) UnmarshalCanoto(bytes []byte) error {
	r := canoto.Reader{
		B: bytes,
	}
	return c.UnmarshalCanotoFrom(&r)
}

// UnmarshalCanotoFrom populates the struct from a canoto.Reader. Most users
// should just use UnmarshalCanoto.
//
// The struct is not cleared before unmarshaling, any fields not present in the
// bytes will retain their previous values.
//
// This function enables configuration of reader options.
func (c *LargestFieldNumber) UnmarshalCanotoFrom(r *canoto.Reader) error {
	var minField uint32
	for canoto.HasNext(r) {
		field, wireType, err := canoto.ReadTag(r)
		if err != nil {
			return err
		}
		if field < minField {
			return canoto.ErrInvalidFieldOrder
		}

		switch field {
		case 536870911:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if err := canoto.ReadInt(r, &c.Int32); err != nil {
				return err
			}
			if canoto.IsZero(c.Int32) {
				return canoto.ErrZeroValue
			}
		default:
			return canoto.ErrUnknownField
		}

		minField = field + 1
	}
	return nil
}

// ValidCanoto validates that the struct can be correctly marshaled into the
// Canoto format.
//
// Specifically, ValidCanoto ensures that all strings are valid utf-8 and all
// custom types are ValidCanoto.
func (c *LargestFieldNumber) ValidCanoto() bool {
	return true
}

// CalculateCanotoSize calculates the size of the Canoto representation and
// caches it.
func (c *LargestFieldNumber) CalculateCanotoSize() int {
	var size int
	if !canoto.IsZero(c.Int32) {
		size += len(canoto__LargestFieldNumber__Int32__tag) + canoto.SizeInt(c.Int32)
	}
	c.canotoData.size.Store(int64(size))
	return size
}

// CachedCanotoSize returns the previously calculated size of the Canoto
// representation from CalculateCanotoSize.
//
// If CalculateCanotoSize has not yet been called, it will return 0.
//
// If the struct has been modified since the last call to CalculateCanotoSize,
// the returned size may be incorrect.
func (c *LargestFieldNumber) CachedCanotoSize() int {
	return int(c.canotoData.size.Load())
}

// MarshalCanoto returns the Canoto representation of this struct.
//
// It is assumed that this struct is ValidCanoto.
func (c *LargestFieldNumber) MarshalCanoto() []byte {
	w := canoto.Writer{
		B: make([]byte, 0, c.CalculateCanotoSize()),
	}
	c.MarshalCanotoInto(&w)
	return w.B
}

// MarshalCanotoInto writes the struct into a canoto.Writer. Most users should
// just use MarshalCanoto.
//
// It is assumed that CalculateCanotoSize has been called since the last
// modification to this struct.
//
// It is assumed that this struct is ValidCanoto.
func (c *LargestFieldNumber) MarshalCanotoInto(w *canoto.Writer) {
	if !canoto.IsZero(c.Int32) {
		canoto.Append(w, canoto__LargestFieldNumber__Int32__tag)
		canoto.AppendInt(w, c.Int32)
	}
}
