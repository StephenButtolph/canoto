package big

import (
	"math/big"
	"reflect"

	"github.com/StephenButtolph/canoto"
)

var _ canoto.Field = (*Int)(nil)

type Int struct {
	Int *big.Int
}

func (*Int) DescribeCanoto(...reflect.Type) *canoto.Spec {
	// Nil indicates that the type does not have a valid spec. This type will be
	// treated as opaque bytes.
	return nil
}

func (i *Int) UnmarshalCanotoFrom(r canoto.Reader) error {
	if i.Int == nil {
		i.Int = new(big.Int)
	}
	i.Int.SetBytes(r.B)
	if i.SizeCanoto() != uint64(len(r.B)) {
		return canoto.ErrPaddedZeroes
	}
	return nil
}

func (*Int) CheckCanoto() bool { return true }
func (*Int) CacheCanoto()      {}

func (i *Int) SizeCanoto() uint64 {
	if i.Int == nil {
		return 0
	}
	return uint64(i.Int.BitLen()+7) / 8 //#nosec G115 // False positive
}

func (i *Int) AppendCanoto(w canoto.Writer) canoto.Writer {
	if i.Int == nil {
		return w
	}
	startIndex := len(w.B)
	w.B = append(w.B, make([]byte, i.SizeCanoto())...)
	i.Int.FillBytes(w.B[startIndex:])
	return w
}
