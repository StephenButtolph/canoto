// Code generated by canoto. DO NOT EDIT.
// versions:
// 	canoto v0.11.1
// source: spec.go

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

	_ = io.ErrUnexpectedEOF
	_ = utf8.ValidString
)

const (
	canoto__Spec__Name__tag   = "\x0a" // canoto.Tag(1, canoto.Len)
	canoto__Spec__Fields__tag = "\x12" // canoto.Tag(2, canoto.Len)
)

type canotoData_Spec struct {
	size atomic.Int64
}

func (c *Spec) CanotoSpec(types ...reflect.Type) *Spec {
	types = append(types, reflect.TypeOf(Spec{}))
	s := &Spec{
		Name:   "Spec",
		Fields: make([]*FieldType, 0, 2),
	}
	s.Fields = append(s.Fields, &FieldType{
		FieldNumber: 1,
		Name:        "Name",
		TypeString:  true,
	})
	{
		f := &FieldType{
			FieldNumber: 2,
			Name:        "Fields",
			Repeated:    true,
		}
		if index := slices.Index(types, reflect.TypeOf(FieldType{})); index >= 0 {
			f.TypeRecursive = uint64(len(types) - index)
		} else {
			f.TypeMessage = (*FieldType)(nil).CanotoSpec(types...)
		}
		s.Fields = append(s.Fields, f)
	}
	s.CalculateCanotoCache()
	return s
}

// MakeCanoto creates a new empty value.
func (*Spec) MakeCanoto() *Spec {
	return new(Spec)
}

// UnmarshalCanoto unmarshals a Canoto-encoded byte slice into the struct.
//
// During parsing, the canoto cache is saved.
func (c *Spec) UnmarshalCanoto(bytes []byte) error {
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
func (c *Spec) UnmarshalCanotoFrom(r canoto.Reader) error {
	// Zero the struct before unmarshaling.
	*c = Spec{}
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
			if wireType != canoto.Len {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadString(&r, &c.Name); err != nil {
				return err
			}
			if len(c.Name) == 0 {
				return canoto.ErrZeroValue
			}
		case 2:
			if wireType != canoto.Len {
				return canoto.ErrUnexpectedWireType
			}

			// Read the first entry manually because the tag is already
			// stripped.
			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			if err := canoto.ReadBytes(&r, &msgBytes); err != nil {
				return err
			}
			r.Unsafe = originalUnsafe

			// Count the number of additional entries after the first entry.
			countMinus1, err := canoto.CountBytes(r.B, canoto__Spec__Fields__tag)
			if err != nil {
				return err
			}

			c.Fields = canoto.MakeSlice(c.Fields, countMinus1+1)
			if len(msgBytes) != 0 {
				remainingBytes := r.B
				r.B = msgBytes
				c.Fields[0] = canoto.MakePointer(c.Fields[0])
				if err := (c.Fields[0]).UnmarshalCanotoFrom(r); err != nil {
					return err
				}
				r.B = remainingBytes
			}

			// Read the rest of the entries, stripping the tag each time.
			for i := range countMinus1 {
				r.B = r.B[len(canoto__Spec__Fields__tag):]
				r.Unsafe = true
				if err := canoto.ReadBytes(&r, &msgBytes); err != nil {
					return err
				}
				if len(msgBytes) == 0 {
					continue
				}
				r.Unsafe = originalUnsafe

				remainingBytes := r.B
				r.B = msgBytes
				c.Fields[1+i] = canoto.MakePointer(c.Fields[1+i])
				if err := (c.Fields[1+i]).UnmarshalCanotoFrom(r); err != nil {
					return err
				}
				r.B = remainingBytes
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
func (c *Spec) ValidCanoto() bool {
	if c == nil {
		return true
	}
	if !utf8.ValidString(string(c.Name)) {
		return false
	}
	for i := range c.Fields {
		if c.Fields[i] != nil && !(c.Fields[i]).ValidCanoto() {
			return false
		}
	}
	return true
}

// CalculateCanotoCache populates size and OneOf caches based on the current
// values in the struct.
func (c *Spec) CalculateCanotoCache() {
	if c == nil {
		return
	}
	var (
		size int
	)
	if len(c.Name) != 0 {
		size += len(canoto__Spec__Name__tag) + canoto.SizeBytes(c.Name)
	}
	for i := range c.Fields {
		var fieldSize int
		if c.Fields[i] != nil {
			(c.Fields[i]).CalculateCanotoCache()
			fieldSize = (c.Fields[i]).CachedCanotoSize()
		}
		size += len(canoto__Spec__Fields__tag) + canoto.SizeInt(int64(fieldSize)) + fieldSize
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
func (c *Spec) CachedCanotoSize() int {
	if c == nil {
		return 0
	}
	return int(c.canotoData.size.Load())
}

// MarshalCanoto returns the Canoto representation of this struct.
//
// It is assumed that this struct is ValidCanoto.
func (c *Spec) MarshalCanoto() []byte {
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
func (c *Spec) MarshalCanotoInto(w canoto.Writer) canoto.Writer {
	if c == nil {
		return w
	}
	if len(c.Name) != 0 {
		canoto.Append(&w, canoto__Spec__Name__tag)
		canoto.AppendBytes(&w, c.Name)
	}
	for i := range c.Fields {
		canoto.Append(&w, canoto__Spec__Fields__tag)
		var fieldSize int
		if c.Fields[i] != nil {
			fieldSize = (c.Fields[i]).CachedCanotoSize()
		}
		canoto.AppendInt(&w, int64(fieldSize))
		if fieldSize != 0 {
			w = (c.Fields[i]).MarshalCanotoInto(w)
		}
	}
	return w
}

const (
	canoto__FieldType__FieldNumber__tag    = "\x08" // canoto.Tag(1, canoto.Varint)
	canoto__FieldType__Name__tag           = "\x12" // canoto.Tag(2, canoto.Len)
	canoto__FieldType__FixedLength__tag    = "\x18" // canoto.Tag(3, canoto.Varint)
	canoto__FieldType__Repeated__tag       = "\x20" // canoto.Tag(4, canoto.Varint)
	canoto__FieldType__TypeInt__tag        = "\x28" // canoto.Tag(5, canoto.Varint)
	canoto__FieldType__TypeUint__tag       = "\x30" // canoto.Tag(6, canoto.Varint)
	canoto__FieldType__TypeSint__tag       = "\x38" // canoto.Tag(7, canoto.Varint)
	canoto__FieldType__TypeFint__tag       = "\x40" // canoto.Tag(8, canoto.Varint)
	canoto__FieldType__TypeSFint__tag      = "\x48" // canoto.Tag(9, canoto.Varint)
	canoto__FieldType__TypeBool__tag       = "\x50" // canoto.Tag(10, canoto.Varint)
	canoto__FieldType__TypeString__tag     = "\x58" // canoto.Tag(11, canoto.Varint)
	canoto__FieldType__TypeBytes__tag      = "\x60" // canoto.Tag(12, canoto.Varint)
	canoto__FieldType__TypeFixedBytes__tag = "\x68" // canoto.Tag(13, canoto.Varint)
	canoto__FieldType__TypeRecursive__tag  = "\x70" // canoto.Tag(14, canoto.Varint)
	canoto__FieldType__TypeMessage__tag    = "\x7a" // canoto.Tag(15, canoto.Len)
)

type canotoData_FieldType struct {
	size atomic.Int64

	TypeOneOf atomic.Uint32
}

func (c *FieldType) CanotoSpec(types ...reflect.Type) *Spec {
	types = append(types, reflect.TypeOf(FieldType{}))
	s := &Spec{
		Name:   "FieldType",
		Fields: make([]*FieldType, 0, 15),
	}
	s.Fields = append(s.Fields, FieldTypeFromInt(
		FieldType{}.FieldNumber, 
		1, 
		"FieldNumber",
	))
	s.Fields = append(s.Fields, &FieldType{
		FieldNumber: 2, 
		Name: "Name", 
		TypeString: true,
	})
	s.Fields = append(s.Fields, FieldTypeFromInt(
		FieldType{}.FixedLength, 
		3, 
		"FixedLength",
	))
	s.Fields = append(s.Fields, &FieldType{
		FieldNumber: 4, 
		Name: "Repeated", 
		TypeBool: true,
	})
	s.Fields = append(s.Fields, FieldTypeFromInt(
		FieldType{}.TypeInt, 
		5, 
		"TypeInt",
	))
	s.Fields = append(s.Fields, FieldTypeFromInt(
		FieldType{}.TypeUint, 
		6, 
		"TypeUint",
	))
	s.Fields = append(s.Fields, FieldTypeFromInt(
		FieldType{}.TypeSint, 
		7, 
		"TypeSint",
	))
	s.Fields = append(s.Fields, FieldTypeFromInt(
		FieldType{}.TypeFint, 
		8, 
		"TypeFint",
	))
	s.Fields = append(s.Fields, FieldTypeFromInt(
		FieldType{}.TypeSFint, 
		9, 
		"TypeSFint",
	))
	s.Fields = append(s.Fields, &FieldType{
		FieldNumber: 10, 
		Name: "TypeBool", 
		TypeBool: true,
	})
	s.Fields = append(s.Fields, &FieldType{
		FieldNumber: 11, 
		Name: "TypeString", 
		TypeBool: true,
	})
	s.Fields = append(s.Fields, &FieldType{
		FieldNumber: 12, 
		Name: "TypeBytes", 
		TypeBool: true,
	})
	s.Fields = append(s.Fields, FieldTypeFromInt(
		FieldType{}.TypeFixedBytes, 
		13, 
		"TypeFixedBytes",
	))
	s.Fields = append(s.Fields, FieldTypeFromInt(
		FieldType{}.TypeRecursive, 
		14, 
		"TypeRecursive",
	))
	{
		f := &FieldType{
			FieldNumber: 15,
			Name:        "TypeMessage",
			Repeated:    true,
		}
		if index := slices.Index(types, reflect.TypeOf(Spec{})); index >= 0 {
			f.TypeRecursive = uint64(len(types) - index)
		} else {
			f.TypeMessage = (*Spec)(nil).CanotoSpec(types...)
		}
		s.Fields = append(s.Fields, f)
	}
	s.CalculateCanotoCache()
	return s
}

// MakeCanoto creates a new empty value.
func (*FieldType) MakeCanoto() *FieldType {
	return new(FieldType)
}

// UnmarshalCanoto unmarshals a Canoto-encoded byte slice into the struct.
//
// During parsing, the canoto cache is saved.
func (c *FieldType) UnmarshalCanoto(bytes []byte) error {
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
func (c *FieldType) UnmarshalCanotoFrom(r canoto.Reader) error {
	// Zero the struct before unmarshaling.
	*c = FieldType{}
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

			if err := canoto.ReadInt(&r, &c.FieldNumber); err != nil {
				return err
			}
			if canoto.IsZero(c.FieldNumber) {
				return canoto.ErrZeroValue
			}
		case 2:
			if wireType != canoto.Len {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadString(&r, &c.Name); err != nil {
				return err
			}
			if len(c.Name) == 0 {
				return canoto.ErrZeroValue
			}
		case 3:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadInt(&r, &c.FixedLength); err != nil {
				return err
			}
			if canoto.IsZero(c.FixedLength) {
				return canoto.ErrZeroValue
			}
		case 4:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}

			if err := canoto.ReadBool(&r, &c.Repeated); err != nil {
				return err
			}
			if canoto.IsZero(c.Repeated) {
				return canoto.ErrZeroValue
			}
		case 5:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(5) != 0 {
				return canoto.ErrDuplicateOneOf
			}

			if err := canoto.ReadInt(&r, &c.TypeInt); err != nil {
				return err
			}
			if canoto.IsZero(c.TypeInt) {
				return canoto.ErrZeroValue
			}
		case 6:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(6) != 0 {
				return canoto.ErrDuplicateOneOf
			}

			if err := canoto.ReadInt(&r, &c.TypeUint); err != nil {
				return err
			}
			if canoto.IsZero(c.TypeUint) {
				return canoto.ErrZeroValue
			}
		case 7:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(7) != 0 {
				return canoto.ErrDuplicateOneOf
			}

			if err := canoto.ReadInt(&r, &c.TypeSint); err != nil {
				return err
			}
			if canoto.IsZero(c.TypeSint) {
				return canoto.ErrZeroValue
			}
		case 8:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(8) != 0 {
				return canoto.ErrDuplicateOneOf
			}

			if err := canoto.ReadInt(&r, &c.TypeFint); err != nil {
				return err
			}
			if canoto.IsZero(c.TypeFint) {
				return canoto.ErrZeroValue
			}
		case 9:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(9) != 0 {
				return canoto.ErrDuplicateOneOf
			}

			if err := canoto.ReadInt(&r, &c.TypeSFint); err != nil {
				return err
			}
			if canoto.IsZero(c.TypeSFint) {
				return canoto.ErrZeroValue
			}
		case 10:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(10) != 0 {
				return canoto.ErrDuplicateOneOf
			}

			if err := canoto.ReadBool(&r, &c.TypeBool); err != nil {
				return err
			}
			if canoto.IsZero(c.TypeBool) {
				return canoto.ErrZeroValue
			}
		case 11:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(11) != 0 {
				return canoto.ErrDuplicateOneOf
			}

			if err := canoto.ReadBool(&r, &c.TypeString); err != nil {
				return err
			}
			if canoto.IsZero(c.TypeString) {
				return canoto.ErrZeroValue
			}
		case 12:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(12) != 0 {
				return canoto.ErrDuplicateOneOf
			}

			if err := canoto.ReadBool(&r, &c.TypeBytes); err != nil {
				return err
			}
			if canoto.IsZero(c.TypeBytes) {
				return canoto.ErrZeroValue
			}
		case 13:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(13) != 0 {
				return canoto.ErrDuplicateOneOf
			}

			if err := canoto.ReadInt(&r, &c.TypeFixedBytes); err != nil {
				return err
			}
			if canoto.IsZero(c.TypeFixedBytes) {
				return canoto.ErrZeroValue
			}
		case 14:
			if wireType != canoto.Varint {
				return canoto.ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(14) != 0 {
				return canoto.ErrDuplicateOneOf
			}

			if err := canoto.ReadInt(&r, &c.TypeRecursive); err != nil {
				return err
			}
			if canoto.IsZero(c.TypeRecursive) {
				return canoto.ErrZeroValue
			}
		case 15:
			if wireType != canoto.Len {
				return canoto.ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(15) != 0 {
				return canoto.ErrDuplicateOneOf
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
			c.TypeMessage = canoto.MakePointer(c.TypeMessage)
			if err := (c.TypeMessage).UnmarshalCanotoFrom(r); err != nil {
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
func (c *FieldType) ValidCanoto() bool {
	if c == nil {
		return true
	}
	var (
		TypeOneOf uint32
	)
	if !canoto.IsZero(c.TypeInt) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 5
	}
	if !canoto.IsZero(c.TypeUint) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 6
	}
	if !canoto.IsZero(c.TypeSint) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 7
	}
	if !canoto.IsZero(c.TypeFint) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 8
	}
	if !canoto.IsZero(c.TypeSFint) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 9
	}
	if !canoto.IsZero(c.TypeBool) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 10
	}
	if !canoto.IsZero(c.TypeString) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 11
	}
	if !canoto.IsZero(c.TypeBytes) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 12
	}
	if !canoto.IsZero(c.TypeFixedBytes) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 13
	}
	if !canoto.IsZero(c.TypeRecursive) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 14
	}
	if c.TypeMessage != nil {
		(c.TypeMessage).CalculateCanotoCache()
		if (c.TypeMessage).CachedCanotoSize() != 0 {
			if TypeOneOf != 0 {
				return false
			}
			TypeOneOf = 15
		}
	}
	if !utf8.ValidString(string(c.Name)) {
		return false
	}
	if c.TypeMessage != nil && !(c.TypeMessage).ValidCanoto() {
		return false
	}
	return true
}

// CalculateCanotoCache populates size and OneOf caches based on the current
// values in the struct.
func (c *FieldType) CalculateCanotoCache() {
	if c == nil {
		return
	}
	var (
		size      int
		TypeOneOf uint32
	)
	if !canoto.IsZero(c.FieldNumber) {
		size += len(canoto__FieldType__FieldNumber__tag) + canoto.SizeInt(c.FieldNumber)
	}
	if len(c.Name) != 0 {
		size += len(canoto__FieldType__Name__tag) + canoto.SizeBytes(c.Name)
	}
	if !canoto.IsZero(c.FixedLength) {
		size += len(canoto__FieldType__FixedLength__tag) + canoto.SizeInt(c.FixedLength)
	}
	if !canoto.IsZero(c.Repeated) {
		size += len(canoto__FieldType__Repeated__tag) + canoto.SizeBool
	}
	if !canoto.IsZero(c.TypeInt) {
		size += len(canoto__FieldType__TypeInt__tag) + canoto.SizeInt(c.TypeInt)
		TypeOneOf = 5
	}
	if !canoto.IsZero(c.TypeUint) {
		size += len(canoto__FieldType__TypeUint__tag) + canoto.SizeInt(c.TypeUint)
		TypeOneOf = 6
	}
	if !canoto.IsZero(c.TypeSint) {
		size += len(canoto__FieldType__TypeSint__tag) + canoto.SizeInt(c.TypeSint)
		TypeOneOf = 7
	}
	if !canoto.IsZero(c.TypeFint) {
		size += len(canoto__FieldType__TypeFint__tag) + canoto.SizeInt(c.TypeFint)
		TypeOneOf = 8
	}
	if !canoto.IsZero(c.TypeSFint) {
		size += len(canoto__FieldType__TypeSFint__tag) + canoto.SizeInt(c.TypeSFint)
		TypeOneOf = 9
	}
	if !canoto.IsZero(c.TypeBool) {
		size += len(canoto__FieldType__TypeBool__tag) + canoto.SizeBool
		TypeOneOf = 10
	}
	if !canoto.IsZero(c.TypeString) {
		size += len(canoto__FieldType__TypeString__tag) + canoto.SizeBool
		TypeOneOf = 11
	}
	if !canoto.IsZero(c.TypeBytes) {
		size += len(canoto__FieldType__TypeBytes__tag) + canoto.SizeBool
		TypeOneOf = 12
	}
	if !canoto.IsZero(c.TypeFixedBytes) {
		size += len(canoto__FieldType__TypeFixedBytes__tag) + canoto.SizeInt(c.TypeFixedBytes)
		TypeOneOf = 13
	}
	if !canoto.IsZero(c.TypeRecursive) {
		size += len(canoto__FieldType__TypeRecursive__tag) + canoto.SizeInt(c.TypeRecursive)
		TypeOneOf = 14
	}
	if c.TypeMessage != nil {
		(c.TypeMessage).CalculateCanotoCache()
		if fieldSize := (c.TypeMessage).CachedCanotoSize(); fieldSize != 0 {
			size += len(canoto__FieldType__TypeMessage__tag) + canoto.SizeInt(int64(fieldSize)) + fieldSize
			TypeOneOf = 15
		}
	}
	c.canotoData.size.Store(int64(size))
	c.canotoData.TypeOneOf.Store(TypeOneOf)
}

// CachedCanotoSize returns the previously calculated size of the Canoto
// representation from CalculateCanotoCache.
//
// If CalculateCanotoCache has not yet been called, it will return 0.
//
// If the struct has been modified since the last call to CalculateCanotoCache,
// the returned size may be incorrect.
func (c *FieldType) CachedCanotoSize() int {
	if c == nil {
		return 0
	}
	return int(c.canotoData.size.Load())
}

// CachedWhichOneOfType returns the previously calculated field number used
// to represent Type.
//
// This field is cached by UnmarshalCanoto, UnmarshalCanotoFrom, and
// CalculateCanotoCache.
//
// If the field has not yet been cached, it will return 0.
//
// If the struct has been modified since the field was last cached, the returned
// field number may be incorrect.
func (c *FieldType) CachedWhichOneOfType() uint32 {
	return c.canotoData.TypeOneOf.Load()
}

// MarshalCanoto returns the Canoto representation of this struct.
//
// It is assumed that this struct is ValidCanoto.
func (c *FieldType) MarshalCanoto() []byte {
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
func (c *FieldType) MarshalCanotoInto(w canoto.Writer) canoto.Writer {
	if c == nil {
		return w
	}
	if !canoto.IsZero(c.FieldNumber) {
		canoto.Append(&w, canoto__FieldType__FieldNumber__tag)
		canoto.AppendInt(&w, c.FieldNumber)
	}
	if len(c.Name) != 0 {
		canoto.Append(&w, canoto__FieldType__Name__tag)
		canoto.AppendBytes(&w, c.Name)
	}
	if !canoto.IsZero(c.FixedLength) {
		canoto.Append(&w, canoto__FieldType__FixedLength__tag)
		canoto.AppendInt(&w, c.FixedLength)
	}
	if !canoto.IsZero(c.Repeated) {
		canoto.Append(&w, canoto__FieldType__Repeated__tag)
		canoto.AppendBool(&w, true)
	}
	if !canoto.IsZero(c.TypeInt) {
		canoto.Append(&w, canoto__FieldType__TypeInt__tag)
		canoto.AppendInt(&w, c.TypeInt)
	}
	if !canoto.IsZero(c.TypeUint) {
		canoto.Append(&w, canoto__FieldType__TypeUint__tag)
		canoto.AppendInt(&w, c.TypeUint)
	}
	if !canoto.IsZero(c.TypeSint) {
		canoto.Append(&w, canoto__FieldType__TypeSint__tag)
		canoto.AppendInt(&w, c.TypeSint)
	}
	if !canoto.IsZero(c.TypeFint) {
		canoto.Append(&w, canoto__FieldType__TypeFint__tag)
		canoto.AppendInt(&w, c.TypeFint)
	}
	if !canoto.IsZero(c.TypeSFint) {
		canoto.Append(&w, canoto__FieldType__TypeSFint__tag)
		canoto.AppendInt(&w, c.TypeSFint)
	}
	if !canoto.IsZero(c.TypeBool) {
		canoto.Append(&w, canoto__FieldType__TypeBool__tag)
		canoto.AppendBool(&w, true)
	}
	if !canoto.IsZero(c.TypeString) {
		canoto.Append(&w, canoto__FieldType__TypeString__tag)
		canoto.AppendBool(&w, true)
	}
	if !canoto.IsZero(c.TypeBytes) {
		canoto.Append(&w, canoto__FieldType__TypeBytes__tag)
		canoto.AppendBool(&w, true)
	}
	if !canoto.IsZero(c.TypeFixedBytes) {
		canoto.Append(&w, canoto__FieldType__TypeFixedBytes__tag)
		canoto.AppendInt(&w, c.TypeFixedBytes)
	}
	if !canoto.IsZero(c.TypeRecursive) {
		canoto.Append(&w, canoto__FieldType__TypeRecursive__tag)
		canoto.AppendInt(&w, c.TypeRecursive)
	}
	if c.TypeMessage != nil {
		if fieldSize := (c.TypeMessage).CachedCanotoSize(); fieldSize != 0 {
			canoto.Append(&w, canoto__FieldType__TypeMessage__tag)
			canoto.AppendInt(&w, int64(fieldSize))
			w = (c.TypeMessage).MarshalCanotoInto(w)
		}
	}
	return w
}
