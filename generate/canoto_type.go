package generate

import (
	"errors"
	"slices"

	"github.com/StephenButtolph/canoto"
)

const (
	canotoInt        canotoType = "int"
	canotoUint       canotoType = "uint"
	canotoFint32     canotoType = "fint32"
	canotoFint64     canotoType = "fint64"
	canotoBool       canotoType = "bool"
	canotoString     canotoType = "string"
	canotoBytes      canotoType = "bytes"
	canotoFixedBytes canotoType = "fixed bytes"
	canotoValue      canotoType = "value"
	canotoPointer    canotoType = "pointer"
	canotoField      canotoType = "field"

	canotoRepeatedInt        = "repeated " + canotoInt
	canotoRepeatedUint       = "repeated " + canotoUint
	canotoRepeatedFint32     = "repeated " + canotoFint32
	canotoRepeatedFint64     = "repeated " + canotoFint64
	canotoRepeatedBool       = "repeated " + canotoBool
	canotoRepeatedString     = "repeated " + canotoString
	canotoRepeatedBytes      = "repeated " + canotoBytes
	canotoRepeatedFixedBytes = "repeated " + canotoFixedBytes
	canotoRepeatedValue      = "repeated " + canotoValue
	canotoRepeatedPointer    = "repeated " + canotoPointer
	canotoRepeatedField      = "repeated " + canotoField

	canotoFixedRepeatedInt        = "fixed " + canotoRepeatedInt
	canotoFixedRepeatedUint       = "fixed " + canotoRepeatedUint
	canotoFixedRepeatedFint32     = "fixed " + canotoRepeatedFint32
	canotoFixedRepeatedFint64     = "fixed " + canotoRepeatedFint64
	canotoFixedRepeatedBool       = "fixed " + canotoRepeatedBool
	canotoFixedRepeatedString     = "fixed " + canotoRepeatedString
	canotoFixedRepeatedBytes      = "fixed " + canotoRepeatedBytes
	canotoFixedRepeatedFixedBytes = "fixed " + canotoRepeatedFixedBytes
	canotoFixedRepeatedValue      = "fixed " + canotoRepeatedValue
	canotoFixedRepeatedPointer    = "fixed " + canotoRepeatedPointer
	canotoFixedRepeatedField      = "fixed " + canotoRepeatedField
)

var (
	canotoTypes = []canotoType{
		canotoInt,
		canotoUint,
		canotoFint32,
		canotoFint64,
		canotoBool,
		canotoString,
		canotoBytes,
		canotoFixedBytes,
		canotoValue,
		canotoPointer,
		canotoField,

		canotoRepeatedInt,
		canotoRepeatedUint,
		canotoRepeatedFint32,
		canotoRepeatedFint64,
		canotoRepeatedBool,
		canotoRepeatedString,
		canotoRepeatedBytes,
		canotoRepeatedFixedBytes,
		canotoRepeatedValue,
		canotoRepeatedPointer,
		canotoRepeatedField,

		canotoFixedRepeatedInt,
		canotoFixedRepeatedUint,
		canotoFixedRepeatedFint32,
		canotoFixedRepeatedFint64,
		canotoFixedRepeatedBool,
		canotoFixedRepeatedString,
		canotoFixedRepeatedBytes,
		canotoFixedRepeatedFixedBytes,
		canotoFixedRepeatedValue,
		canotoFixedRepeatedPointer,
		canotoFixedRepeatedField,
	}
	canotoIntTypes = []canotoType{
		canotoInt,
		canotoRepeatedInt,
		canotoFixedRepeatedInt,
	}
	canotoUintTypes = []canotoType{
		canotoUint,
		canotoRepeatedUint,
		canotoFixedRepeatedUint,
	}
	canotoVarintTypes   = append(canotoIntTypes, canotoUintTypes...)
	canotoRepeatedTypes = append(
		[]canotoType{
			canotoRepeatedInt,
			canotoRepeatedUint,
			canotoRepeatedFint32,
			canotoRepeatedFint64,
			canotoRepeatedBool,
			canotoRepeatedString,
			canotoRepeatedBytes,
			canotoRepeatedFixedBytes,
			canotoRepeatedValue,
			canotoRepeatedPointer,
			canotoRepeatedField,
		},
		canotoFixedRepeatedTypes...,
	)
	canotoFixedRepeatedTypes = []canotoType{
		canotoFixedRepeatedInt,
		canotoFixedRepeatedUint,
		canotoFixedRepeatedFint32,
		canotoFixedRepeatedFint64,
		canotoFixedRepeatedBool,
		canotoFixedRepeatedString,
		canotoFixedRepeatedBytes,
		canotoFixedRepeatedFixedBytes,
		canotoFixedRepeatedValue,
		canotoFixedRepeatedPointer,
		canotoFixedRepeatedField,
	}
	canotoMessageTypes = []canotoType{
		canotoValue,
		canotoPointer,
		canotoField,

		canotoRepeatedValue,
		canotoRepeatedPointer,
		canotoRepeatedField,

		canotoFixedRepeatedValue,
		canotoFixedRepeatedPointer,
		canotoFixedRepeatedField,
	}

	goIntToProto = map[string]string{
		"int8":  "sint32",
		"int16": "sint32",
		"int32": "sint32",
		"int64": "sint64",
		"rune":  "sint32",
	}
	goUintToProto = map[string]string{
		"uint8":  "uint32",
		"uint16": "uint32",
		"uint32": "uint32",
		"uint64": "uint64",
		"byte":   "uint32",
	}
	goFint32ToProto = map[string]string{
		"int32":  "sfixed32",
		"uint32": "fixed32",
	}
	goFint64ToProto = map[string]string{
		"int64":  "sfixed64",
		"uint64": "fixed64",
	}

	errUnexpectedCanotoType = errors.New("unexpected canoto type")
)

type canotoType string

func (c canotoType) IsValid() bool {
	return slices.Contains(canotoTypes, c)
}

func (c canotoType) IsInt() bool {
	return slices.Contains(canotoIntTypes, c)
}

func (c canotoType) IsVarint() bool {
	return slices.Contains(canotoVarintTypes, c)
}

func (c canotoType) IsRepeated() bool {
	return slices.Contains(canotoRepeatedTypes, c) || c.IsFixed()
}

func (c canotoType) IsFixed() bool {
	return slices.Contains(canotoFixedRepeatedTypes, c)
}

func (c canotoType) IsMessage() bool {
	return slices.Contains(canotoMessageTypes, c)
}

func (c canotoType) WireType() canoto.WireType {
	switch c {
	case canotoInt, canotoUint, canotoBool:
		return canoto.Varint
	case canotoFint32:
		return canoto.I32
	case canotoFint64:
		return canoto.I64
	default:
		return canoto.Len
	}
}

func (c canotoType) ProtoType(goType string) string {
	switch c {
	case canotoInt, canotoRepeatedInt, canotoFixedRepeatedInt:
		return goIntToProto[goType]
	case canotoUint, canotoRepeatedUint, canotoFixedRepeatedUint:
		return goUintToProto[goType]
	case canotoFint32, canotoRepeatedFint32, canotoFixedRepeatedFint32:
		return goFint32ToProto[goType]
	case canotoFint64, canotoRepeatedFint64, canotoFixedRepeatedFint64:
		return goFint64ToProto[goType]
	default:
		return ""
	}
}

func (c canotoType) ProtoTypePrefix() string {
	switch c {
	case canotoInt, canotoUint, canotoFint32, canotoFint64, canotoBool, canotoString, canotoBytes, canotoFixedBytes, canotoValue, canotoPointer, canotoField:
		return ""
	default:
		return "repeated "
	}
}

func (c canotoType) ProtoTypeSuffix() string {
	switch c {
	case canotoInt, canotoRepeatedInt, canotoFixedRepeatedInt:
		return "sint64"
	case canotoUint, canotoRepeatedUint, canotoFixedRepeatedUint:
		return "uint64"
	case canotoFint32, canotoRepeatedFint32, canotoFixedRepeatedFint32:
		return "fixed32"
	case canotoFint64, canotoRepeatedFint64, canotoFixedRepeatedFint64:
		return "fixed64"
	case canotoBool, canotoRepeatedBool, canotoFixedRepeatedBool:
		return "bool"
	case canotoString, canotoRepeatedString, canotoFixedRepeatedString:
		return "string"
	default:
		return "bytes"
	}
}

func (c canotoType) Suffix() string {
	switch c {
	case canotoInt, canotoRepeatedInt, canotoFixedRepeatedInt:
		return "Int"
	case canotoUint, canotoRepeatedUint, canotoFixedRepeatedUint:
		return "Uint"
	case canotoFint32, canotoRepeatedFint32, canotoFixedRepeatedFint32:
		return "Fint32"
	case canotoFint64, canotoRepeatedFint64, canotoFixedRepeatedFint64:
		return "Fint64"
	case canotoBool, canotoRepeatedBool, canotoFixedRepeatedBool:
		return "Bool"
	case canotoString, canotoRepeatedString, canotoFixedRepeatedString:
		return "String"
	default:
		return "Bytes"
	}
}
