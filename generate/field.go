package generate

import (
	"cmp"
)

type field struct {
	name              string
	canonicalizedName string
	goType            string
	protoType         string
	canotoType        canotoType
	fieldNumber       uint32
	oneOfName         string
	templateArgs      map[string]string
}

func (f field) Compare(other field) int {
	return cmp.Compare(f.fieldNumber, other.fieldNumber)
}
