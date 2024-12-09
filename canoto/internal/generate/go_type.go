package generate

import (
	"errors"
	"slices"
)

const (
	goInt32  goType = "int32"
	goInt64  goType = "int64"
	goUint32 goType = "uint32"
	goUint64 goType = "uint64"
	goBool   goType = "bool"
	goString goType = "string"
	goBytes  goType = "[]byte"
)

var (
	primitiveGoTypes = []goType{
		goInt32,
		goInt64,
		goUint32,
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
