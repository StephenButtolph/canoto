package big

import (
	"math/big"

	"github.com/StephenButtolph/canoto"
)

var _ canoto.Field = (*Int)(nil)

type Int struct {
	Int *big.Int
}

func (i *Int) UnmarshalCanotoFrom(r *canoto.Reader) error {
	if i.Int == nil {
		i.Int = new(big.Int)
	}
	i.Int.SetBytes(r.B)
	if i.CachedCanotoSize() != len(r.B) {
		return canoto.ErrPaddedZeroes
	}
	return nil
}

func (*Int) ValidCanoto() bool     { return true }
func (*Int) CalculateCanotoCache() {}

func (i *Int) CachedCanotoSize() int {
	if i.Int == nil {
		return 0
	}
	return (i.Int.BitLen() + 7) / 8
}

func (i *Int) MarshalCanotoInto(w *canoto.Writer) {
	if i.Int == nil {
		return
	}
	startIndex := len(w.B)
	w.B = append(w.B, make([]byte, i.CachedCanotoSize())...)
	i.Int.FillBytes(w.B[startIndex:])
}
