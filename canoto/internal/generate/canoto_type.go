package generate

import (
	"errors"
	"slices"
)

const (
	canotoInt    canotoType = "int"
	canotoSint   canotoType = "sint"   // signed int
	canotoFint32 canotoType = "fint32" // fixed 32-bit int
	canotoFint64 canotoType = "fint64" // fixed 64-bit int
	canotoBool   canotoType = "bool"
	canotoString canotoType = "string"
	canotoBytes  canotoType = "bytes"
	canotoField  canotoType = "field"
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
		canotoField,
	}
	canotoVarintTypes = []canotoType{
		canotoInt,
		canotoSint,
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
