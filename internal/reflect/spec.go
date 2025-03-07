//go:generate canoto $GOFILE

package reflect

import (
	"errors"
	"fmt"
	"slices"

	"github.com/StephenButtolph/canoto"
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
		TypeInt        uint8  `canoto:"int,5,Type"      json:"typeInt,omitempty"`        // can be any of 8, 16, 32, or 64.
		TypeUint       uint8  `canoto:"int,6,Type"      json:"typeUint,omitempty"`       // can be any of 8, 16, 32, or 64.
		TypeSint       uint8  `canoto:"int,7,Type"      json:"typeSint,omitempty"`       // can be any of 8, 16, 32, or 64.
		TypeFint       uint8  `canoto:"int,8,Type"      json:"typeFint,omitempty"`       // can be either 32 or 64.
		TypeSFint      uint8  `canoto:"int,9,Type"      json:"typeSFint,omitempty"`      // can be either 32 or 64.
		TypeBool       bool   `canoto:"bool,10,Type"    json:"typeBool,omitempty"`       // can only be true.
		TypeString     bool   `canoto:"bool,11,Type"    json:"typeString,omitempty"`     // can only be true.
		TypeBytes      bool   `canoto:"bool,12,Type"    json:"typeBytes,omitempty"`      // can only be true.
		TypeFixedBytes uint64 `canoto:"int,13,Type"     json:"typeFixedBytes,omitempty"` // length of the fixed bytes.
		TypeRecursive  uint64 `canoto:"int,14,Type"     json:"typeRecursive,omitempty"`  // depth of the recursion.
		TypeMessage    *Spec  `canoto:"pointer,15,Type" json:"typeMessage,omitempty"`

		canotoData canotoData_FieldType
	}
	unmarshaler func(f *FieldType, r *canoto.Reader, specs []*Spec) (any, error)
	Any         map[string]any
)

func (s *Spec) Unmarshal(b []byte) (Any, error) {
	s.CalculateCanotoCache()
	r := canoto.Reader{
		B: b,
	}
	return s.unmarshal(&r, nil)
}

func (s *Spec) unmarshal(r *canoto.Reader, specs []*Spec) (Any, error) {
	specs = append(specs, s)
	var (
		minField uint32
		a        = make(Any)
	)
	for canoto.HasNext(r) {
		fieldNumber, wireType, err := canoto.ReadTag(r)
		if err != nil {
			return Any{}, fmt.Errorf("reading tag: %w", err)
		}
		if fieldNumber < minField {
			return Any{}, fmt.Errorf("fieldNumber %d < minField %d: %w", fieldNumber, minField, canoto.ErrInvalidFieldOrder)
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
			return Any{}, fmt.Errorf("fieldNumber %d: %w", fieldNumber, canoto.ErrInvalidWireType)
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
	return nil, canoto.ErrUnknownField
}

func (f *FieldType) wireType() (canoto.WireType, error) {
	whichOneOf := f.CachedWhichOneOfType()
	switch whichOneOf {
	case 5, 6, 7, 10:
		if f.Repeated {
			return canoto.Len, nil
		}
		return canoto.Varint, nil
	case 8:
		switch f.TypeFint {
		case 32:
			return canoto.I32, nil
		case 64:
			return canoto.I64, nil
		default:
			return 0, canoto.ErrUnknownField
		}
	case 9:
		switch f.TypeSFint {
		case 32:
			return canoto.I32, nil
		case 64:
			return canoto.I64, nil
		default:
			return 0, canoto.ErrUnknownField
		}
	case 11, 12, 13, 14, 15:
		return canoto.Len, nil
	default:
		return 0, canoto.ErrUnknownField
	}
}

func (f *FieldType) unmarshal(r *canoto.Reader, specs []*Spec) (any, error) {
	whichOneOf := f.CachedWhichOneOfType()
	unmarshal, ok := map[uint32]unmarshaler{
		5:  (*FieldType).unmarshalInt,
		6:  (*FieldType).unmarshalUint,
		7:  (*FieldType).unmarshalSint,
		8:  (*FieldType).unmarshalFint,
		9:  (*FieldType).unmarshalSFint,
		10: (*FieldType).unmarshalBool,
		11: (*FieldType).unmarshalString,
		12: (*FieldType).unmarshalBytes,
		13: (*FieldType).unmarshalFixedBytes,
		14: (*FieldType).unmarshalRecursive,
		15: (*FieldType).unmarshalSpec,
	}[whichOneOf]
	if !ok {
		return nil, canoto.ErrUnknownField
	}
	value, err := unmarshal(f, r, specs)
	if err != nil {
		return nil, fmt.Errorf("%d: %w", whichOneOf, err)
	}
	return value, nil
}

func (f *FieldType) unmarshalInt(r *canoto.Reader, _ []*Spec) (any, error) {
	switch f.TypeInt {
	case 8:
		var v int8
		err := canoto.ReadInt(r, &v)
		return v, err
	case 16:
		var v int16
		err := canoto.ReadInt(r, &v)
		return v, err
	case 32:
		var v int32
		err := canoto.ReadInt(r, &v)
		return v, err
	case 64:
		var v int64
		err := canoto.ReadInt(r, &v)
		return v, err
	default:
		return nil, canoto.ErrUnknownField
	}
}

func (f *FieldType) unmarshalUint(r *canoto.Reader, _ []*Spec) (any, error) {
	switch f.TypeUint {
	case 8:
		var v uint8
		err := canoto.ReadInt(r, &v)
		return v, err
	case 16:
		var v uint16
		err := canoto.ReadInt(r, &v)
		return v, err
	case 32:
		var v uint32
		err := canoto.ReadInt(r, &v)
		return v, err
	case 64:
		var v uint64
		err := canoto.ReadInt(r, &v)
		return v, err
	default:
		return nil, canoto.ErrUnknownField
	}
}

func (f *FieldType) unmarshalSint(r *canoto.Reader, _ []*Spec) (any, error) {
	switch f.TypeSint {
	case 8:
		var v int8
		err := canoto.ReadSint(r, &v)
		return v, err
	case 16:
		var v int16
		err := canoto.ReadSint(r, &v)
		return v, err
	case 32:
		var v int32
		err := canoto.ReadSint(r, &v)
		return v, err
	case 64:
		var v int64
		err := canoto.ReadSint(r, &v)
		return v, err
	default:
		return nil, canoto.ErrUnknownField
	}
}

func (f *FieldType) unmarshalFint(r *canoto.Reader, _ []*Spec) (any, error) {
	switch f.TypeFint {
	case 32:
		var v uint32
		err := canoto.ReadFint32(r, &v)
		return v, err
	case 64:
		var v uint64
		err := canoto.ReadFint64(r, &v)
		return v, err
	default:
		return nil, canoto.ErrUnknownField
	}
}

func (f *FieldType) unmarshalSFint(r *canoto.Reader, _ []*Spec) (any, error) {
	switch f.TypeSFint {
	case 32:
		var v int32
		err := canoto.ReadFint32(r, &v)
		return v, err
	case 64:
		var v int64
		err := canoto.ReadFint64(r, &v)
		return v, err
	default:
		return nil, canoto.ErrUnknownField
	}
}

func (f *FieldType) unmarshalBool(r *canoto.Reader, _ []*Spec) (any, error) {
	var v bool
	err := canoto.ReadBool(r, &v)
	return v, err
}

func (f *FieldType) unmarshalString(r *canoto.Reader, _ []*Spec) (any, error) {
	var v string
	err := canoto.ReadString(r, &v)
	return v, err
}

func (f *FieldType) unmarshalBytes(r *canoto.Reader, _ []*Spec) (any, error) {
	var v []byte
	err := canoto.ReadBytes(r, &v)
	return v, err
}

func (f *FieldType) unmarshalFixedBytes(r *canoto.Reader, _ []*Spec) (any, error) {
	var v []byte
	if err := canoto.ReadBytes(r, &v); err != nil {
		return nil, err
	}
	if uint64(len(v)) != f.TypeFixedBytes {
		return nil, canoto.ErrInvalidLength
	}
	return v, nil
}

func (f *FieldType) unmarshalRecursive(r *canoto.Reader, specs []*Spec) (any, error) {
	var b []byte
	if err := canoto.ReadBytes(r, &b); err != nil {
		return nil, err
	}

	r = &canoto.Reader{
		B: b,
	}
	numSpecs := uint64(len(specs))
	if f.TypeRecursive > numSpecs {
		return nil, errors.New("invalid depth")
	}
	index := numSpecs - f.TypeRecursive
	spec := specs[index]
	specs = slices.Clone(specs[:index])
	return spec.unmarshal(r, specs)
}

func (f *FieldType) unmarshalSpec(r *canoto.Reader, specs []*Spec) (any, error) {
	// Read the first entry manually because the tag is already stripped.
	var msgBytes []byte
	if err := canoto.ReadBytes(r, &msgBytes); err != nil {
		return nil, err
	}
	msg, err := f.TypeMessage.unmarshal(
		&canoto.Reader{
			B: msgBytes,
		},
		specs,
	)
	if err != nil {
		return nil, err
	}
	if !f.Repeated {
		// If there is only one entry, return it.
		return msg, nil
	}

	// Count the number of additional entries after the first entry.
	wireType, err := f.wireType()
	if err != nil {
		return nil, err
	}
	expectedTag := canoto.Tag(f.FieldNumber, wireType)
	countMinus1, err := canoto.CountBytes(r.B, string(expectedTag))
	if err != nil {
		return nil, err
	}

	// If there should be a specific number of entries, check that the count is
	// correct.
	totalCount := countMinus1 + 1
	if f.FixedLength > 0 && uint64(totalCount) != f.FixedLength {
		return nil, canoto.ErrInvalidLength
	}

	msgs := make([]any, totalCount)
	msgs[0] = msg

	// Read the rest of the entries, stripping the tag each time.
	for i := range countMinus1 {
		r.B = r.B[len(expectedTag):]
		if err := canoto.ReadBytes(r, &msgBytes); err != nil {
			return nil, err
		}
		if len(msgBytes) == 0 {
			continue
		}

		msgs[1+i], err = f.TypeMessage.unmarshal(
			&canoto.Reader{
				B: msgBytes,
			},
			specs,
		)
		if err != nil {
			return nil, err
		}
	}
	return msgs, nil
}
