package generate

import (
	"errors"
	"slices"
)

const (
	goInt8   goType = "int8"
	goUint8  goType = "uint8"
	goInt16  goType = "int16"
	goUint16 goType = "uint16"
	goInt32  goType = "int32"
	goUint32 goType = "uint32"
	goInt64  goType = "int64"
	goUint64 goType = "uint64"
	goBool   goType = "bool"
	goString goType = "string"
	goBytes  goType = "[]byte"
)

var (
	primitiveGoTypes = []goType{
		goInt8,
		goUint8,
		goInt16,
		goUint16,
		goInt32,
		goUint32,
		goInt64,
		goUint64,
		goBool,
		goString,
		goBytes,
	}

	errUnexpectedGoType = errors.New("unexpected go type")
)

type goType string

func (g goType) IsPrimitive() bool {
	return slices.Contains(primitiveGoTypes, g)
}
