package generate

import (
	"cmp"
)

type field struct {
	name              string
	canonicalizedName string
	goType            string
	// isLocal is true if goType is not imported or a type parameter, so it may
	// refer to a message declared in the parsed file.
	isLocal      bool
	protoType    string
	canotoType   canotoType
	fieldNumber  uint32
	oneOfName    string
	templateArgs map[string]string
}

func (f field) Compare(other field) int {
	return cmp.Compare(f.fieldNumber, other.fieldNumber)
}
