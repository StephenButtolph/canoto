package generate

import (
	"cmp"
	"fmt"

	"github.com/StephenButtolph/canoto"
)

type field struct {
	name              string
	canonicalizedName string
	fixedLength       [2]bool
	repeated          bool
	goType            goType
	canotoType        canotoType
	fieldNumber       uint32
	templateArgs      map[string]string
}

func (f field) Compare(other field) int {
	return cmp.Compare(f.fieldNumber, other.fieldNumber)
}

func (f field) WireType() (canoto.WireType, error) {
	if f.repeated {
		return canoto.Len, nil
	}
	switch f.canotoType {
	case canotoInt, canotoSint, canotoBool:
		return canoto.Varint, nil
	case canotoFint:
		switch f.goType {
		case goInt32, goUint32:
			return canoto.I32, nil
		case goInt64, goUint64:
			return canoto.I64, nil
		default:
			return 0, fmt.Errorf("%w: %q with canotoType %q", errUnexpectedGoType, f.goType, f.canotoType)
		}
	case canotoBytes:
		return canoto.Len, nil
	default:
		return 0, fmt.Errorf("%w: %q", errUnexpectedCanotoType, f.canotoType)
	}
}
