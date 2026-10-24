package big

import (
	"math/big"
	"math/bits"
	"reflect"

	"github.com/StephenButtolph/canoto"
)

var _ canoto.Field = (*Int)(nil)

// Int is a non-negative integer, encoded as big-endian bytes without leading
// zeros.
type Int struct {
	Int *big.Int
}

func (*Int) DescribeCanoto(...reflect.Type) *canoto.Spec {
	// Nil indicates that the type does not have a valid spec. This type will be
	// treated as opaque bytes.
	return nil
}

func (i *Int) UnmarshalCanotoFrom(r canoto.Reader) (canoto.UnmarshalResult, error) {
	*i = Int{}
	if len(r.B) == 0 {
		return canoto.UnmarshalResult{MemoryRemaining: r.MemoryRemaining}, nil
	}
	if r.B[0] == 0 {
		return canoto.UnmarshalResult{}, canoto.ErrPaddedZeroes
	}

	const wordSize = bits.UintSize / 8
	var words []big.Word
	numWords := (uint64(len(r.B)) + wordSize - 1) / wordSize
	if err := canoto.MakeSlice(&r, &words, numWords); err != nil {
		return canoto.UnmarshalResult{}, err
	}
	if err := canoto.MakePointer(&r, &i.Int); err != nil {
		return canoto.UnmarshalResult{}, err
	}
	i.Int.SetBits(words[:0]).SetBytes(r.B)
	return canoto.UnmarshalResult{MemoryRemaining: r.MemoryRemaining}, nil
}

func (i *Int) CheckCanoto() bool {
	return i.Int == nil || i.Int.Sign() >= 0
}

func (*Int) CacheCanoto() {}

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
