// Code generated by canoto. DO NOT EDIT.
// versions:
// 	canoto v0.11.1
// source: spec_test.go

package reflect

import (
	"io"
	"reflect"
	"slices"
	"sync/atomic"
	"unicode/utf8"

	"github.com/StephenButtolph/canoto"
)

// Ensure that unused imports do not error
var (
	_ atomic.Int64

	_ = slices.Index[[]reflect.Type, reflect.Type]
	_ = io.ErrUnexpectedEOF
	_ = utf8.ValidString
)

const (
	canoto__testMessage__Int8__tag       = "\x08"     // canoto.Tag(1, canoto.Varint)
	canoto__testMessage__Int16__tag      = "\x10"     // canoto.Tag(2, canoto.Varint)
	canoto__testMessage__Int32__tag      = "\x18"     // canoto.Tag(3, canoto.Varint)
	canoto__testMessage__Int64__tag      = "\x20"     // canoto.Tag(4, canoto.Varint)
	canoto__testMessage__Uint8__tag      = "\x28"     // canoto.Tag(5, canoto.Varint)
	canoto__testMessage__Uint16__tag     = "\x30"     // canoto.Tag(6, canoto.Varint)
	canoto__testMessage__Uint32__tag     = "\x38"     // canoto.Tag(7, canoto.Varint)
	canoto__testMessage__Uint64__tag     = "\x40"     // canoto.Tag(8, canoto.Varint)
	canoto__testMessage__Sint8__tag      = "\x48"     // canoto.Tag(9, canoto.Varint)
	canoto__testMessage__Sint16__tag     = "\x50"     // canoto.Tag(10, canoto.Varint)
	canoto__testMessage__Sint32__tag     = "\x58"     // canoto.Tag(11, canoto.Varint)
	canoto__testMessage__Sint64__tag     = "\x60"     // canoto.Tag(12, canoto.Varint)
	canoto__testMessage__Fixed32__tag    = "\x6d"     // canoto.Tag(13, canoto.I32)
	canoto__testMessage__Fixed64__tag    = "\x71"     // canoto.Tag(14, canoto.I64)
	canoto__testMessage__Sfixed32__tag   = "\x7d"     // canoto.Tag(15, canoto.I32)
	canoto__testMessage__Sfixed64__tag   = "\x81\x01" // canoto.Tag(16, canoto.I64)
	canoto__testMessage__Bool__tag       = "\x88\x01" // canoto.Tag(17, canoto.Varint)
	canoto__testMessage__String__tag     = "\x92\x01" // canoto.Tag(18, canoto.Len)
	canoto__testMessage__Bytes__tag      = "\x9a\x01" // canoto.Tag(19, canoto.Len)
	canoto__testMessage__FixedBytes__tag = "\xa2\x01" // canoto.Tag(20, canoto.Len)
	canoto__testMessage__Recursive__tag  = "\xaa\x01" // canoto.Tag(21, canoto.Len)
	canoto__testMessage__Message__tag    = "\xb2\x01" // canoto.Tag(22, canoto.Len)
)

type canotoData_testMessage struct {
	size atomic.Int64
}

// CanotoSpec returns the specification of this canoto message.
func (*testMessage) CanotoSpec(types ...reflect.Type) *canoto.Spec {
	types = append(types, reflect.TypeOf(testMessage{}))
	var zero testMessage
	s := &canoto.Spec{
		Name: "testMessage",
		Fields: []*canoto.FieldType{
			canoto.FieldTypeFromInt(
				zero.Int8,
				1,
				"Int8",
				"",
			),
			canoto.FieldTypeFromInt(
				zero.Int16,
				2,
				"Int16",
				"",
			),
			canoto.FieldTypeFromInt(
				zero.Int32,
				3,
				"Int32",
				"",
			),
			canoto.FieldTypeFromInt(
				zero.Int64,
				4,
				"Int64",
				"",
			),
			canoto.FieldTypeFromInt(
				zero.Uint8,
				5,
				"Uint8",
				"",
			),
			canoto.FieldTypeFromInt(
				zero.Uint16,
				6,
				"Uint16",
				"",
			),
			canoto.FieldTypeFromInt(
				zero.Uint32,
				7,
				"Uint32",
				"",
			),
			canoto.FieldTypeFromInt(
				zero.Uint64,
				8,
				"Uint64",
				"",
			),
			canoto.FieldTypeFromSint(
				zero.Sint8,
				9,
				"Sint8",
				"",
			),
			canoto.FieldTypeFromSint(
				zero.Sint16,
				10,
				"Sint16",
				"",
			),
			canoto.FieldTypeFromSint(
				zero.Sint32,
				11,
				"Sint32",
				"",
			),
			canoto.FieldTypeFromSint(
				zero.Sint64,
				12,
				"Sint64",
				"",
			),
			canoto.FieldTypeFromFint(
				zero.Fixed32,
				13,
				"Fixed32",
				"",
			),
			canoto.FieldTypeFromFint(
				zero.Fixed64,
				14,
				"Fixed64",
				"",
			),
			canoto.FieldTypeFromFint(
				zero.Sfixed32,
				15,
				"Sfixed32",
				"",
			),
			canoto.FieldTypeFromFint(
				zero.Sfixed64,
				16,
				"Sfixed64",
				"",
			),
			{
				FieldNumber: 17,
				Name:        "Bool",
				OneOf:       "",
				TypeBool:    true,
			},
			{
				FieldNumber: 18,
				Name:        "String",
				OneOf:       "",
				TypeString:  true,
			},
			{
				FieldNumber: 19,
				Name:        "Bytes",
				OneOf:       "",
				TypeBytes:   true,
			},
			{
				FieldNumber:    20,
				Name:           "FixedBytes",
				OneOf:          "",
				TypeFixedBytes: uint64(len(zero.FixedBytes)),
			},
			canoto.FieldTypeFromPointer(
				(zero.Recursive),
				21,
				"Recursive",
				0,
				false,
				"",
				types,
			),
			canoto.FieldTypeFromPointer(
				(&zero.Message),
				22,
				"Message",
				0,
				false,
				"",
				types,
			),
		},
	}
	s.CalculateCanotoCache()
	return s
}

// MakeCanoto creates a new empty value.
func (*testMessage) MakeCanoto() *testMessage {
	return new(testMessage)
}

// UnmarshalCanoto unmarshals a Canoto-encoded byte slice into the struct.
//
// During parsing, the canoto cache is saved.
func (c *testMessage) UnmarshalCanoto(bytes []byte) error {
	r := canoto.Reader{
		B: bytes,
	}
	return c.UnmarshalCanotoFrom(r)
}

// UnmarshalCanotoFrom populates the struct from a canoto.Reader. Most users
// should just use UnmarshalCanoto.
//
// During parsing, the canoto cache is saved.
//
// This function enables configuration of reader options.
func (c *testMessage) UnmarshalCanotoFrom(r canoto.Reader) error {
	// Zero the struct before unmarshaling.
	*c = testMessage{}
	c.canotoData.size.Store(int64(len(r.B)))

	var minField uint32
	for canoto.HasNext(&r) {
		field, wireType, err := canoto.ReadTag(&r)
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

			if err := canoto.ReadInt(&r, &c.Int8); err != nil {
				return err
			}
			if canoto.IsZero(c.Int8) {
				return canoto.ErrZeroValue
			}
		case 2:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadInt(&r, &c.Int16); err != nil {
				return err
			}
			if canoto.IsZero(c.Int16) {
				return canoto.ErrZeroValue
			}
		case 3:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadInt(&r, &c.Int32); err != nil {
				return err
			}
			if canoto.IsZero(c.Int32) {
				return canoto.ErrZeroValue
			}
		case 4:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadInt(&r, &c.Int64); err != nil {
				return err
			}
			if canoto.IsZero(c.Int64) {
				return canoto.ErrZeroValue
			}
		case 5:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadInt(&r, &c.Uint8); err != nil {
				return err
			}
			if canoto.IsZero(c.Uint8) {
				return canoto.ErrZeroValue
			}
		case 6:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadInt(&r, &c.Uint16); err != nil {
				return err
			}
			if canoto.IsZero(c.Uint16) {
				return canoto.ErrZeroValue
			}
		case 7:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadInt(&r, &c.Uint32); err != nil {
				return err
			}
			if canoto.IsZero(c.Uint32) {
				return canoto.ErrZeroValue
			}
		case 8:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadInt(&r, &c.Uint64); err != nil {
				return err
			}
			if canoto.IsZero(c.Uint64) {
				return canoto.ErrZeroValue
			}
		case 9:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadSint(&r, &c.Sint8); err != nil {
				return err
			}
			if canoto.IsZero(c.Sint8) {
				return canoto.ErrZeroValue
			}
		case 10:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadSint(&r, &c.Sint16); err != nil {
				return err
			}
			if canoto.IsZero(c.Sint16) {
				return canoto.ErrZeroValue
			}
		case 11:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadSint(&r, &c.Sint32); err != nil {
				return err
			}
			if canoto.IsZero(c.Sint32) {
				return canoto.ErrZeroValue
			}
		case 12:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadSint(&r, &c.Sint64); err != nil {
				return err
			}
			if canoto.IsZero(c.Sint64) {
				return canoto.ErrZeroValue
			}
		case 13:
			if wireType != canoto.I32 {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadFint32(&r, &c.Fixed32); err != nil {
				return err
			}
			if canoto.IsZero(c.Fixed32) {
				return canoto.ErrZeroValue
			}
		case 14:
			if wireType != canoto.I64 {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadFint64(&r, &c.Fixed64); err != nil {
				return err
			}
			if canoto.IsZero(c.Fixed64) {
				return canoto.ErrZeroValue
			}
		case 15:
			if wireType != canoto.I32 {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadFint32(&r, &c.Sfixed32); err != nil {
				return err
			}
			if canoto.IsZero(c.Sfixed32) {
				return canoto.ErrZeroValue
			}
		case 16:
			if wireType != canoto.I64 {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadFint64(&r, &c.Sfixed64); err != nil {
				return err
			}
			if canoto.IsZero(c.Sfixed64) {
				return canoto.ErrZeroValue
			}
		case 17:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadBool(&r, &c.Bool); err != nil {
				return err
			}
			if canoto.IsZero(c.Bool) {
				return canoto.ErrZeroValue
			}
		case 18:
			if wireType != canoto.Len {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadString(&r, &c.String); err != nil {
				return err
			}
			if len(c.String) == 0 {
				return canoto.ErrZeroValue
			}
		case 19:
			if wireType != canoto.Len {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadBytes(&r, &c.Bytes); err != nil {
				return err
			}
			if len(c.Bytes) == 0 {
				return canoto.ErrZeroValue
			}
		case 20:
			if wireType != canoto.Len {
				return canoto.ErrUnexpectedWireType
			}

			const (
				expectedLength      = len(c.FixedBytes)
				expectedLengthInt64 = int64(expectedLength)
			)
			var length int64
			if err := canoto.ReadInt(&r, &length); err != nil {
				return err
			}
			if expectedLength > len(r.B) {
				return io.ErrUnexpectedEOF
			}
			if length != expectedLengthInt64 {
				return canoto.ErrInvalidLength
			}

			copy((&c.FixedBytes)[:], r.B)
			if canoto.IsZero(c.FixedBytes) {
				return canoto.ErrZeroValue
			}
			r.B = r.B[expectedLength:]
		case 21:
			if wireType != canoto.Len {
				return canoto.ErrUnexpectedWireType
			}

			// Read the bytes for the field.
			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			if err := canoto.ReadBytes(&r, &msgBytes); err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return canoto.ErrZeroValue
			}
			r.Unsafe = originalUnsafe

			// Unmarshal the field from the bytes.
			remainingBytes := r.B
			r.B = msgBytes
			c.Recursive = canoto.MakePointer(c.Recursive)
			if err := (c.Recursive).UnmarshalCanotoFrom(r); err != nil {
				return err
			}
			r.B = remainingBytes
		case 22:
			if wireType != canoto.Len {
				return canoto.ErrUnexpectedWireType
			}

			// Read the bytes for the field.
			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			if err := canoto.ReadBytes(&r, &msgBytes); err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return canoto.ErrZeroValue
			}
			r.Unsafe = originalUnsafe

			// Unmarshal the field from the bytes.
			remainingBytes := r.B
			r.B = msgBytes
			if err := (&c.Message).UnmarshalCanotoFrom(r); err != nil {
				return err
			}
			r.B = remainingBytes
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
// Specifically, ValidCanoto ensures:
// 1. All OneOfs are specified at most once.
// 2. All strings are valid utf-8.
// 3. All custom fields are ValidCanoto.
func (c *testMessage) ValidCanoto() bool {
	if c == nil {
		return true
	}
	if !utf8.ValidString(string(c.String)) {
		return false
	}
	if c.Recursive != nil && !(c.Recursive).ValidCanoto() {
		return false
	}
	if !(&c.Message).ValidCanoto() {
		return false
	}
	return true
}

// CalculateCanotoCache populates size and OneOf caches based on the current
// values in the struct.
func (c *testMessage) CalculateCanotoCache() {
	if c == nil {
		return
	}
	var (
		size int
	)
	if !canoto.IsZero(c.Int8) {
		size += len(canoto__testMessage__Int8__tag) + canoto.SizeInt(c.Int8)
	}
	if !canoto.IsZero(c.Int16) {
		size += len(canoto__testMessage__Int16__tag) + canoto.SizeInt(c.Int16)
	}
	if !canoto.IsZero(c.Int32) {
		size += len(canoto__testMessage__Int32__tag) + canoto.SizeInt(c.Int32)
	}
	if !canoto.IsZero(c.Int64) {
		size += len(canoto__testMessage__Int64__tag) + canoto.SizeInt(c.Int64)
	}
	if !canoto.IsZero(c.Uint8) {
		size += len(canoto__testMessage__Uint8__tag) + canoto.SizeInt(c.Uint8)
	}
	if !canoto.IsZero(c.Uint16) {
		size += len(canoto__testMessage__Uint16__tag) + canoto.SizeInt(c.Uint16)
	}
	if !canoto.IsZero(c.Uint32) {
		size += len(canoto__testMessage__Uint32__tag) + canoto.SizeInt(c.Uint32)
	}
	if !canoto.IsZero(c.Uint64) {
		size += len(canoto__testMessage__Uint64__tag) + canoto.SizeInt(c.Uint64)
	}
	if !canoto.IsZero(c.Sint8) {
		size += len(canoto__testMessage__Sint8__tag) + canoto.SizeSint(c.Sint8)
	}
	if !canoto.IsZero(c.Sint16) {
		size += len(canoto__testMessage__Sint16__tag) + canoto.SizeSint(c.Sint16)
	}
	if !canoto.IsZero(c.Sint32) {
		size += len(canoto__testMessage__Sint32__tag) + canoto.SizeSint(c.Sint32)
	}
	if !canoto.IsZero(c.Sint64) {
		size += len(canoto__testMessage__Sint64__tag) + canoto.SizeSint(c.Sint64)
	}
	if !canoto.IsZero(c.Fixed32) {
		size += len(canoto__testMessage__Fixed32__tag) + canoto.SizeFint32
	}
	if !canoto.IsZero(c.Fixed64) {
		size += len(canoto__testMessage__Fixed64__tag) + canoto.SizeFint64
	}
	if !canoto.IsZero(c.Sfixed32) {
		size += len(canoto__testMessage__Sfixed32__tag) + canoto.SizeFint32
	}
	if !canoto.IsZero(c.Sfixed64) {
		size += len(canoto__testMessage__Sfixed64__tag) + canoto.SizeFint64
	}
	if !canoto.IsZero(c.Bool) {
		size += len(canoto__testMessage__Bool__tag) + canoto.SizeBool
	}
	if len(c.String) != 0 {
		size += len(canoto__testMessage__String__tag) + canoto.SizeBytes(c.String)
	}
	if len(c.Bytes) != 0 {
		size += len(canoto__testMessage__Bytes__tag) + canoto.SizeBytes(c.Bytes)
	}
	if !canoto.IsZero(c.FixedBytes) {
		size += len(canoto__testMessage__FixedBytes__tag) + canoto.SizeBytes((&c.FixedBytes)[:])
	}
	if c.Recursive != nil {
		(c.Recursive).CalculateCanotoCache()
		if fieldSize := (c.Recursive).CachedCanotoSize(); fieldSize != 0 {
			size += len(canoto__testMessage__Recursive__tag) + canoto.SizeInt(int64(fieldSize)) + fieldSize
		}
	}
	(&c.Message).CalculateCanotoCache()
	if fieldSize := (&c.Message).CachedCanotoSize(); fieldSize != 0 {
		size += len(canoto__testMessage__Message__tag) + canoto.SizeInt(int64(fieldSize)) + fieldSize
	}
	c.canotoData.size.Store(int64(size))
}

// CachedCanotoSize returns the previously calculated size of the Canoto
// representation from CalculateCanotoCache.
//
// If CalculateCanotoCache has not yet been called, it will return 0.
//
// If the struct has been modified since the last call to CalculateCanotoCache,
// the returned size may be incorrect.
func (c *testMessage) CachedCanotoSize() int {
	if c == nil {
		return 0
	}
	return int(c.canotoData.size.Load())
}

// MarshalCanoto returns the Canoto representation of this struct.
//
// It is assumed that this struct is ValidCanoto.
func (c *testMessage) MarshalCanoto() []byte {
	c.CalculateCanotoCache()
	w := canoto.Writer{
		B: make([]byte, 0, c.CachedCanotoSize()),
	}
	w = c.MarshalCanotoInto(w)
	return w.B
}

// MarshalCanotoInto writes the struct into a canoto.Writer and returns the
// resulting canoto.Writer. Most users should just use MarshalCanoto.
//
// It is assumed that CalculateCanotoCache has been called since the last
// modification to this struct.
//
// It is assumed that this struct is ValidCanoto.
func (c *testMessage) MarshalCanotoInto(w canoto.Writer) canoto.Writer {
	if c == nil {
		return w
	}
	if !canoto.IsZero(c.Int8) {
		canoto.Append(&w, canoto__testMessage__Int8__tag)
		canoto.AppendInt(&w, c.Int8)
	}
	if !canoto.IsZero(c.Int16) {
		canoto.Append(&w, canoto__testMessage__Int16__tag)
		canoto.AppendInt(&w, c.Int16)
	}
	if !canoto.IsZero(c.Int32) {
		canoto.Append(&w, canoto__testMessage__Int32__tag)
		canoto.AppendInt(&w, c.Int32)
	}
	if !canoto.IsZero(c.Int64) {
		canoto.Append(&w, canoto__testMessage__Int64__tag)
		canoto.AppendInt(&w, c.Int64)
	}
	if !canoto.IsZero(c.Uint8) {
		canoto.Append(&w, canoto__testMessage__Uint8__tag)
		canoto.AppendInt(&w, c.Uint8)
	}
	if !canoto.IsZero(c.Uint16) {
		canoto.Append(&w, canoto__testMessage__Uint16__tag)
		canoto.AppendInt(&w, c.Uint16)
	}
	if !canoto.IsZero(c.Uint32) {
		canoto.Append(&w, canoto__testMessage__Uint32__tag)
		canoto.AppendInt(&w, c.Uint32)
	}
	if !canoto.IsZero(c.Uint64) {
		canoto.Append(&w, canoto__testMessage__Uint64__tag)
		canoto.AppendInt(&w, c.Uint64)
	}
	if !canoto.IsZero(c.Sint8) {
		canoto.Append(&w, canoto__testMessage__Sint8__tag)
		canoto.AppendSint(&w, c.Sint8)
	}
	if !canoto.IsZero(c.Sint16) {
		canoto.Append(&w, canoto__testMessage__Sint16__tag)
		canoto.AppendSint(&w, c.Sint16)
	}
	if !canoto.IsZero(c.Sint32) {
		canoto.Append(&w, canoto__testMessage__Sint32__tag)
		canoto.AppendSint(&w, c.Sint32)
	}
	if !canoto.IsZero(c.Sint64) {
		canoto.Append(&w, canoto__testMessage__Sint64__tag)
		canoto.AppendSint(&w, c.Sint64)
	}
	if !canoto.IsZero(c.Fixed32) {
		canoto.Append(&w, canoto__testMessage__Fixed32__tag)
		canoto.AppendFint32(&w, c.Fixed32)
	}
	if !canoto.IsZero(c.Fixed64) {
		canoto.Append(&w, canoto__testMessage__Fixed64__tag)
		canoto.AppendFint64(&w, c.Fixed64)
	}
	if !canoto.IsZero(c.Sfixed32) {
		canoto.Append(&w, canoto__testMessage__Sfixed32__tag)
		canoto.AppendFint32(&w, c.Sfixed32)
	}
	if !canoto.IsZero(c.Sfixed64) {
		canoto.Append(&w, canoto__testMessage__Sfixed64__tag)
		canoto.AppendFint64(&w, c.Sfixed64)
	}
	if !canoto.IsZero(c.Bool) {
		canoto.Append(&w, canoto__testMessage__Bool__tag)
		canoto.AppendBool(&w, true)
	}
	if len(c.String) != 0 {
		canoto.Append(&w, canoto__testMessage__String__tag)
		canoto.AppendBytes(&w, c.String)
	}
	if len(c.Bytes) != 0 {
		canoto.Append(&w, canoto__testMessage__Bytes__tag)
		canoto.AppendBytes(&w, c.Bytes)
	}
	if !canoto.IsZero(c.FixedBytes) {
		canoto.Append(&w, canoto__testMessage__FixedBytes__tag)
		canoto.AppendBytes(&w, (&c.FixedBytes)[:])
	}
	if c.Recursive != nil {
		if fieldSize := (c.Recursive).CachedCanotoSize(); fieldSize != 0 {
			canoto.Append(&w, canoto__testMessage__Recursive__tag)
			canoto.AppendInt(&w, int64(fieldSize))
			w = (c.Recursive).MarshalCanotoInto(w)
		}
	}
	if fieldSize := (&c.Message).CachedCanotoSize(); fieldSize != 0 {
		canoto.Append(&w, canoto__testMessage__Message__tag)
		canoto.AppendInt(&w, int64(fieldSize))
		w = (&c.Message).MarshalCanotoInto(w)
	}
	return w
}

const (
	canoto__testSimpleMessage__Int8__tag = "\x08" // canoto.Tag(1, canoto.Varint)
)

type canotoData_testSimpleMessage struct {
	size atomic.Int64
}

// CanotoSpec returns the specification of this canoto message.
func (*testSimpleMessage) CanotoSpec(types ...reflect.Type) *canoto.Spec {
	types = append(types, reflect.TypeOf(testSimpleMessage{}))
	var zero testSimpleMessage
	s := &canoto.Spec{
		Name: "testSimpleMessage",
		Fields: []*canoto.FieldType{
			canoto.FieldTypeFromInt(
				zero.Int8,
				1,
				"Int8",
				"",
			),
		},
	}
	s.CalculateCanotoCache()
	return s
}

// MakeCanoto creates a new empty value.
func (*testSimpleMessage) MakeCanoto() *testSimpleMessage {
	return new(testSimpleMessage)
}

// UnmarshalCanoto unmarshals a Canoto-encoded byte slice into the struct.
//
// During parsing, the canoto cache is saved.
func (c *testSimpleMessage) UnmarshalCanoto(bytes []byte) error {
	r := canoto.Reader{
		B: bytes,
	}
	return c.UnmarshalCanotoFrom(r)
}

// UnmarshalCanotoFrom populates the struct from a canoto.Reader. Most users
// should just use UnmarshalCanoto.
//
// During parsing, the canoto cache is saved.
//
// This function enables configuration of reader options.
func (c *testSimpleMessage) UnmarshalCanotoFrom(r canoto.Reader) error {
	// Zero the struct before unmarshaling.
	*c = testSimpleMessage{}
	c.canotoData.size.Store(int64(len(r.B)))

	var minField uint32
	for canoto.HasNext(&r) {
		field, wireType, err := canoto.ReadTag(&r)
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

			if err := canoto.ReadInt(&r, &c.Int8); err != nil {
				return err
			}
			if canoto.IsZero(c.Int8) {
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
// Specifically, ValidCanoto ensures:
// 1. All OneOfs are specified at most once.
// 2. All strings are valid utf-8.
// 3. All custom fields are ValidCanoto.
func (c *testSimpleMessage) ValidCanoto() bool {
	if c == nil {
		return true
	}
	return true
}

// CalculateCanotoCache populates size and OneOf caches based on the current
// values in the struct.
func (c *testSimpleMessage) CalculateCanotoCache() {
	if c == nil {
		return
	}
	var (
		size int
	)
	if !canoto.IsZero(c.Int8) {
		size += len(canoto__testSimpleMessage__Int8__tag) + canoto.SizeInt(c.Int8)
	}
	c.canotoData.size.Store(int64(size))
}

// CachedCanotoSize returns the previously calculated size of the Canoto
// representation from CalculateCanotoCache.
//
// If CalculateCanotoCache has not yet been called, it will return 0.
//
// If the struct has been modified since the last call to CalculateCanotoCache,
// the returned size may be incorrect.
func (c *testSimpleMessage) CachedCanotoSize() int {
	if c == nil {
		return 0
	}
	return int(c.canotoData.size.Load())
}

// MarshalCanoto returns the Canoto representation of this struct.
//
// It is assumed that this struct is ValidCanoto.
func (c *testSimpleMessage) MarshalCanoto() []byte {
	c.CalculateCanotoCache()
	w := canoto.Writer{
		B: make([]byte, 0, c.CachedCanotoSize()),
	}
	w = c.MarshalCanotoInto(w)
	return w.B
}

// MarshalCanotoInto writes the struct into a canoto.Writer and returns the
// resulting canoto.Writer. Most users should just use MarshalCanoto.
//
// It is assumed that CalculateCanotoCache has been called since the last
// modification to this struct.
//
// It is assumed that this struct is ValidCanoto.
func (c *testSimpleMessage) MarshalCanotoInto(w canoto.Writer) canoto.Writer {
	if c == nil {
		return w
	}
	if !canoto.IsZero(c.Int8) {
		canoto.Append(&w, canoto__testSimpleMessage__Int8__tag)
		canoto.AppendInt(&w, c.Int8)
	}
	return w
}
