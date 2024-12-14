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
	canotoField      canotoType = "field"

	canotoRepeatedInt        = "repeated " + canotoInt
	canotoRepeatedSint       = "repeated " + canotoSint
	canotoRepeatedFint32     = "repeated " + canotoFint32
	canotoRepeatedFint64     = "repeated " + canotoFint64
	canotoRepeatedBool       = "repeated " + canotoBool
	canotoRepeatedString     = "repeated " + canotoString
	canotoRepeatedBytes      = "repeated " + canotoBytes
	canotoRepeatedFixedBytes = "repeated " + canotoFixedBytes
	canotoRepeatedField      = "repeated " + canotoField

	canotoFixedRepeatedInt        = "fixed " + canotoRepeatedInt
	canotoFixedRepeatedSint       = "fixed " + canotoRepeatedSint
	canotoFixedRepeatedFint32     = "fixed " + canotoRepeatedFint32
	canotoFixedRepeatedFint64     = "fixed " + canotoRepeatedFint64
	canotoFixedRepeatedBool       = "fixed " + canotoRepeatedBool
	canotoFixedRepeatedString     = "fixed " + canotoRepeatedString
	canotoFixedRepeatedBytes      = "fixed " + canotoRepeatedBytes
	canotoFixedRepeatedFixedBytes = "fixed " + canotoRepeatedFixedBytes
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
		canotoField,

		canotoRepeatedInt,
		canotoRepeatedSint,
		canotoRepeatedFint32,
		canotoRepeatedFint64,
		canotoRepeatedBool,
		canotoRepeatedString,
		canotoRepeatedBytes,
		canotoRepeatedFixedBytes,
		canotoRepeatedField,

		canotoFixedRepeatedInt,
		canotoFixedRepeatedSint,
		canotoFixedRepeatedFint32,
		canotoFixedRepeatedFint64,
		canotoFixedRepeatedBool,
		canotoFixedRepeatedString,
		canotoFixedRepeatedBytes,
		canotoFixedRepeatedFixedBytes,
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
	canotoRepeatedTypes = []canotoType{
		canotoRepeatedInt,
		canotoRepeatedSint,
		canotoRepeatedFint32,
		canotoRepeatedFint64,
		canotoRepeatedBool,
		canotoRepeatedString,
		canotoRepeatedBytes,
		canotoRepeatedFixedBytes,
		canotoRepeatedField,

		canotoFixedRepeatedInt,
		canotoFixedRepeatedSint,
		canotoFixedRepeatedFint32,
		canotoFixedRepeatedFint64,
		canotoFixedRepeatedBool,
		canotoFixedRepeatedString,
		canotoFixedRepeatedBytes,
		canotoFixedRepeatedFixedBytes,
		canotoFixedRepeatedField,
	}
	canotoFixedRepeatedTypes = []canotoType{
		canotoFixedRepeatedInt,
		canotoFixedRepeatedSint,
		canotoFixedRepeatedFint32,
		canotoFixedRepeatedFint64,
		canotoFixedRepeatedBool,
		canotoFixedRepeatedString,
		canotoFixedRepeatedBytes,
		canotoFixedRepeatedFixedBytes,
		canotoFixedRepeatedField,
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
