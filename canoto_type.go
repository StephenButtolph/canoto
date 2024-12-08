package canoto

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

	errUnexpectedCanotoType = errors.New("unexpected canoto type")
)

type canotoType string

func (c canotoType) IsValid() bool {
	return slices.Contains(canotoTypes, c)
}
