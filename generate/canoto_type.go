package generate

import (
	"errors"
	"slices"

	"github.com/StephenButtolph/canoto"
)

const (
	canotoInt        canotoType = "int"
	canotoSint       canotoType = "sint"   // signed int
	canotoFint32     canotoType = "fint32" // fixed 32-bit int
	canotoFint64     canotoType = "fint64" // fixed 64-bit int
	canotoBool       canotoType = "bool"
	canotoString     canotoType = "string"
	canotoBytes      canotoType = "bytes"
	canotoFixedBytes canotoType = "fixed bytes"
	canotoValue      canotoType = "value"
	canotoPointer    canotoType = "pointer"
	canotoField      canotoType = "field"

	canotoRepeatedInt        = "repeated " + canotoInt
	canotoRepeatedSint       = "repeated " + canotoSint
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
	canotoFixedRepeatedSint       = "fixed " + canotoRepeatedSint
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
		canotoSint,
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
		canotoRepeatedSint,
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
		canotoFixedRepeatedSint,
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
	canotoVarintTypes = []canotoType{
		canotoInt,
		canotoSint,

		canotoRepeatedInt,
		canotoRepeatedSint,

		canotoFixedRepeatedInt,
		canotoFixedRepeatedSint,
	}
	canotoRepeatedTypes = append(
		[]canotoType{
			canotoRepeatedInt,
			canotoRepeatedSint,
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
		canotoFixedRepeatedSint,
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

	goIntToProtoInt = map[string]string{
		"int8":   "int32",
		"int16":  "int32",
		"int32":  "int32",
		"int64":  "int64",
		"uint8":  "uint32",
		"uint16": "uint32",
		"uint32": "uint32",
		"uint64": "uint64",
		"byte":   "uint32",
		"rune":   "int32",
	}
	goSintToProtoSint = map[string]string{
		"int8":  "sint32",
		"int16": "sint32",
		"int32": "sint32",
		"int64": "sint64",
		"rune":  "sint32",
	}
	goFint32ToProtoSint = map[string]string{
		"int32":  "sfixed32",
		"uint32": "fixed32",
	}
	goFint64ToProtoSint = map[string]string{
		"int64":  "sfixed64",
		"uint64": "fixed64",
	}

	errUnexpectedCanotoType = errors.New("unexpected canoto type")
)

type canotoType string

func (c canotoType) IsValid() bool {
	return slices.Contains(canotoTypes, c)
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

func (c canotoType) WireType() canoto.WireType {
	switch c {
	case canotoInt, canotoSint, canotoBool:
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
		return goIntToProtoInt[goType]
	case canotoSint, canotoRepeatedSint, canotoFixedRepeatedSint:
		return goSintToProtoSint[goType]
	case canotoFint32, canotoRepeatedFint32, canotoFixedRepeatedFint32:
		return goFint32ToProtoSint[goType]
	case canotoFint64, canotoRepeatedFint64, canotoFixedRepeatedFint64:
		return goFint64ToProtoSint[goType]
	default:
		return ""
	}
}

func (c canotoType) ProtoTypePrefix() string {
	switch c {
	case canotoInt, canotoSint, canotoFint32, canotoFint64, canotoBool, canotoString, canotoBytes, canotoFixedBytes, canotoValue, canotoPointer, canotoField:
		return ""
	default:
		return "repeated "
	}
}

func (c canotoType) ProtoTypeSuffix() string {
	switch c {
	case canotoInt, canotoRepeatedInt, canotoFixedRepeatedInt:
		return "uint64"
	case canotoSint, canotoRepeatedSint, canotoFixedRepeatedSint:
		return "sint64"
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
	case canotoSint, canotoRepeatedSint, canotoFixedRepeatedSint:
		return "Sint"
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
