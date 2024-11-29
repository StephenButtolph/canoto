package canoto

import (
	"math/bits"
)

const (
	SizeFixed32  = 4
	SizeFixed64  = 8
	SizeSfixed32 = SizeFixed32
	SizeSfixed64 = SizeFixed64
	SizeBool     = 1
)

func SizeTag(fieldNumber uint32, wireType WireType) int {
	return SizeUint64(tagToUint64(fieldNumber, wireType))
}

func SizeInt32(v int32) int {
	return SizeUint64(uint64(v))
}

func SizeInt64(v int64) int {
	return SizeUint64(uint64(v))
}

func SizeUint32(v uint32) int {
	return SizeUint64(uint64(v))
}

func SizeUint64(v uint64) int {
	if v == 0 {
		return 1
	}
	return (bits.Len64(v) + 6) / 7
}

func SizeSint32(v int32) int {
	return SizeUint64(sint32ToUint64(v))
}

func SizeSint64(v int64) int {
	return SizeUint64(sint64ToUint64(v))
}

func SizeString(v string) int {
	return SizeUint64(uint64(len(v))) + len(v)
}

func SizeBytes(v []byte) int {
	return SizeUint64(uint64(len(v))) + len(v)
}
