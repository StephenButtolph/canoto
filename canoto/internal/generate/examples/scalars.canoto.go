// Code generated by Canoto. DO NOT EDIT.

package examples

import (
	"unicode/utf8"

	"github.com/StephenButtolph/canoto"
)

// Ensure that "unicode/utf8" is imported without error
var _ = utf8.ValidString

const (
	canoto__Scalars__Int32__tag = "\x08" // canoto.Tag(1, canoto.Varint)
	canoto__Scalars__Int64__tag = "\x10" // canoto.Tag(2, canoto.Varint)
	canoto__Scalars__Uint32__tag = "\x18" // canoto.Tag(3, canoto.Varint)
	canoto__Scalars__Uint64__tag = "\x20" // canoto.Tag(4, canoto.Varint)
	canoto__Scalars__Sint32__tag = "\x28" // canoto.Tag(5, canoto.Varint)
	canoto__Scalars__Sint64__tag = "\x30" // canoto.Tag(6, canoto.Varint)
	canoto__Scalars__Fixed32__tag = "\x3d" // canoto.Tag(7, canoto.I32)
	canoto__Scalars__Fixed64__tag = "\x41" // canoto.Tag(8, canoto.I64)
	canoto__Scalars__Sfixed32__tag = "\x4d" // canoto.Tag(9, canoto.I32)
	canoto__Scalars__Sfixed64__tag = "\x51" // canoto.Tag(10, canoto.I64)
	canoto__Scalars__Bool__tag = "\x58" // canoto.Tag(11, canoto.Varint)
	canoto__Scalars__String__tag = "\x62" // canoto.Tag(12, canoto.Len)
	canoto__Scalars__Bytes__tag = "\x6a" // canoto.Tag(13, canoto.Len)
	canoto__Scalars__LargestFieldNumber__tag = "\x72" // canoto.Tag(14, canoto.Len)
	canoto__Scalars__RepeatedInt32__tag = "\x7a" // canoto.Tag(15, canoto.Len)
	canoto__Scalars__RepeatedInt64__tag = "\x82\x01" // canoto.Tag(16, canoto.Len)
	canoto__Scalars__RepeatedUint32__tag = "\x8a\x01" // canoto.Tag(17, canoto.Len)
	canoto__Scalars__RepeatedUint64__tag = "\x92\x01" // canoto.Tag(18, canoto.Len)
	canoto__Scalars__RepeatedSint32__tag = "\x9a\x01" // canoto.Tag(19, canoto.Len)
	canoto__Scalars__RepeatedSint64__tag = "\xa2\x01" // canoto.Tag(20, canoto.Len)
	canoto__Scalars__RepeatedFixed32__tag = "\xaa\x01" // canoto.Tag(21, canoto.Len)
	canoto__Scalars__RepeatedFixed64__tag = "\xb2\x01" // canoto.Tag(22, canoto.Len)
	canoto__Scalars__RepeatedSfixed32__tag = "\xba\x01" // canoto.Tag(23, canoto.Len)
	canoto__Scalars__RepeatedSfixed64__tag = "\xc2\x01" // canoto.Tag(24, canoto.Len)
	canoto__Scalars__RepeatedBool__tag = "\xca\x01" // canoto.Tag(25, canoto.Len)

	canoto__Scalars__Int32__tag__size = len(canoto__Scalars__Int32__tag)
	canoto__Scalars__Int64__tag__size = len(canoto__Scalars__Int64__tag)
	canoto__Scalars__Uint32__tag__size = len(canoto__Scalars__Uint32__tag)
	canoto__Scalars__Uint64__tag__size = len(canoto__Scalars__Uint64__tag)
	canoto__Scalars__Sint32__tag__size = len(canoto__Scalars__Sint32__tag)
	canoto__Scalars__Sint64__tag__size = len(canoto__Scalars__Sint64__tag)
	canoto__Scalars__Fixed32__tag__size = len(canoto__Scalars__Fixed32__tag)
	canoto__Scalars__Fixed64__tag__size = len(canoto__Scalars__Fixed64__tag)
	canoto__Scalars__Sfixed32__tag__size = len(canoto__Scalars__Sfixed32__tag)
	canoto__Scalars__Sfixed64__tag__size = len(canoto__Scalars__Sfixed64__tag)
	canoto__Scalars__Bool__tag__size = len(canoto__Scalars__Bool__tag)
	canoto__Scalars__String__tag__size = len(canoto__Scalars__String__tag)
	canoto__Scalars__Bytes__tag__size = len(canoto__Scalars__Bytes__tag)
	canoto__Scalars__LargestFieldNumber__tag__size = len(canoto__Scalars__LargestFieldNumber__tag)
	canoto__Scalars__RepeatedInt32__tag__size = len(canoto__Scalars__RepeatedInt32__tag)
	canoto__Scalars__RepeatedInt64__tag__size = len(canoto__Scalars__RepeatedInt64__tag)
	canoto__Scalars__RepeatedUint32__tag__size = len(canoto__Scalars__RepeatedUint32__tag)
	canoto__Scalars__RepeatedUint64__tag__size = len(canoto__Scalars__RepeatedUint64__tag)
	canoto__Scalars__RepeatedSint32__tag__size = len(canoto__Scalars__RepeatedSint32__tag)
	canoto__Scalars__RepeatedSint64__tag__size = len(canoto__Scalars__RepeatedSint64__tag)
	canoto__Scalars__RepeatedFixed32__tag__size = len(canoto__Scalars__RepeatedFixed32__tag)
	canoto__Scalars__RepeatedFixed64__tag__size = len(canoto__Scalars__RepeatedFixed64__tag)
	canoto__Scalars__RepeatedSfixed32__tag__size = len(canoto__Scalars__RepeatedSfixed32__tag)
	canoto__Scalars__RepeatedSfixed64__tag__size = len(canoto__Scalars__RepeatedSfixed64__tag)
	canoto__Scalars__RepeatedBool__tag__size = len(canoto__Scalars__RepeatedBool__tag)
)

type canotoData_Scalars struct {
	size int
	RepeatedInt32Size int
	RepeatedInt64Size int
	RepeatedUint32Size int
	RepeatedUint64Size int
	RepeatedSint32Size int
	RepeatedSint64Size int
}


func (c *Scalars) UnmarshalCanoto(bytes []byte) error {
	r := canoto.Reader{
		B: bytes,
	}
	return c.UnmarshalCanotoFrom(&r)
}

func (c *Scalars) UnmarshalCanotoFrom(r *canoto.Reader) error {
	var minField uint32
	for canoto.HasNext(r) {
		field, wireType, err := canoto.ReadTag(r)
		if err != nil {
			return err
		}
		if field < minField {
			return canoto.ErrInvalidFieldOrder
		}

		switch field {
		case 1:
			if wireType != canoto.Varint {
				return canoto.ErrInvalidWireType
			}
			c.Int32, err = canoto.ReadInt[int32](r)
			if err != nil {
				return err
			}
			if c.Int32 == 0 {
				return canoto.ErrZeroValue
			}
		case 2:
			if wireType != canoto.Varint {
				return canoto.ErrInvalidWireType
			}
			c.Int64, err = canoto.ReadInt[int64](r)
			if err != nil {
				return err
			}
			if c.Int64 == 0 {
				return canoto.ErrZeroValue
			}
		case 3:
			if wireType != canoto.Varint {
				return canoto.ErrInvalidWireType
			}
			c.Uint32, err = canoto.ReadInt[uint32](r)
			if err != nil {
				return err
			}
			if c.Uint32 == 0 {
				return canoto.ErrZeroValue
			}
		case 4:
			if wireType != canoto.Varint {
				return canoto.ErrInvalidWireType
			}
			c.Uint64, err = canoto.ReadInt[uint64](r)
			if err != nil {
				return err
			}
			if c.Uint64 == 0 {
				return canoto.ErrZeroValue
			}
		case 5:
			if wireType != canoto.Varint {
				return canoto.ErrInvalidWireType
			}
			c.Sint32, err = canoto.ReadSint[int32](r)
			if err != nil {
				return err
			}
			if c.Sint32 == 0 {
				return canoto.ErrZeroValue
			}
		case 6:
			if wireType != canoto.Varint {
				return canoto.ErrInvalidWireType
			}
			c.Sint64, err = canoto.ReadSint[int64](r)
			if err != nil {
				return err
			}
			if c.Sint64 == 0 {
				return canoto.ErrZeroValue
			}
		case 7:
			if wireType != canoto.I32 {
				return canoto.ErrInvalidWireType
			}
			c.Fixed32, err = canoto.ReadFint32[uint32](r)
			if err != nil {
				return err
			}
			if c.Fixed32 == 0 {
				return canoto.ErrZeroValue
			}
		case 8:
			if wireType != canoto.I64 {
				return canoto.ErrInvalidWireType
			}
			c.Fixed64, err = canoto.ReadFint64[uint64](r)
			if err != nil {
				return err
			}
			if c.Fixed64 == 0 {
				return canoto.ErrZeroValue
			}
		case 9:
			if wireType != canoto.I32 {
				return canoto.ErrInvalidWireType
			}
			c.Sfixed32, err = canoto.ReadFint32[int32](r)
			if err != nil {
				return err
			}
			if c.Sfixed32 == 0 {
				return canoto.ErrZeroValue
			}
		case 10:
			if wireType != canoto.I64 {
				return canoto.ErrInvalidWireType
			}
			c.Sfixed64, err = canoto.ReadFint64[int64](r)
			if err != nil {
				return err
			}
			if c.Sfixed64 == 0 {
				return canoto.ErrZeroValue
			}
		case 11:
			if wireType != canoto.Varint {
				return canoto.ErrInvalidWireType
			}
			c.Bool, err = canoto.ReadBool(r)
			if err != nil {
				return err
			}
			if !c.Bool {
				return canoto.ErrZeroValue
			}
		case 12:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}
			c.String, err = canoto.ReadString(r)
			if err != nil {
				return err
			}
			if len(c.String) == 0 {
				return canoto.ErrZeroValue
			}
		case 13:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}
			c.Bytes, err = canoto.ReadBytes(r)
			if err != nil {
				return err
			}
			if len(c.Bytes) == 0 {
				return canoto.ErrZeroValue
			}
		case 14:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			err = c.LargestFieldNumber.UnmarshalCanotoFrom(r)
			r.B = remainingBytes
			if err != nil {
				return err
			}
		case 15:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.RepeatedInt32 = make([]int32, 0, canoto.CountInts(msgBytes))
			for canoto.HasNext(r) {
				v, err := canoto.ReadInt[int32](r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.RepeatedInt32 = append(c.RepeatedInt32, v)
			}
			r.B = remainingBytes
		case 16:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.RepeatedInt64 = make([]int64, 0, canoto.CountInts(msgBytes))
			for canoto.HasNext(r) {
				v, err := canoto.ReadInt[int64](r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.RepeatedInt64 = append(c.RepeatedInt64, v)
			}
			r.B = remainingBytes
		case 17:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.RepeatedUint32 = make([]uint32, 0, canoto.CountInts(msgBytes))
			for canoto.HasNext(r) {
				v, err := canoto.ReadInt[uint32](r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.RepeatedUint32 = append(c.RepeatedUint32, v)
			}
			r.B = remainingBytes
		case 18:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.RepeatedUint64 = make([]uint64, 0, canoto.CountInts(msgBytes))
			for canoto.HasNext(r) {
				v, err := canoto.ReadInt[uint64](r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.RepeatedUint64 = append(c.RepeatedUint64, v)
			}
			r.B = remainingBytes
		case 19:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.RepeatedSint32 = make([]int32, 0, canoto.CountInts(msgBytes))
			for canoto.HasNext(r) {
				v, err := canoto.ReadSint[int32](r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.RepeatedSint32 = append(c.RepeatedSint32, v)
			}
			r.B = remainingBytes
		case 20:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			if len(msgBytes) == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.RepeatedSint64 = make([]int64, 0, canoto.CountInts(msgBytes))
			for canoto.HasNext(r) {
				v, err := canoto.ReadSint[int64](r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.RepeatedSint64 = append(c.RepeatedSint64, v)
			}
			r.B = remainingBytes
		case 21:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			numMsgBytes := len(msgBytes)
			if numMsgBytes == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.RepeatedFixed32 = make([]uint32, 0, numMsgBytes / canoto.SizeFint32)
			for canoto.HasNext(r) {
				v, err := canoto.ReadFint32[uint32](r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.RepeatedFixed32 = append(c.RepeatedFixed32, v)
			}
			r.B = remainingBytes
		case 22:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			numMsgBytes := len(msgBytes)
			if numMsgBytes == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.RepeatedFixed64 = make([]uint64, 0, numMsgBytes / canoto.SizeFint64)
			for canoto.HasNext(r) {
				v, err := canoto.ReadFint64[uint64](r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.RepeatedFixed64 = append(c.RepeatedFixed64, v)
			}
			r.B = remainingBytes
		case 23:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			numMsgBytes := len(msgBytes)
			if numMsgBytes == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.RepeatedSfixed32 = make([]int32, 0, numMsgBytes / canoto.SizeFint32)
			for canoto.HasNext(r) {
				v, err := canoto.ReadFint32[int32](r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.RepeatedSfixed32 = append(c.RepeatedSfixed32, v)
			}
			r.B = remainingBytes
		case 24:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			numMsgBytes := len(msgBytes)
			if numMsgBytes == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.RepeatedSfixed64 = make([]int64, 0, numMsgBytes / canoto.SizeFint64)
			for canoto.HasNext(r) {
				v, err := canoto.ReadFint64[int64](r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.RepeatedSfixed64 = append(c.RepeatedSfixed64, v)
			}
			r.B = remainingBytes
		case 25:
			if wireType != canoto.Len {
				return canoto.ErrInvalidWireType
			}

			originalUnsafe := r.Unsafe
			r.Unsafe = true
			var msgBytes []byte
			msgBytes, err = canoto.ReadBytes(r)
			r.Unsafe = originalUnsafe
			if err != nil {
				return err
			}
			numMsgBytes := len(msgBytes)
			if numMsgBytes == 0 {
				return canoto.ErrZeroValue
			}

			remainingBytes := r.B
			r.B = msgBytes
			c.RepeatedBool = make([]bool, 0, numMsgBytes / canoto.SizeBool)
			for canoto.HasNext(r) {
				v, err := canoto.ReadBool(r)
				if err != nil {
					r.B = remainingBytes
					return err
				}
				c.RepeatedBool = append(c.RepeatedBool, v)
			}
			r.B = remainingBytes
		default:
			return canoto.ErrUnknownField
		}

		minField = field + 1
	}
	return nil
}

func (c *Scalars) ValidCanoto() bool {
	return utf8.ValidString(c.String) && c.LargestFieldNumber.ValidCanoto()
}

func (c *Scalars) CalculateCanotoSize() int {
	c.canotoData.size = 0
	if c.Int32 != 0 {
		c.canotoData.size += canoto__Scalars__Int32__tag__size + canoto.SizeInt(c.Int32)
	}
	if c.Int64 != 0 {
		c.canotoData.size += canoto__Scalars__Int64__tag__size + canoto.SizeInt(c.Int64)
	}
	if c.Uint32 != 0 {
		c.canotoData.size += canoto__Scalars__Uint32__tag__size + canoto.SizeInt(c.Uint32)
	}
	if c.Uint64 != 0 {
		c.canotoData.size += canoto__Scalars__Uint64__tag__size + canoto.SizeInt(c.Uint64)
	}
	if c.Sint32 != 0 {
		c.canotoData.size += canoto__Scalars__Sint32__tag__size + canoto.SizeSint(c.Sint32)
	}
	if c.Sint64 != 0 {
		c.canotoData.size += canoto__Scalars__Sint64__tag__size + canoto.SizeSint(c.Sint64)
	}
	if c.Fixed32 != 0 {
		c.canotoData.size += canoto__Scalars__Fixed32__tag__size + canoto.SizeFint32
	}
	if c.Fixed64 != 0 {
		c.canotoData.size += canoto__Scalars__Fixed64__tag__size + canoto.SizeFint64
	}
	if c.Sfixed32 != 0 {
		c.canotoData.size += canoto__Scalars__Sfixed32__tag__size + canoto.SizeFint32
	}
	if c.Sfixed64 != 0 {
		c.canotoData.size += canoto__Scalars__Sfixed64__tag__size + canoto.SizeFint64
	}
	if c.Bool {
		c.canotoData.size += canoto__Scalars__Bool__tag__size + canoto.SizeBool
	}
	if len(c.String) != 0 {
		c.canotoData.size += canoto__Scalars__String__tag__size + canoto.SizeBytes(c.String)
	}
	if len(c.Bytes) != 0 {
		c.canotoData.size += canoto__Scalars__Bytes__tag__size + canoto.SizeBytes(c.Bytes)
	}
	if fieldSize := c.LargestFieldNumber.CalculateCanotoSize(); fieldSize != 0 {
		c.canotoData.size += canoto__Scalars__LargestFieldNumber__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
	}
	if len(c.RepeatedInt32) != 0 {
		c.canotoData.RepeatedInt32Size = 0
		for _, v := range c.RepeatedInt32 {
			c.canotoData.RepeatedInt32Size += canoto.SizeInt(v)
		}
		c.canotoData.size += canoto__Scalars__RepeatedInt32__tag__size + canoto.SizeInt(int64(c.canotoData.RepeatedInt32Size)) + c.canotoData.RepeatedInt32Size
	}
	if len(c.RepeatedInt64) != 0 {
		c.canotoData.RepeatedInt64Size = 0
		for _, v := range c.RepeatedInt64 {
			c.canotoData.RepeatedInt64Size += canoto.SizeInt(v)
		}
		c.canotoData.size += canoto__Scalars__RepeatedInt64__tag__size + canoto.SizeInt(int64(c.canotoData.RepeatedInt64Size)) + c.canotoData.RepeatedInt64Size
	}
	if len(c.RepeatedUint32) != 0 {
		c.canotoData.RepeatedUint32Size = 0
		for _, v := range c.RepeatedUint32 {
			c.canotoData.RepeatedUint32Size += canoto.SizeInt(v)
		}
		c.canotoData.size += canoto__Scalars__RepeatedUint32__tag__size + canoto.SizeInt(int64(c.canotoData.RepeatedUint32Size)) + c.canotoData.RepeatedUint32Size
	}
	if len(c.RepeatedUint64) != 0 {
		c.canotoData.RepeatedUint64Size = 0
		for _, v := range c.RepeatedUint64 {
			c.canotoData.RepeatedUint64Size += canoto.SizeInt(v)
		}
		c.canotoData.size += canoto__Scalars__RepeatedUint64__tag__size + canoto.SizeInt(int64(c.canotoData.RepeatedUint64Size)) + c.canotoData.RepeatedUint64Size
	}
	if len(c.RepeatedSint32) != 0 {
		c.canotoData.RepeatedSint32Size = 0
		for _, v := range c.RepeatedSint32 {
			c.canotoData.RepeatedSint32Size += canoto.SizeSint(v)
		}
		c.canotoData.size += canoto__Scalars__RepeatedSint32__tag__size + canoto.SizeInt(int64(c.canotoData.RepeatedSint32Size)) + c.canotoData.RepeatedSint32Size
	}
	if len(c.RepeatedSint64) != 0 {
		c.canotoData.RepeatedSint64Size = 0
		for _, v := range c.RepeatedSint64 {
			c.canotoData.RepeatedSint64Size += canoto.SizeSint(v)
		}
		c.canotoData.size += canoto__Scalars__RepeatedSint64__tag__size + canoto.SizeInt(int64(c.canotoData.RepeatedSint64Size)) + c.canotoData.RepeatedSint64Size
	}
	if num := len(c.RepeatedFixed32); num != 0 {
		fieldSize := num * canoto.SizeFint32
		c.canotoData.size += canoto__Scalars__RepeatedFixed32__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
	}
	if num := len(c.RepeatedFixed64); num != 0 {
		fieldSize := num * canoto.SizeFint64
		c.canotoData.size += canoto__Scalars__RepeatedFixed64__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
	}
	if num := len(c.RepeatedSfixed32); num != 0 {
		fieldSize := num * canoto.SizeFint32
		c.canotoData.size += canoto__Scalars__RepeatedSfixed32__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
	}
	if num := len(c.RepeatedSfixed64); num != 0 {
		fieldSize := num * canoto.SizeFint64
		c.canotoData.size += canoto__Scalars__RepeatedSfixed64__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
	}
	if num := len(c.RepeatedBool); num != 0 {
		fieldSize := num * canoto.SizeBool
		c.canotoData.size += canoto__Scalars__RepeatedBool__tag__size + canoto.SizeInt(int64(fieldSize)) + fieldSize
	}
	return c.canotoData.size
}

func (c *Scalars) CachedCanotoSize() int {
	return c.canotoData.size
}

func (c *Scalars) MarshalCanoto() []byte {
	w := canoto.Writer{
		B: make([]byte, 0, c.CalculateCanotoSize()),
	}
	c.MarshalCanotoInto(&w)
	return w.B
}

func (c *Scalars) MarshalCanotoInto(w *canoto.Writer) {
	if c.Int32 != 0 {
		canoto.Append(w, canoto__Scalars__Int32__tag)
		canoto.AppendInt(w, c.Int32)
	}
	if c.Int64 != 0 {
		canoto.Append(w, canoto__Scalars__Int64__tag)
		canoto.AppendInt(w, c.Int64)
	}
	if c.Uint32 != 0 {
		canoto.Append(w, canoto__Scalars__Uint32__tag)
		canoto.AppendInt(w, c.Uint32)
	}
	if c.Uint64 != 0 {
		canoto.Append(w, canoto__Scalars__Uint64__tag)
		canoto.AppendInt(w, c.Uint64)
	}
	if c.Sint32 != 0 {
		canoto.Append(w, canoto__Scalars__Sint32__tag)
		canoto.AppendSint(w, c.Sint32)
	}
	if c.Sint64 != 0 {
		canoto.Append(w, canoto__Scalars__Sint64__tag)
		canoto.AppendSint(w, c.Sint64)
	}
	if c.Fixed32 != 0 {
		canoto.Append(w, canoto__Scalars__Fixed32__tag)
		canoto.AppendFint32(w, c.Fixed32)
	}
	if c.Fixed64 != 0 {
		canoto.Append(w, canoto__Scalars__Fixed64__tag)
		canoto.AppendFint64(w, c.Fixed64)
	}
	if c.Sfixed32 != 0 {
		canoto.Append(w, canoto__Scalars__Sfixed32__tag)
		canoto.AppendFint32(w, c.Sfixed32)
	}
	if c.Sfixed64 != 0 {
		canoto.Append(w, canoto__Scalars__Sfixed64__tag)
		canoto.AppendFint64(w, c.Sfixed64)
	}
	if c.Bool {
		canoto.Append(w, canoto__Scalars__Bool__tag)
		canoto.AppendBool(w, true)
	}
	if len(c.String) != 0 {
		canoto.Append(w, canoto__Scalars__String__tag)
		canoto.AppendBytes(w, c.String)
	}
	if len(c.Bytes) != 0 {
		canoto.Append(w, canoto__Scalars__Bytes__tag)
		canoto.AppendBytes(w, c.Bytes)
	}
	if fieldSize := c.LargestFieldNumber.CachedCanotoSize(); fieldSize != 0 {
		canoto.Append(w, canoto__Scalars__LargestFieldNumber__tag)
		canoto.AppendInt(w, int64(fieldSize))
		c.LargestFieldNumber.MarshalCanotoInto(w)
	}
	if len(c.RepeatedInt32) != 0 {
		canoto.Append(w, canoto__Scalars__RepeatedInt32__tag)
		canoto.AppendInt(w, int64(c.canotoData.RepeatedInt32Size))
		for _, v := range c.RepeatedInt32 {
			canoto.AppendInt(w, v)
		}
	}
	if len(c.RepeatedInt64) != 0 {
		canoto.Append(w, canoto__Scalars__RepeatedInt64__tag)
		canoto.AppendInt(w, int64(c.canotoData.RepeatedInt64Size))
		for _, v := range c.RepeatedInt64 {
			canoto.AppendInt(w, v)
		}
	}
	if len(c.RepeatedUint32) != 0 {
		canoto.Append(w, canoto__Scalars__RepeatedUint32__tag)
		canoto.AppendInt(w, int64(c.canotoData.RepeatedUint32Size))
		for _, v := range c.RepeatedUint32 {
			canoto.AppendInt(w, v)
		}
	}
	if len(c.RepeatedUint64) != 0 {
		canoto.Append(w, canoto__Scalars__RepeatedUint64__tag)
		canoto.AppendInt(w, int64(c.canotoData.RepeatedUint64Size))
		for _, v := range c.RepeatedUint64 {
			canoto.AppendInt(w, v)
		}
	}
	if len(c.RepeatedSint32) != 0 {
		canoto.Append(w, canoto__Scalars__RepeatedSint32__tag)
		canoto.AppendInt(w, int64(c.canotoData.RepeatedSint32Size))
		for _, v := range c.RepeatedSint32 {
			canoto.AppendSint(w, v)
		}
	}
	if len(c.RepeatedSint64) != 0 {
		canoto.Append(w, canoto__Scalars__RepeatedSint64__tag)
		canoto.AppendInt(w, int64(c.canotoData.RepeatedSint64Size))
		for _, v := range c.RepeatedSint64 {
			canoto.AppendSint(w, v)
		}
	}
	if num := len(c.RepeatedFixed32); num != 0 {
		canoto.Append(w, canoto__Scalars__RepeatedFixed32__tag)
		canoto.AppendInt(w, int64(num * canoto.SizeFint32))
		for _, v := range c.RepeatedFixed32 {
			canoto.AppendFint32(w, v)
		}
	}
	if num := len(c.RepeatedFixed64); num != 0 {
		canoto.Append(w, canoto__Scalars__RepeatedFixed64__tag)
		canoto.AppendInt(w, int64(num * canoto.SizeFint64))
		for _, v := range c.RepeatedFixed64 {
			canoto.AppendFint64(w, v)
		}
	}
	if num := len(c.RepeatedSfixed32); num != 0 {
		canoto.Append(w, canoto__Scalars__RepeatedSfixed32__tag)
		canoto.AppendInt(w, int64(num * canoto.SizeFint32))
		for _, v := range c.RepeatedSfixed32 {
			canoto.AppendFint32(w, v)
		}
	}
	if num := len(c.RepeatedSfixed64); num != 0 {
		canoto.Append(w, canoto__Scalars__RepeatedSfixed64__tag)
		canoto.AppendInt(w, int64(num * canoto.SizeFint64))
		for _, v := range c.RepeatedSfixed64 {
			canoto.AppendFint64(w, v)
		}
	}
	if num := len(c.RepeatedBool); num != 0 {
		canoto.Append(w, canoto__Scalars__RepeatedBool__tag)
		canoto.AppendInt(w, int64(num * canoto.SizeBool))
		for _, v := range c.RepeatedBool {
			canoto.AppendBool(w, v)
		}
	}
}
