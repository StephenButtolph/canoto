package canoto

import "errors"

const (
	VarintType WireType = iota
	I64Type
	LenType
	_ // SGROUP is deprecated and not supported
	_ // EGROUP is deprecated and not supported
	I32Type

	MaxFieldNumber = 1<<29 - 1

	wireTypeLength = 3
	wireTypeMask   = 0x07
)

var errInvalidWireType = errors.New("invalid wire type")

type WireType byte

func (w WireType) IsValid() bool {
	switch w {
	case VarintType, I64Type, LenType, I32Type:
		return true
	default:
		return false
	}
}
