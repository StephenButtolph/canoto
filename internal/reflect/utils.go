//go:generate canoto $GOFILE

package reflect

import (
	"github.com/StephenButtolph/canoto"
)

func FieldTypeFromInt[T canoto.Int](
	_ T,
	fieldNumber uint32,
	name string,
) *FieldType {
	f := &FieldType{
		FieldNumber: fieldNumber,
		Name:        name,
	}
	if isSigned[T]() {
		f.TypeInt = bitLength[T]()
	} else {
		f.TypeUint = bitLength[T]()
	}
	return f
}

func FieldTypeFromSint[T canoto.Sint](
	_ T,
	fieldNumber uint32,
	name string,
) *FieldType {
	return &FieldType{
		FieldNumber: fieldNumber,
		Name:        name,
		TypeSint:    bitLength[T](),
	}
}

func isSigned[T canoto.Int]() bool {
	return ^T(0) < T(0)
}

func bitLength[T canoto.Int]() uint8 {
	for _, l := range []uint8{8, 16, 32, 64} {
		if T(1)<<l == T(0) {
			return l
		}
	}
	panic("unsupported integer size")
}
