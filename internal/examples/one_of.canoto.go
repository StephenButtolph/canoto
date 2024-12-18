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
	canoto__OneOf__A1__tag = "\x08" // canoto.Tag(1, canoto.Varint)
	canoto__OneOf__B1__tag = "\x18" // canoto.Tag(3, canoto.Varint)
	canoto__OneOf__B2__tag = "\x20" // canoto.Tag(4, canoto.Varint)
	canoto__OneOf__C__tag  = "\x28" // canoto.Tag(5, canoto.Varint)
	canoto__OneOf__D__tag  = "\x30" // canoto.Tag(6, canoto.Varint)
	canoto__OneOf__A2__tag = "\x38" // canoto.Tag(7, canoto.Varint)
)

type canotoData_OneOf struct {
	// Enforce noCopy before atomic usage.
	// See https://github.com/StephenButtolph/canoto/pull/32
	_ atomic.Int64

	size int

	aOneOf uint32
	bOneOf uint32
}

// UnmarshalCanoto unmarshals a Canoto-encoded byte slice into the struct.
//
// The struct is not cleared before unmarshaling, any fields not present in the
// bytes will retain their previous values.
func (c *OneOf) UnmarshalCanoto(bytes []byte) error {
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
func (c *OneOf) UnmarshalCanotoFrom(r *canoto.Reader) error {
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
		case 1:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if err := canoto.ReadInt(r, &c.A1); err != nil {
				return err
			}
			if canoto.IsZero(c.A1) {
				return canoto.ErrZeroValue
			}
		case 3:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if err := canoto.ReadInt(r, &c.B1); err != nil {
				return err
			}
			if canoto.IsZero(c.B1) {
				return canoto.ErrZeroValue
			}
		case 4:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if err := canoto.ReadInt(r, &c.B2); err != nil {
				return err
			}
			if canoto.IsZero(c.B2) {
				return canoto.ErrZeroValue
			}
		case 5:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if err := canoto.ReadInt(r, &c.C); err != nil {
				return err
			}
			if canoto.IsZero(c.C) {
				return canoto.ErrZeroValue
			}
		case 6:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if err := canoto.ReadInt(r, &c.D); err != nil {
				return err
			}
			if canoto.IsZero(c.D) {
				return canoto.ErrZeroValue
			}
		case 7:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if err := canoto.ReadInt(r, &c.A2); err != nil {
				return err
			}
			if canoto.IsZero(c.A2) {
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
func (c *OneOf) ValidCanoto() bool {
	return true
}

// CalculateCanotoSize calculates the size of the Canoto representation and
// caches it.
//
// It is not safe to call this function concurrently.
func (c *OneOf) CalculateCanotoSize() int {
	c.canotoData.aOneOf = 0
	c.canotoData.bOneOf = 0

	c.canotoData.size = 0
	if !canoto.IsZero(c.A1) {
		c.canotoData.size += len(canoto__OneOf__A1__tag) + canoto.SizeInt(c.A1)
		c.canotoData.aOneOf = 1
	}
	if !canoto.IsZero(c.B1) {
		c.canotoData.size += len(canoto__OneOf__B1__tag) + canoto.SizeInt(c.B1)
		c.canotoData.bOneOf = 3
	}
	if !canoto.IsZero(c.B2) {
		c.canotoData.size += len(canoto__OneOf__B2__tag) + canoto.SizeInt(c.B2)
		c.canotoData.bOneOf = 4
	}
	if !canoto.IsZero(c.C) {
		c.canotoData.size += len(canoto__OneOf__C__tag) + canoto.SizeInt(c.C)
	}
	if !canoto.IsZero(c.D) {
		c.canotoData.size += len(canoto__OneOf__D__tag) + canoto.SizeInt(c.D)
	}
	if !canoto.IsZero(c.A2) {
		c.canotoData.size += len(canoto__OneOf__A2__tag) + canoto.SizeInt(c.A2)
		c.canotoData.aOneOf = 7
	}
	return c.canotoData.size
}

// CachedCanotoSize returns the previously calculated size of the Canoto
// representation from CalculateCanotoSize.
//
// If CalculateCanotoSize has not yet been called, it will return 0.
//
// If the struct has been modified since the last call to CalculateCanotoSize,
// the returned size may be incorrect.
func (c *OneOf) CachedCanotoSize() int {
	return c.canotoData.size
}

// MarshalCanoto returns the Canoto representation of this struct.
//
// It is assumed that this struct is ValidCanoto.
//
// It is not safe to call this function concurrently.
func (c *OneOf) MarshalCanoto() []byte {
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
//
// It is not safe to call this function concurrently.
func (c *OneOf) MarshalCanotoInto(w *canoto.Writer) {
	if !canoto.IsZero(c.A1) {
		canoto.Append(w, canoto__OneOf__A1__tag)
		canoto.AppendInt(w, c.A1)
	}
	if !canoto.IsZero(c.B1) {
		canoto.Append(w, canoto__OneOf__B1__tag)
		canoto.AppendInt(w, c.B1)
	}
	if !canoto.IsZero(c.B2) {
		canoto.Append(w, canoto__OneOf__B2__tag)
		canoto.AppendInt(w, c.B2)
	}
	if !canoto.IsZero(c.C) {
		canoto.Append(w, canoto__OneOf__C__tag)
		canoto.AppendInt(w, c.C)
	}
	if !canoto.IsZero(c.D) {
		canoto.Append(w, canoto__OneOf__D__tag)
		canoto.AppendInt(w, c.D)
	}
	if !canoto.IsZero(c.A2) {
		canoto.Append(w, canoto__OneOf__A2__tag)
		canoto.AppendInt(w, c.A2)
	}
}