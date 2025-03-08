//go:generate canoto --internal $GOFILE

package canoto

import (
	"errors"
	"fmt"
	"io"
	"slices"
)

type (
	Spec struct {
		Name   string       `canoto:"string,1"           json:"name"`
		Fields []*FieldType `canoto:"repeated pointer,2" json:"fields"` // TODO: Replace this with a map.

		canotoData canotoData_Spec
	}
	FieldType struct {
		FieldNumber    uint32 `canoto:"int,1"           json:"fieldNumber"`
		Name           string `canoto:"string,2"        json:"name"`
		FixedLength    uint64 `canoto:"int,3"           json:"fixedLength,omitempty"`
		Repeated       bool   `canoto:"bool,4"          json:"repeated,omitempty"`
		OneOf          string `canoto:"string,5"        json:"oneOf,omitempty"`
		TypeInt        uint8  `canoto:"int,6,Type"      json:"typeInt,omitempty"`        // can be any of 8, 16, 32, or 64.
		TypeUint       uint8  `canoto:"int,7,Type"      json:"typeUint,omitempty"`       // can be any of 8, 16, 32, or 64.
		TypeSint       uint8  `canoto:"int,8,Type"      json:"typeSint,omitempty"`       // can be any of 8, 16, 32, or 64.
		TypeFint       uint8  `canoto:"int,9,Type"      json:"typeFint,omitempty"`       // can be either 32 or 64.
		TypeSFint      uint8  `canoto:"int,10,Type"     json:"typeSFint,omitempty"`      // can be either 32 or 64.
		TypeBool       bool   `canoto:"bool,11,Type"    json:"typeBool,omitempty"`       // can only be true.
		TypeString     bool   `canoto:"bool,12,Type"    json:"typeString,omitempty"`     // can only be true.
		TypeBytes      bool   `canoto:"bool,13,Type"    json:"typeBytes,omitempty"`      // can only be true.
		TypeFixedBytes uint64 `canoto:"int,14,Type"     json:"typeFixedBytes,omitempty"` // length of the fixed bytes.
		TypeRecursive  uint64 `canoto:"int,15,Type"     json:"typeRecursive,omitempty"`  // depth of the recursion.
		TypeMessage    *Spec  `canoto:"pointer,16,Type" json:"typeMessage,omitempty"`

		canotoData canotoData_FieldType
	}
	unmarshaler func(f *FieldType, r *Reader, specs []*Spec) (any, error)
	Any         map[string]any
)

// Unmarshal unmarshals the given bytes into a map of fields based on the
// specification.
func Unmarshal(s *Spec, b []byte) (Any, error) {
	r := Reader{
		B: b,
	}
	return s.unmarshal(&r, nil)
}

func (s *Spec) unmarshal(r *Reader, specs []*Spec) (Any, error) {
	specs = append(specs, s)
	var (
		minField uint32
		a        = make(Any)
		oneOfs   = make(map[string]struct{})
	)
	for HasNext(r) {
		fieldNumber, wireType, err := ReadTag(r)
		if err != nil {
			return Any{}, fmt.Errorf("reading tag: %w", err)
		}
		if fieldNumber < minField {
			return Any{}, fmt.Errorf("fieldNumber %d < minField %d: %w", fieldNumber, minField, ErrInvalidFieldOrder)
		}

		fieldType, err := s.findField(fieldNumber)
		if err != nil {
			return Any{}, fmt.Errorf("find field %d: %w", fieldNumber, err)
		}

		expectedWireType, err := fieldType.wireType()
		if err != nil {
			return Any{}, fmt.Errorf("wireType for %d: %w", fieldNumber, err)
		}
		if wireType != expectedWireType {
			return Any{}, fmt.Errorf("fieldNumber %d: %w", fieldNumber, ErrUnexpectedWireType)
		}

		if fieldType.OneOf != "" {
			if _, ok := oneOfs[fieldType.OneOf]; ok {
				return Any{}, fmt.Errorf("fieldNumber %d: %w", fieldNumber, ErrDuplicateOneOf)
			}
			oneOfs[fieldType.OneOf] = struct{}{}
		}

		value, err := fieldType.unmarshal(r, specs)
		if err != nil {
			return Any{}, fmt.Errorf("unmarshal fieldNumber %d: %w", fieldNumber, err)
		}
		a[fieldType.Name] = value

		minField = fieldNumber + 1
	}
	return a, nil
}

func (s *Spec) findField(fieldNumber uint32) (*FieldType, error) {
	for _, f := range s.Fields {
		if f.FieldNumber == fieldNumber {
			return f, nil
		}
	}
	return nil, ErrUnknownField
}

func (f *FieldType) wireType() (WireType, error) {
	whichOneOf := f.CachedWhichOneOfType()
	switch whichOneOf {
	case 6, 7, 8, 11:
		if f.Repeated {
			return Len, nil
		}
		return Varint, nil
	case 9:
		if f.Repeated {
			return Len, nil
		}
		switch f.TypeFint {
		case 3:
			return I32, nil
		case 4:
			return I64, nil
		default:
			return 0, ErrUnknownField
		}
	case 10:
		if f.Repeated {
			return Len, nil
		}
		switch f.TypeSFint {
		case 3:
			return I32, nil
		case 4:
			return I64, nil
		default:
			return 0, ErrUnknownField
		}
	case 12, 13, 14, 15, 16:
		return Len, nil
	default:
		return 0, ErrUnknownField
	}
}

func (f *FieldType) unmarshal(r *Reader, specs []*Spec) (any, error) {
	whichOneOf := f.CachedWhichOneOfType()
	unmarshal, ok := map[uint32]unmarshaler{
		6:  (*FieldType).unmarshalInt,
		7:  (*FieldType).unmarshalUint,
		8:  (*FieldType).unmarshalSint,
		9:  (*FieldType).unmarshalFint,
		10: (*FieldType).unmarshalSFint,
		11: (*FieldType).unmarshalBool,
		12: (*FieldType).unmarshalString,
		13: (*FieldType).unmarshalBytes,
		14: (*FieldType).unmarshalFixedBytes,
		15: (*FieldType).unmarshalRecursive,
		16: (*FieldType).unmarshalSpec,
	}[whichOneOf]
	if !ok {
		return nil, ErrUnknownField
	}
	value, err := unmarshal(f, r, specs)
	if err != nil {
		return nil, fmt.Errorf("%d: %w", whichOneOf, err)
	}
	return value, nil
}

func (f *FieldType) unmarshalInt(r *Reader, _ []*Spec) (any, error) {
	return unmarshalPackedVarint(
		f,
		r,
		func(r *Reader) (int64, error) {
			switch f.TypeInt {
			case 1:
				var v int8
				err := ReadInt(r, &v)
				return int64(v), err
			case 2:
				var v int16
				err := ReadInt(r, &v)
				return int64(v), err
			case 3:
				var v int32
				err := ReadInt(r, &v)
				return int64(v), err
			case 4:
				var v int64
				err := ReadInt(r, &v)
				return v, err
			default:
				return 0, ErrUnknownField
			}
		},
	)
}

func (f *FieldType) unmarshalUint(r *Reader, _ []*Spec) (any, error) {
	return unmarshalPackedVarint(
		f,
		r,
		func(r *Reader) (uint64, error) {
			switch f.TypeUint {
			case 1:
				var v uint8
				err := ReadInt(r, &v)
				return uint64(v), err
			case 2:
				var v uint16
				err := ReadInt(r, &v)
				return uint64(v), err
			case 3:
				var v uint32
				err := ReadInt(r, &v)
				return uint64(v), err
			case 4:
				var v uint64
				err := ReadInt(r, &v)
				return v, err
			default:
				return 0, ErrUnknownField
			}
		},
	)
}

func (f *FieldType) unmarshalSint(r *Reader, _ []*Spec) (any, error) {
	return unmarshalPackedVarint(
		f,
		r,
		func(r *Reader) (int64, error) {
			switch f.TypeSint {
			case 1:
				var v int8
				err := ReadSint(r, &v)
				return int64(v), err
			case 2:
				var v int16
				err := ReadSint(r, &v)
				return int64(v), err
			case 3:
				var v int32
				err := ReadSint(r, &v)
				return int64(v), err
			case 4:
				var v int64
				err := ReadSint(r, &v)
				return v, err
			default:
				return 0, ErrUnknownField
			}
		},
	)
}

func (f *FieldType) unmarshalFint(r *Reader, _ []*Spec) (any, error) {
	return unmarshalPackedFixed(
		f,
		r,
		func(r *Reader) (uint64, error) {
			switch f.TypeFint {
			case 3:
				var v uint32
				err := ReadFint32(r, &v)
				return uint64(v), err
			case 4:
				var v uint64
				err := ReadFint64(r, &v)
				return v, err
			default:
				return 0, ErrUnknownField
			}
		},
		f.TypeFint,
	)
}

func (f *FieldType) unmarshalSFint(r *Reader, _ []*Spec) (any, error) {
	return unmarshalPackedFixed(
		f,
		r,
		func(r *Reader) (int64, error) {
			switch f.TypeSFint {
			case 3:
				var v int32
				err := ReadFint32(r, &v)
				return int64(v), err
			case 4:
				var v int64
				err := ReadFint64(r, &v)
				return v, err
			default:
				return 0, ErrUnknownField
			}
		},
		f.TypeSFint,
	)
}

func (f *FieldType) unmarshalBool(r *Reader, _ []*Spec) (any, error) {
	return unmarshalPackedFixed(
		f,
		r,
		func(r *Reader) (bool, error) {
			var v bool
			err := ReadBool(r, &v)
			return v, err
		},
		1,
	)
}

func (f *FieldType) unmarshalString(r *Reader, _ []*Spec) (any, error) {
	return unmarshalUnpacked(
		f,
		r,
		func(r *Reader) (string, bool, error) {
			var s string
			if err := ReadString(r, &s); err != nil {
				return "", false, err
			}
			return s, len(s) == 0, nil
		},
	)
}

func (f *FieldType) unmarshalBytes(r *Reader, _ []*Spec) (any, error) {
	return unmarshalUnpacked(
		f,
		r,
		func(r *Reader) ([]byte, bool, error) {
			var msgBytes []byte
			err := ReadBytes(r, &msgBytes)
			return msgBytes, len(msgBytes) == 0, err
		},
	)
}

func (f *FieldType) unmarshalFixedBytes(r *Reader, _ []*Spec) (any, error) {
	return unmarshalUnpacked(
		f,
		r,
		func(r *Reader) ([]byte, bool, error) {
			var length uint64
			if err := ReadInt(r, &length); err != nil {
				return nil, false, err
			}
			if f.TypeFixedBytes > uint64(len(r.B)) {
				return nil, false, io.ErrUnexpectedEOF
			}
			if length != f.TypeFixedBytes {
				return nil, false, ErrInvalidLength
			}
			msgBytes := slices.Clone(r.B[:length])
			r.B = r.B[length:]
			return msgBytes, isBytesEmpty(msgBytes), nil
		},
	)
}

func (f *FieldType) unmarshalRecursive(r *Reader, specs []*Spec) (any, error) {
	numSpecs := uint64(len(specs))
	if f.TypeRecursive > numSpecs {
		return nil, errors.New("invalid depth")
	}
	index := numSpecs - f.TypeRecursive
	spec := specs[index]
	specs = slices.Clone(specs[:index])

	return unmarshalUnpacked(
		f,
		r,
		func(r *Reader) (Any, bool, error) {
			var msgBytes []byte
			if err := ReadBytes(r, &msgBytes); err != nil {
				return nil, false, err
			}
			if len(msgBytes) == 0 {
				return nil, true, nil
			}
			a, err := spec.unmarshal(
				&Reader{
					B: msgBytes,
				},
				specs,
			)
			return a, false, err
		},
	)
}

func (f *FieldType) unmarshalSpec(r *Reader, specs []*Spec) (any, error) {
	return unmarshalUnpacked(
		f,
		r,
		func(r *Reader) (Any, bool, error) {
			var msgBytes []byte
			if err := ReadBytes(r, &msgBytes); err != nil {
				return nil, false, err
			}
			if len(msgBytes) == 0 {
				return nil, true, nil
			}
			a, err := f.TypeMessage.unmarshal(
				&Reader{
					B: msgBytes,
				},
				specs,
			)
			return a, false, err
		},
	)
}

func unmarshalPackedVarint[T comparable](
	f *FieldType,
	r *Reader,
	unmarshal func(r *Reader) (T, error),
) (any, error) {
	if !f.Repeated {
		// If there is only one entry, read it.
		value, err := unmarshal(r)
		if err != nil {
			return nil, err
		}
		if IsZero(value) {
			return nil, ErrZeroValue
		}
		return value, nil
	}

	// Read the full packed bytes.
	var msgBytes []byte
	if err := ReadBytes(r, &msgBytes); err != nil {
		return nil, err
	}

	count := f.FixedLength
	if count == 0 {
		if len(msgBytes) == 0 {
			return nil, ErrZeroValue
		}

		count = uint64(CountInts(msgBytes))
	}
	values := make([]T, count)
	r = &Reader{
		B: msgBytes,
	}
	isZero := true
	for i := range values {
		value, err := unmarshal(r)
		if err != nil {
			return nil, err
		}
		values[i] = value
		isZero = isZero && IsZero(value)
	}
	if HasNext(r) {
		return nil, ErrInvalidLength
	}
	if f.FixedLength > 0 && isZero {
		return nil, ErrZeroValue
	}
	return values, nil
}

func unmarshalPackedFixed[T comparable](
	f *FieldType,
	r *Reader,
	unmarshal func(r *Reader) (T, error),
	sizeEnum uint8,
) (any, error) {
	if !f.Repeated {
		// If there is only one entry, read it.
		value, err := unmarshal(r)
		if err != nil {
			return nil, err
		}
		if IsZero(value) {
			return nil, ErrZeroValue
		}
		return value, nil
	}

	// Read the full packed bytes.
	var msgBytes []byte
	if err := ReadBytes(r, &msgBytes); err != nil {
		return nil, err
	}

	count := f.FixedLength
	if count == 0 {
		numMsgBytes := uint64(len(msgBytes))
		if numMsgBytes == 0 {
			return nil, ErrZeroValue
		}

		var size uint64
		switch sizeEnum {
		case 1:
			size = SizeBool
		case 3:
			size = SizeFint32
		case 4:
			size = SizeFint64
		default:
			return 0, ErrUnknownField
		}

		if numMsgBytes%size != 0 {
			return nil, ErrInvalidLength
		}
		count = numMsgBytes / size
	}

	values := make([]T, count)
	r = &Reader{
		B: msgBytes,
	}
	isZero := true
	for i := range values {
		value, err := unmarshal(r)
		if err != nil {
			return nil, err
		}
		values[i] = value
		isZero = isZero && IsZero(value)
	}
	if HasNext(r) {
		return nil, ErrInvalidLength
	}
	if f.FixedLength > 0 && isZero {
		return nil, ErrZeroValue
	}
	return values, nil
}

func unmarshalUnpacked[T any](
	f *FieldType,
	r *Reader,
	unmarshal func(*Reader) (T, bool, error),
) (any, error) {
	// Read the first entry manually because the tag is already stripped.
	value, isZero, err := unmarshal(r)
	if err != nil {
		return nil, err
	}
	if !f.Repeated {
		// If there is only one entry, return it.
		if isZero {
			return nil, ErrZeroValue
		}
		return value, nil
	}

	// Count the number of additional entries after the first entry.
	expectedTag := Tag(f.FieldNumber, Len)

	count := f.FixedLength
	if count == 0 {
		countMinus1, err := CountBytes(r.B, string(expectedTag))
		if err != nil {
			return nil, err
		}
		count = uint64(countMinus1 + 1)
	}

	values := make([]T, count)
	values[0] = value

	// Read the rest of the entries, stripping the tag each time.
	for i := range count - 1 {
		if !HasPrefix(r.B, string(expectedTag)) {
			return nil, ErrUnknownField
		}
		r.B = r.B[len(expectedTag):]

		var isFieldZero bool
		values[1+i], isFieldZero, err = unmarshal(r)
		if err != nil {
			return nil, err
		}
		isZero = isZero && isFieldZero
	}
	if f.FixedLength > 0 && isZero {
		return nil, ErrZeroValue
	}
	return values, nil
}

// isBytesEmpty returns true if the byte slice is all zeros.
func isBytesEmpty(b []byte) bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}
