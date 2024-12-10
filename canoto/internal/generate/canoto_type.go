package generate

import (
	"errors"
	"slices"
)

const (
	canotoInt   canotoType = "int"
	canotoSint  canotoType = "sint" // signed int
	canotoFint  canotoType = "fint" // fixed int
	canotoBool  canotoType = "bool"
	canotoBytes canotoType = "bytes"
)

var (
	canotoTypes = []canotoType{
		canotoInt,
		canotoSint,
		canotoFint,
		canotoBool,
		canotoBytes,
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
