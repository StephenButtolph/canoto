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
	if c.CalculateCanotoSize() != len(r.B) {
		return canoto.ErrPaddedZeroes
	}
	return nil
}

func (*CustomType) ValidCanoto() bool {
	return true
}

func (c *CustomType) CalculateCanotoSize() int {
	if c.Int == nil {
		return 0
	}
	return (c.Int.BitLen() + 7) / 8
}

func (c *CustomType) CachedCanotoSize() int {
	return c.CalculateCanotoSize()
}

func (c *CustomType) MarshalCanotoInto(w *canoto.Writer) {
	if c.Int == nil {
		return
	}
	startIndex := len(w.B)
	w.B = append(w.B, make([]byte, c.CalculateCanotoSize())...)
	c.Int.FillBytes(w.B[startIndex:])
}
