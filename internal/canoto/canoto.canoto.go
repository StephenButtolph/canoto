// Code generated by canoto. DO NOT EDIT.
// versions:
// 	canoto v0.15.0
// source: canoto.go

package canoto

import (
	"io"
	"reflect"
	"sync/atomic"
)

// Ensure that unused imports do not error
var (
	_ atomic.Uint64

	_ = io.ErrUnexpectedEOF
)

const (
	canoto__Spec__Name__tag   = "\x0a" // canoto.Tag(1, canoto.Len)
	canoto__Spec__Fields__tag = "\x12" // canoto.Tag(2, canoto.Len)
)

type canotoData_Spec struct {
	size atomic.Uint64
}

// CanotoSpec returns the specification of this canoto message.
func (*Spec) CanotoSpec(types ...reflect.Type) *Spec {
	types = append(types, reflect.TypeOf(Spec{}))
	var zero Spec
	s := &Spec{
		Name: "Spec",
		Fields: []FieldType{
			{
				FieldNumber: 1,
				Name:        "Name",
				OneOf:       "",
				TypeString:  true,
			},
			FieldTypeFromField(
				/*type inference:*/ (MakeEntryNilPointer(zero.Fields)),
				/*FieldNumber:   */ 2,
				/*Name:          */ "Fields",
				/*FixedLength:   */ 0,
				/*Repeated:      */ true,
				/*OneOf:         */ "",
				/*types:         */ types,
			),
		},
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
	r := Reader{
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
func (c *Spec) UnmarshalCanotoFrom(r Reader) error {
	// Zero the struct before unmarshaling.
	*c = Spec{}
	c.canotoData.size.Store(uint64(len(r.B)))

	var minField uint32
	for HasNext(&r) {
		field, wireType, err := ReadTag(&r)
		if err != nil {
			return err
		}
		if field < minField {
			return ErrInvalidFieldOrder
		}

		switch field {
		case 1:
			if wireType != Len {
				return ErrUnexpectedWireType
			}

			if err := ReadString(&r, &c.Name); err != nil {
				return err
			}
			if len(c.Name) == 0 {
				return ErrZeroValue
			}
		case 2:
			if wireType != Len {
				return ErrUnexpectedWireType
			}

			// Read the first entry manually because the tag is already
			// stripped.
			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			if err := ReadBytes(&r, &msgBytes); err != nil {
				return err
			}
			r.Unsafe = originalUnsafe

			// Count the number of additional entries after the first entry.
			countMinus1, err := CountBytes(r.B, canoto__Spec__Fields__tag)
			if err != nil {
				return err
			}

			c.Fields = MakeSlice(c.Fields, countMinus1+1)
			field := c.Fields
			additionalField := field[1:]
			if len(msgBytes) != 0 {
				remainingBytes := r.B
				r.B = msgBytes
				if err := (&field[0]).UnmarshalCanotoFrom(r); err != nil {
					return err
				}
				r.B = remainingBytes
			}

			// Read the rest of the entries, stripping the tag each time.
			for i := range additionalField {
				r.B = r.B[len(canoto__Spec__Fields__tag):]
				r.Unsafe = true
				if err := ReadBytes(&r, &msgBytes); err != nil {
					return err
				}
				if len(msgBytes) == 0 {
					continue
				}
				r.Unsafe = originalUnsafe

				remainingBytes := r.B
				r.B = msgBytes
				if err := (&additionalField[i]).UnmarshalCanotoFrom(r); err != nil {
					return err
				}
				r.B = remainingBytes
			}
		default:
			return ErrUnknownField
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
	if !ValidString(c.Name) {
		return false
	}
	{
		field := c.Fields
		for i := range field {
			if !(&field[i]).ValidCanoto() {
				return false
			}
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
	var size uint64
	if len(c.Name) != 0 {
		size += uint64(len(canoto__Spec__Name__tag)) + SizeBytes(c.Name)
	}
	{
		field := c.Fields
		for i := range field {
			(&field[i]).CalculateCanotoCache()
			fieldSize := (&field[i]).CachedCanotoSize()
			size += uint64(len(canoto__Spec__Fields__tag)) + SizeUint(fieldSize) + fieldSize
		}
	}
	c.canotoData.size.Store(size)
}

// CachedCanotoSize returns the previously calculated size of the Canoto
// representation from CalculateCanotoCache.
//
// If CalculateCanotoCache has not yet been called, it will return 0.
//
// If the struct has been modified since the last call to CalculateCanotoCache,
// the returned size may be incorrect.
func (c *Spec) CachedCanotoSize() uint64 {
	if c == nil {
		return 0
	}
	return c.canotoData.size.Load()
}

// MarshalCanoto returns the Canoto representation of this struct.
//
// It is assumed that this struct is ValidCanoto.
func (c *Spec) MarshalCanoto() []byte {
	c.CalculateCanotoCache()
	w := Writer{
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
func (c *Spec) MarshalCanotoInto(w Writer) Writer {
	if c == nil {
		return w
	}
	if len(c.Name) != 0 {
		Append(&w, canoto__Spec__Name__tag)
		AppendBytes(&w, c.Name)
	}
	{
		field := c.Fields
		for i := range field {
			Append(&w, canoto__Spec__Fields__tag)
			AppendUint(&w, (&field[i]).CachedCanotoSize())
			w = (&field[i]).MarshalCanotoInto(w)
		}
	}
	return w
}

const (
	canoto__FieldType__FieldNumber__tag    = "\x08" // canoto.Tag(1, canoto.Varint)
	canoto__FieldType__Name__tag           = "\x12" // canoto.Tag(2, canoto.Len)
	canoto__FieldType__FixedLength__tag    = "\x18" // canoto.Tag(3, canoto.Varint)
	canoto__FieldType__Repeated__tag       = "\x20" // canoto.Tag(4, canoto.Varint)
	canoto__FieldType__OneOf__tag          = "\x2a" // canoto.Tag(5, canoto.Len)
	canoto__FieldType__TypeInt__tag        = "\x30" // canoto.Tag(6, canoto.Varint)
	canoto__FieldType__TypeUint__tag       = "\x38" // canoto.Tag(7, canoto.Varint)
	canoto__FieldType__TypeFixedInt__tag   = "\x40" // canoto.Tag(8, canoto.Varint)
	canoto__FieldType__TypeFixedUint__tag  = "\x48" // canoto.Tag(9, canoto.Varint)
	canoto__FieldType__TypeBool__tag       = "\x50" // canoto.Tag(10, canoto.Varint)
	canoto__FieldType__TypeString__tag     = "\x58" // canoto.Tag(11, canoto.Varint)
	canoto__FieldType__TypeBytes__tag      = "\x60" // canoto.Tag(12, canoto.Varint)
	canoto__FieldType__TypeFixedBytes__tag = "\x68" // canoto.Tag(13, canoto.Varint)
	canoto__FieldType__TypeRecursive__tag  = "\x70" // canoto.Tag(14, canoto.Varint)
	canoto__FieldType__TypeMessage__tag    = "\x7a" // canoto.Tag(15, canoto.Len)
)

type canotoData_FieldType struct {
	size atomic.Uint64

	TypeOneOf atomic.Uint32
}

// CanotoSpec returns the specification of this canoto message.
func (*FieldType) CanotoSpec(types ...reflect.Type) *Spec {
	types = append(types, reflect.TypeOf(FieldType{}))
	var zero FieldType
	s := &Spec{
		Name: "FieldType",
		Fields: []FieldType{
			{
				FieldNumber: 1,
				Name:        "FieldNumber",
				OneOf:       "",
				TypeUint:    SizeOf(zero.FieldNumber),
			},
			{
				FieldNumber: 2,
				Name:        "Name",
				OneOf:       "",
				TypeString:  true,
			},
			{
				FieldNumber: 3,
				Name:        "FixedLength",
				OneOf:       "",
				TypeUint:    SizeOf(zero.FixedLength),
			},
			{
				FieldNumber: 4,
				Name:        "Repeated",
				OneOf:       "",
				TypeBool:    true,
			},
			{
				FieldNumber: 5,
				Name:        "OneOf",
				OneOf:       "",
				TypeString:  true,
			},
			{
				FieldNumber: 6,
				Name:        "TypeInt",
				OneOf:       "Type",
				TypeUint:    SizeOf(zero.TypeInt),
			},
			{
				FieldNumber: 7,
				Name:        "TypeUint",
				OneOf:       "Type",
				TypeUint:    SizeOf(zero.TypeUint),
			},
			{
				FieldNumber: 8,
				Name:        "TypeFixedInt",
				OneOf:       "Type",
				TypeUint:    SizeOf(zero.TypeFixedInt),
			},
			{
				FieldNumber: 9,
				Name:        "TypeFixedUint",
				OneOf:       "Type",
				TypeUint:    SizeOf(zero.TypeFixedUint),
			},
			{
				FieldNumber: 10,
				Name:        "TypeBool",
				OneOf:       "Type",
				TypeBool:    true,
			},
			{
				FieldNumber: 11,
				Name:        "TypeString",
				OneOf:       "Type",
				TypeBool:    true,
			},
			{
				FieldNumber: 12,
				Name:        "TypeBytes",
				OneOf:       "Type",
				TypeBool:    true,
			},
			{
				FieldNumber: 13,
				Name:        "TypeFixedBytes",
				OneOf:       "Type",
				TypeUint:    SizeOf(zero.TypeFixedBytes),
			},
			{
				FieldNumber: 14,
				Name:        "TypeRecursive",
				OneOf:       "Type",
				TypeUint:    SizeOf(zero.TypeRecursive),
			},
			FieldTypeFromField(
				/*type inference:*/ (zero.TypeMessage),
				/*FieldNumber:   */ 15,
				/*Name:          */ "TypeMessage",
				/*FixedLength:   */ 0,
				/*Repeated:      */ false,
				/*OneOf:         */ "Type",
				/*types:         */ types,
			),
		},
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
	r := Reader{
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
func (c *FieldType) UnmarshalCanotoFrom(r Reader) error {
	// Zero the struct before unmarshaling.
	*c = FieldType{}
	c.canotoData.size.Store(uint64(len(r.B)))

	var minField uint32
	for HasNext(&r) {
		field, wireType, err := ReadTag(&r)
		if err != nil {
			return err
		}
		if field < minField {
			return ErrInvalidFieldOrder
		}

		switch field {
		case 1:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}

			if err := ReadUint(&r, &c.FieldNumber); err != nil {
				return err
			}
			if IsZero(c.FieldNumber) {
				return ErrZeroValue
			}
		case 2:
			if wireType != Len {
				return ErrUnexpectedWireType
			}

			if err := ReadString(&r, &c.Name); err != nil {
				return err
			}
			if len(c.Name) == 0 {
				return ErrZeroValue
			}
		case 3:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}

			if err := ReadUint(&r, &c.FixedLength); err != nil {
				return err
			}
			if IsZero(c.FixedLength) {
				return ErrZeroValue
			}
		case 4:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}

			if err := ReadBool(&r, &c.Repeated); err != nil {
				return err
			}
			if IsZero(c.Repeated) {
				return ErrZeroValue
			}
		case 5:
			if wireType != Len {
				return ErrUnexpectedWireType
			}

			if err := ReadString(&r, &c.OneOf); err != nil {
				return err
			}
			if len(c.OneOf) == 0 {
				return ErrZeroValue
			}
		case 6:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(6) != 0 {
				return ErrDuplicateOneOf
			}

			if err := ReadUint(&r, &c.TypeInt); err != nil {
				return err
			}
			if IsZero(c.TypeInt) {
				return ErrZeroValue
			}
		case 7:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(7) != 0 {
				return ErrDuplicateOneOf
			}

			if err := ReadUint(&r, &c.TypeUint); err != nil {
				return err
			}
			if IsZero(c.TypeUint) {
				return ErrZeroValue
			}
		case 8:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(8) != 0 {
				return ErrDuplicateOneOf
			}

			if err := ReadUint(&r, &c.TypeFixedInt); err != nil {
				return err
			}
			if IsZero(c.TypeFixedInt) {
				return ErrZeroValue
			}
		case 9:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(9) != 0 {
				return ErrDuplicateOneOf
			}

			if err := ReadUint(&r, &c.TypeFixedUint); err != nil {
				return err
			}
			if IsZero(c.TypeFixedUint) {
				return ErrZeroValue
			}
		case 10:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(10) != 0 {
				return ErrDuplicateOneOf
			}

			if err := ReadBool(&r, &c.TypeBool); err != nil {
				return err
			}
			if IsZero(c.TypeBool) {
				return ErrZeroValue
			}
		case 11:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(11) != 0 {
				return ErrDuplicateOneOf
			}

			if err := ReadBool(&r, &c.TypeString); err != nil {
				return err
			}
			if IsZero(c.TypeString) {
				return ErrZeroValue
			}
		case 12:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(12) != 0 {
				return ErrDuplicateOneOf
			}

			if err := ReadBool(&r, &c.TypeBytes); err != nil {
				return err
			}
			if IsZero(c.TypeBytes) {
				return ErrZeroValue
			}
		case 13:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(13) != 0 {
				return ErrDuplicateOneOf
			}

			if err := ReadUint(&r, &c.TypeFixedBytes); err != nil {
				return err
			}
			if IsZero(c.TypeFixedBytes) {
				return ErrZeroValue
			}
		case 14:
			if wireType != Varint {
				return ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(14) != 0 {
				return ErrDuplicateOneOf
			}

			if err := ReadUint(&r, &c.TypeRecursive); err != nil {
				return err
			}
			if IsZero(c.TypeRecursive) {
				return ErrZeroValue
			}
		case 15:
			if wireType != Len {
				return ErrUnexpectedWireType
			}
			if c.canotoData.TypeOneOf.Swap(15) != 0 {
				return ErrDuplicateOneOf
			}

			// Read the bytes for the field.
			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			if err := ReadBytes(&r, &msgBytes); err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return ErrZeroValue
			}
			r.Unsafe = originalUnsafe

			// Unmarshal the field from the bytes.
			remainingBytes := r.B
			r.B = msgBytes
			c.TypeMessage = MakePointer(c.TypeMessage)
			if err := (c.TypeMessage).UnmarshalCanotoFrom(r); err != nil {
				return err
			}
			r.B = remainingBytes
		default:
			return ErrUnknownField
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
	var TypeOneOf uint32
	if !IsZero(c.TypeInt) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 6
	}
	if !IsZero(c.TypeUint) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 7
	}
	if !IsZero(c.TypeFixedInt) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 8
	}
	if !IsZero(c.TypeFixedUint) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 9
	}
	if !IsZero(c.TypeBool) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 10
	}
	if !IsZero(c.TypeString) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 11
	}
	if !IsZero(c.TypeBytes) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 12
	}
	if !IsZero(c.TypeFixedBytes) {
		if TypeOneOf != 0 {
			return false
		}
		TypeOneOf = 13
	}
	if !IsZero(c.TypeRecursive) {
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
	if !ValidString(c.Name) {
		return false
	}
	if !ValidString(c.OneOf) {
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
	var size uint64
	var TypeOneOf uint32
	if !IsZero(c.FieldNumber) {
		size += uint64(len(canoto__FieldType__FieldNumber__tag)) + SizeUint(c.FieldNumber)
	}
	if len(c.Name) != 0 {
		size += uint64(len(canoto__FieldType__Name__tag)) + SizeBytes(c.Name)
	}
	if !IsZero(c.FixedLength) {
		size += uint64(len(canoto__FieldType__FixedLength__tag)) + SizeUint(c.FixedLength)
	}
	if !IsZero(c.Repeated) {
		size += uint64(len(canoto__FieldType__Repeated__tag)) + SizeBool
	}
	if len(c.OneOf) != 0 {
		size += uint64(len(canoto__FieldType__OneOf__tag)) + SizeBytes(c.OneOf)
	}
	if !IsZero(c.TypeInt) {
		size += uint64(len(canoto__FieldType__TypeInt__tag)) + SizeUint(c.TypeInt)
		TypeOneOf = 6
	}
	if !IsZero(c.TypeUint) {
		size += uint64(len(canoto__FieldType__TypeUint__tag)) + SizeUint(c.TypeUint)
		TypeOneOf = 7
	}
	if !IsZero(c.TypeFixedInt) {
		size += uint64(len(canoto__FieldType__TypeFixedInt__tag)) + SizeUint(c.TypeFixedInt)
		TypeOneOf = 8
	}
	if !IsZero(c.TypeFixedUint) {
		size += uint64(len(canoto__FieldType__TypeFixedUint__tag)) + SizeUint(c.TypeFixedUint)
		TypeOneOf = 9
	}
	if !IsZero(c.TypeBool) {
		size += uint64(len(canoto__FieldType__TypeBool__tag)) + SizeBool
		TypeOneOf = 10
	}
	if !IsZero(c.TypeString) {
		size += uint64(len(canoto__FieldType__TypeString__tag)) + SizeBool
		TypeOneOf = 11
	}
	if !IsZero(c.TypeBytes) {
		size += uint64(len(canoto__FieldType__TypeBytes__tag)) + SizeBool
		TypeOneOf = 12
	}
	if !IsZero(c.TypeFixedBytes) {
		size += uint64(len(canoto__FieldType__TypeFixedBytes__tag)) + SizeUint(c.TypeFixedBytes)
		TypeOneOf = 13
	}
	if !IsZero(c.TypeRecursive) {
		size += uint64(len(canoto__FieldType__TypeRecursive__tag)) + SizeUint(c.TypeRecursive)
		TypeOneOf = 14
	}
	if c.TypeMessage != nil {
		(c.TypeMessage).CalculateCanotoCache()
		if fieldSize := (c.TypeMessage).CachedCanotoSize(); fieldSize != 0 {
			size += uint64(len(canoto__FieldType__TypeMessage__tag)) + SizeUint(fieldSize) + fieldSize
			TypeOneOf = 15
		}
	}
	c.canotoData.size.Store(size)
	c.canotoData.TypeOneOf.Store(TypeOneOf)
}

// CachedCanotoSize returns the previously calculated size of the Canoto
// representation from CalculateCanotoCache.
//
// If CalculateCanotoCache has not yet been called, it will return 0.
//
// If the struct has been modified since the last call to CalculateCanotoCache,
// the returned size may be incorrect.
func (c *FieldType) CachedCanotoSize() uint64 {
	if c == nil {
		return 0
	}
	return c.canotoData.size.Load()
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
	w := Writer{
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
func (c *FieldType) MarshalCanotoInto(w Writer) Writer {
	if c == nil {
		return w
	}
	if !IsZero(c.FieldNumber) {
		Append(&w, canoto__FieldType__FieldNumber__tag)
		AppendUint(&w, c.FieldNumber)
	}
	if len(c.Name) != 0 {
		Append(&w, canoto__FieldType__Name__tag)
		AppendBytes(&w, c.Name)
	}
	if !IsZero(c.FixedLength) {
		Append(&w, canoto__FieldType__FixedLength__tag)
		AppendUint(&w, c.FixedLength)
	}
	if !IsZero(c.Repeated) {
		Append(&w, canoto__FieldType__Repeated__tag)
		AppendBool(&w, true)
	}
	if len(c.OneOf) != 0 {
		Append(&w, canoto__FieldType__OneOf__tag)
		AppendBytes(&w, c.OneOf)
	}
	if !IsZero(c.TypeInt) {
		Append(&w, canoto__FieldType__TypeInt__tag)
		AppendUint(&w, c.TypeInt)
	}
	if !IsZero(c.TypeUint) {
		Append(&w, canoto__FieldType__TypeUint__tag)
		AppendUint(&w, c.TypeUint)
	}
	if !IsZero(c.TypeFixedInt) {
		Append(&w, canoto__FieldType__TypeFixedInt__tag)
		AppendUint(&w, c.TypeFixedInt)
	}
	if !IsZero(c.TypeFixedUint) {
		Append(&w, canoto__FieldType__TypeFixedUint__tag)
		AppendUint(&w, c.TypeFixedUint)
	}
	if !IsZero(c.TypeBool) {
		Append(&w, canoto__FieldType__TypeBool__tag)
		AppendBool(&w, true)
	}
	if !IsZero(c.TypeString) {
		Append(&w, canoto__FieldType__TypeString__tag)
		AppendBool(&w, true)
	}
	if !IsZero(c.TypeBytes) {
		Append(&w, canoto__FieldType__TypeBytes__tag)
		AppendBool(&w, true)
	}
	if !IsZero(c.TypeFixedBytes) {
		Append(&w, canoto__FieldType__TypeFixedBytes__tag)
		AppendUint(&w, c.TypeFixedBytes)
	}
	if !IsZero(c.TypeRecursive) {
		Append(&w, canoto__FieldType__TypeRecursive__tag)
		AppendUint(&w, c.TypeRecursive)
	}
	if c.TypeMessage != nil {
		if fieldSize := (c.TypeMessage).CachedCanotoSize(); fieldSize != 0 {
			Append(&w, canoto__FieldType__TypeMessage__tag)
			AppendUint(&w, fieldSize)
			w = (c.TypeMessage).MarshalCanotoInto(w)
		}
	}
	return w
}
