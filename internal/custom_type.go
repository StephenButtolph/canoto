package examples

import (
	"math/big"

	"github.com/StephenButtolph/canoto"
)

var _ canoto.Field = (*CustomType)(nil)

type CustomType struct {
	Int *big.Int
}

func (c *CustomType) UnmarshalCanotoFrom(r *canoto.Reader) error {
	if c.Int == nil {
		c.Int = new(big.Int)
	}
	c.Int.SetBytes(r.B)
	if c.CachedCanotoSize() != len(r.B) {
		return canoto.ErrPaddedZeroes
	}
	return nil
}

func (*CustomType) ValidCanoto() bool     { return true }
func (*CustomType) CalculateCanotoCache() {}

func (c *CustomType) CachedCanotoSize() int {
	if c.Int == nil {
		return 0
	}
	return (c.Int.BitLen() + 7) / 8
}

func (c *CustomType) MarshalCanotoInto(w *canoto.Writer) {
	if c.Int == nil {
		return
	}
	startIndex := len(w.B)
	w.B = append(w.B, make([]byte, c.CachedCanotoSize())...)
	c.Int.FillBytes(w.B[startIndex:])
}
