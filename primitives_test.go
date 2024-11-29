package canoto

import (
	"encoding/hex"
	"io"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/StephenButtolph/canoto/internal/proto/pb"
)

type validTest[T any] struct {
	hex  string
	want T
}

func (v validTest[_]) Bytes(t *testing.T) []byte {
	bytes, err := hex.DecodeString(v.hex)
	require.NoError(t, err)
	return bytes
}

type invalidTest struct {
	hex  string
	want error
}

func (v invalidTest) Bytes(t *testing.T) []byte {
	bytes, err := hex.DecodeString(v.hex)
	require.NoError(t, err)
	return bytes
}

func FuzzSizeTag(f *testing.F) {
	f.Fuzz(func(t *testing.T, fieldNumber uint32, wireTypeByte byte) {
		wireType := WireType(wireTypeByte)
		if fieldNumber > MaxFieldNumber || !wireType.IsValid() {
			return
		}

		w := &Writer{}
		AppendTag(w, fieldNumber, wireType)

		size := SizeTag(fieldNumber, wireType)
		require.Len(t, w.b, size)
	})
}

func TestReadTag(t *testing.T) {
	type tag struct {
		fieldNumber uint32
		wireType    WireType
	}
	validTests := []validTest[tag]{
		{"00", tag{fieldNumber: 0, wireType: Varint}},
		{"01", tag{fieldNumber: 0, wireType: I64}},
		{"02", tag{fieldNumber: 0, wireType: Len}},
		{"05", tag{fieldNumber: 0, wireType: I32}},
		{"08", tag{fieldNumber: 1, wireType: Varint}},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			gotField, gotType, err := ReadTag(r)
			require.NoError(err)
			require.Equal(test.want, tag{fieldNumber: gotField, wireType: gotType})
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"03", errInvalidWireType},
		{"04", errInvalidWireType},
		{"", io.ErrUnexpectedEOF},
		{"80", io.ErrUnexpectedEOF},
		{"8080", io.ErrUnexpectedEOF},
		{"808080", io.ErrUnexpectedEOF},
		{"80808080", io.ErrUnexpectedEOF},
		{"8080808080", io.ErrUnexpectedEOF},
		{"808080808080", io.ErrUnexpectedEOF},
		{"80808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080", io.ErrUnexpectedEOF},
		{"808080808080808080", io.ErrUnexpectedEOF},
		{"80808080808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080808080", errOverflow},
		{"ffffffffffffffffff02", errOverflow},
		{"8180808080808080808000", errOverflow},
		{"8100", errPaddedZeroes},
		{"818000", errPaddedZeroes},
		{"81808000", errPaddedZeroes},
		{"8180808000", errPaddedZeroes},
		{"818080808000", errPaddedZeroes},
		{"81808080808000", errPaddedZeroes},
		{"8180808080808000", errPaddedZeroes},
		{"818080808080808000", errPaddedZeroes},
		{"81808080808080808000", errPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, _, err := ReadTag(r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendTag(f *testing.F) {
	f.Fuzz(func(t *testing.T, fieldNumber uint32, wireTypeByte byte) {
		wireType := WireType(wireTypeByte)
		if fieldNumber > MaxFieldNumber || !wireType.IsValid() {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendTag(w, fieldNumber, wireType)

		r := &Reader{b: w.b}
		gotFieldNumber, gotWireType, err := ReadTag(r)
		require.NoError(err)
		require.Equal(fieldNumber, gotFieldNumber)
		require.Equal(wireType, gotWireType)
		require.Empty(r.b)
	})
}

func FuzzSizeInt_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		w := &Writer{}
		AppendInt(w, v)

		size := SizeInt(v)
		require.Len(t, w.b, size)
	})
}

func TestReadInt_int32(t *testing.T) {
	validTests := []validTest[int32]{
		{"00", 0},
		{"01", 1},
		{"7f", 0x7f},
		{"8001", 0x7f + 1},
		{"9601", 150},
		{"ff7f", 0x3fff},
		{"808001", 0x3fff + 1},
		{"ffff7f", 0x1fffff},
		{"80808001", 0x1fffff + 1},
		{"ffffff7f", 0xfffffff},
		{"8080808001", 0xfffffff + 1},
		{"ffffffff07", math.MaxInt32},
		{"80808080f8ffffffff01", math.MinInt32},
		{"ffffffffffffffffff01", -1},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadInt[int32](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"80", io.ErrUnexpectedEOF},
		{"8080", io.ErrUnexpectedEOF},
		{"808080", io.ErrUnexpectedEOF},
		{"80808080", io.ErrUnexpectedEOF},
		{"8080808080", io.ErrUnexpectedEOF},
		{"808080808080", io.ErrUnexpectedEOF},
		{"80808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080", io.ErrUnexpectedEOF},
		{"808080808080808080", io.ErrUnexpectedEOF},
		{"80808080808080808080", io.ErrUnexpectedEOF},
		{"ffffffff7f", errOverflow},
		{"808080808001", errOverflow},
		{"ffffffffff7f", errOverflow},
		{"80808080808001", errOverflow},
		{"ffffffffffff7f", errOverflow},
		{"8080808080808001", errOverflow},
		{"ffffffffffffff7f", errOverflow},
		{"808080808080808001", errOverflow},
		{"ffffffffffffffff7f", errOverflow},
		{"8080808080808080808080", errOverflow},
		{"ffffffffffffffffff02", errOverflow},
		{"8180808080808080808000", errOverflow},
		{"8100", errPaddedZeroes},
		{"818000", errPaddedZeroes},
		{"81808000", errPaddedZeroes},
		{"8180808000", errPaddedZeroes},
		{"818080808000", errPaddedZeroes},
		{"81808080808000", errPaddedZeroes},
		{"8180808080808000", errPaddedZeroes},
		{"818080808080808000", errPaddedZeroes},
		{"81808080808080808000", errPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadInt[int32](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendInt_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		require := require.New(t)

		w := &Writer{}
		AppendInt(w, v)

		r := &Reader{b: w.b}
		got, err := ReadInt[int32](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzSizeInt_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		w := &Writer{}
		AppendInt(w, v)

		size := SizeInt(v)
		require.Len(t, w.b, size)
	})
}

func TestReadInt_int64(t *testing.T) {
	validTests := []validTest[int64]{
		{"00", 0},
		{"01", 1},
		{"7f", 0x7f},
		{"8001", 0x7f + 1},
		{"9601", 150},
		{"ff7f", 0x3fff},
		{"808001", 0x3fff + 1},
		{"ffff7f", 0x1fffff},
		{"80808001", 0x1fffff + 1},
		{"ffffff7f", 0xfffffff},
		{"8080808001", 0xfffffff + 1},
		{"ffffffff7f", 0x7ffffffff},
		{"808080808001", 0x7ffffffff + 1},
		{"ffffffffff7f", 0x3ffffffffff},
		{"80808080808001", 0x3ffffffffff + 1},
		{"ffffffffffff7f", 0x1ffffffffffff},
		{"8080808080808001", 0x1ffffffffffff + 1},
		{"ffffffffffffff7f", 0xffffffffffffff},
		{"808080808080808001", 0xffffffffffffff + 1},
		{"ffffffffffffffff7f", math.MaxInt64},
		{"80808080808080808001", math.MinInt64},
		{"ffffffffffffffffff01", -1},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadInt[int64](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"80", io.ErrUnexpectedEOF},
		{"8080", io.ErrUnexpectedEOF},
		{"808080", io.ErrUnexpectedEOF},
		{"80808080", io.ErrUnexpectedEOF},
		{"8080808080", io.ErrUnexpectedEOF},
		{"808080808080", io.ErrUnexpectedEOF},
		{"80808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080", io.ErrUnexpectedEOF},
		{"808080808080808080", io.ErrUnexpectedEOF},
		{"80808080808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080808080", errOverflow},
		{"ffffffffffffffffff02", errOverflow},
		{"8180808080808080808000", errOverflow},
		{"8100", errPaddedZeroes},
		{"818000", errPaddedZeroes},
		{"81808000", errPaddedZeroes},
		{"8180808000", errPaddedZeroes},
		{"818080808000", errPaddedZeroes},
		{"81808080808000", errPaddedZeroes},
		{"8180808080808000", errPaddedZeroes},
		{"818080808080808000", errPaddedZeroes},
		{"81808080808080808000", errPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadInt[int64](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendInt_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		require := require.New(t)

		w := &Writer{}
		AppendInt(w, v)

		r := &Reader{b: w.b}
		got, err := ReadInt[int64](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzSizeInt_uint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		w := &Writer{}
		AppendInt(w, v)

		size := SizeInt(v)
		require.Len(t, w.b, size)
	})
}

func TestReadInt_uint32(t *testing.T) {
	validTests := []validTest[uint32]{
		{"00", 0},
		{"01", 1},
		{"7f", 0x7f},
		{"8001", 0x7f + 1},
		{"9601", 150},
		{"ff7f", 0x3fff},
		{"808001", 0x3fff + 1},
		{"ffff7f", 0x1fffff},
		{"80808001", 0x1fffff + 1},
		{"ffffff7f", 0xfffffff},
		{"8080808001", 0xfffffff + 1},
		{"ffffffff0f", math.MaxUint32},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadInt[uint32](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"80", io.ErrUnexpectedEOF},
		{"8080", io.ErrUnexpectedEOF},
		{"808080", io.ErrUnexpectedEOF},
		{"80808080", io.ErrUnexpectedEOF},
		{"8080808080", io.ErrUnexpectedEOF},
		{"808080808080", io.ErrUnexpectedEOF},
		{"80808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080", io.ErrUnexpectedEOF},
		{"808080808080808080", io.ErrUnexpectedEOF},
		{"80808080808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080808080", errOverflow},
		{"8080808080808080808080", errOverflow},
		{"ffffffff10", errOverflow},
		{"8180808080808080808000", errOverflow},
		{"8100", errPaddedZeroes},
		{"818000", errPaddedZeroes},
		{"81808000", errPaddedZeroes},
		{"8180808000", errPaddedZeroes},
		{"818080808000", errPaddedZeroes},
		{"81808080808000", errPaddedZeroes},
		{"8180808080808000", errPaddedZeroes},
		{"818080808080808000", errPaddedZeroes},
		{"81808080808080808000", errPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadInt[uint32](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendInt_uint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		require := require.New(t)

		w := &Writer{}
		AppendInt(w, v)

		r := &Reader{b: w.b}
		got, err := ReadInt[uint32](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzSizeInt_uint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		w := &Writer{}
		AppendInt(w, v)

		size := SizeInt(v)
		require.Len(t, w.b, size)
	})
}

func TestReadInt_uint64(t *testing.T) {
	validTests := []validTest[uint64]{
		{"00", 0},
		{"01", 1},
		{"7f", 0x7f},
		{"8001", 0x7f + 1},
		{"9601", 150},
		{"ff7f", 0x3fff},
		{"808001", 0x3fff + 1},
		{"ffff7f", 0x1fffff},
		{"80808001", 0x1fffff + 1},
		{"ffffff7f", 0xfffffff},
		{"8080808001", 0xfffffff + 1},
		{"ffffffff7f", 0x7ffffffff},
		{"808080808001", 0x7ffffffff + 1},
		{"ffffffffff7f", 0x3ffffffffff},
		{"80808080808001", 0x3ffffffffff + 1},
		{"ffffffffffff7f", 0x1ffffffffffff},
		{"8080808080808001", 0x1ffffffffffff + 1},
		{"ffffffffffffff7f", 0xffffffffffffff},
		{"808080808080808001", 0xffffffffffffff + 1},
		{"ffffffffffffffff7f", 0x7fffffffffffffff},
		{"80808080808080808001", 0x7fffffffffffffff + 1},
		{"ffffffffffffffffff01", math.MaxUint64},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadInt[uint64](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"80", io.ErrUnexpectedEOF},
		{"8080", io.ErrUnexpectedEOF},
		{"808080", io.ErrUnexpectedEOF},
		{"80808080", io.ErrUnexpectedEOF},
		{"8080808080", io.ErrUnexpectedEOF},
		{"808080808080", io.ErrUnexpectedEOF},
		{"80808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080", io.ErrUnexpectedEOF},
		{"808080808080808080", io.ErrUnexpectedEOF},
		{"80808080808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080808080", errOverflow},
		{"ffffffffffffffffff02", errOverflow},
		{"8180808080808080808000", errOverflow},
		{"8100", errPaddedZeroes},
		{"818000", errPaddedZeroes},
		{"81808000", errPaddedZeroes},
		{"8180808000", errPaddedZeroes},
		{"818080808000", errPaddedZeroes},
		{"81808080808000", errPaddedZeroes},
		{"8180808080808000", errPaddedZeroes},
		{"818080808080808000", errPaddedZeroes},
		{"81808080808080808000", errPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadInt[uint64](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendInt_uint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		require := require.New(t)

		w := &Writer{}
		AppendInt(w, v)

		r := &Reader{b: w.b}
		got, err := ReadInt[uint64](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzSizeSint_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		w := &Writer{}
		AppendSint(w, v)

		size := SizeSint(v)
		require.Len(t, w.b, size)
	})
}

func TestReadSint_int32(t *testing.T) {
	validTests := []validTest[int32]{
		{"00", 0},
		{"01", -1},
		{"02", +1},
		{"03", -2},
		{"04", +2},
		{"05", -3},
		{"06", +3},
		{"06", +3},
		{"faffffff0f", math.MaxInt32 - 2},
		{"fbffffff0f", math.MinInt32 + 2},
		{"fcffffff0f", math.MaxInt32 - 1},
		{"fdffffff0f", math.MinInt32 + 1},
		{"feffffff0f", math.MaxInt32},
		{"ffffffff0f", math.MinInt32},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadSint[int32](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"80", io.ErrUnexpectedEOF},
		{"8080", io.ErrUnexpectedEOF},
		{"808080", io.ErrUnexpectedEOF},
		{"80808080", io.ErrUnexpectedEOF},
		{"8080808080", io.ErrUnexpectedEOF},
		{"808080808080", io.ErrUnexpectedEOF},
		{"80808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080", io.ErrUnexpectedEOF},
		{"808080808080808080", io.ErrUnexpectedEOF},
		{"80808080808080808080", io.ErrUnexpectedEOF},
		{"ffffffff10", errOverflow},
		{"8080808080808080808080", errOverflow},
		{"ffffffffffffffffff02", errOverflow},
		{"8180808080808080808000", errOverflow},
		{"8100", errPaddedZeroes},
		{"818000", errPaddedZeroes},
		{"81808000", errPaddedZeroes},
		{"8180808000", errPaddedZeroes},
		{"818080808000", errPaddedZeroes},
		{"81808080808000", errPaddedZeroes},
		{"8180808080808000", errPaddedZeroes},
		{"818080808080808000", errPaddedZeroes},
		{"81808080808080808000", errPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadSint[int32](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendSint_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		require := require.New(t)

		w := &Writer{}
		AppendSint(w, v)

		r := &Reader{b: w.b}
		got, err := ReadSint[int32](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzSizeSint_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		w := &Writer{}
		AppendSint(w, v)

		size := SizeSint(v)
		require.Len(t, w.b, size)
	})
}

func TestReadSint_int64(t *testing.T) {
	validTests := []validTest[int64]{
		{"00", 0},
		{"01", -1},
		{"02", +1},
		{"03", -2},
		{"04", +2},
		{"05", -3},
		{"06", +3},
		{"06", +3},
		{"faffffffffffffffff01", math.MaxInt64 - 2},
		{"fbffffffffffffffff01", math.MinInt64 + 2},
		{"fcffffffffffffffff01", math.MaxInt64 - 1},
		{"fdffffffffffffffff01", math.MinInt64 + 1},
		{"feffffffffffffffff01", math.MaxInt64},
		{"ffffffffffffffffff01", math.MinInt64},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadSint[int64](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"80", io.ErrUnexpectedEOF},
		{"8080", io.ErrUnexpectedEOF},
		{"808080", io.ErrUnexpectedEOF},
		{"80808080", io.ErrUnexpectedEOF},
		{"8080808080", io.ErrUnexpectedEOF},
		{"808080808080", io.ErrUnexpectedEOF},
		{"80808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080", io.ErrUnexpectedEOF},
		{"808080808080808080", io.ErrUnexpectedEOF},
		{"80808080808080808080", io.ErrUnexpectedEOF},
		{"8080808080808080808080", errOverflow},
		{"ffffffffffffffffff02", errOverflow},
		{"8180808080808080808000", errOverflow},
		{"8100", errPaddedZeroes},
		{"818000", errPaddedZeroes},
		{"81808000", errPaddedZeroes},
		{"8180808000", errPaddedZeroes},
		{"818080808000", errPaddedZeroes},
		{"81808080808000", errPaddedZeroes},
		{"8180808080808000", errPaddedZeroes},
		{"818080808080808000", errPaddedZeroes},
		{"81808080808080808000", errPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadSint[int64](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendSint_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		require := require.New(t)

		w := &Writer{}
		AppendSint(w, v)

		r := &Reader{b: w.b}
		got, err := ReadSint[int64](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzSizeFint32_uint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		w := &Writer{}
		AppendFint32(w, v)
		require.Len(t, w.b, SizeFint32)
	})
}

func TestReadFint32_uint32(t *testing.T) {
	validTests := []validTest[uint32]{
		{"00000000", 0},
		{"01000000", 1},
		{"ffffffff", math.MaxUint32},
		{"c3d2e1f0", 0xf0e1d2c3},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadFint32[uint32](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"00", io.ErrUnexpectedEOF},
		{"0000", io.ErrUnexpectedEOF},
		{"000000", io.ErrUnexpectedEOF},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadFint32[uint32](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendFint32_uint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		require := require.New(t)

		w := &Writer{}
		AppendFint32(w, v)

		r := &Reader{b: w.b}
		got, err := ReadFint32[uint32](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzSizeFint32_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		w := &Writer{}
		AppendFint32(w, v)
		require.Len(t, w.b, SizeFint32)
	})
}

func TestReadFint32_int32(t *testing.T) {
	validTests := []validTest[int32]{
		{"00000080", math.MinInt32},
		{"ffffffff", -1},
		{"00000000", 0},
		{"01000000", 1},
		{"ffffff7f", math.MaxInt32},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadFint32[int32](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"00", io.ErrUnexpectedEOF},
		{"0000", io.ErrUnexpectedEOF},
		{"000000", io.ErrUnexpectedEOF},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadFint32[int32](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendFint32_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		require := require.New(t)

		w := &Writer{}
		AppendFint32(w, v)

		r := &Reader{b: w.b}
		got, err := ReadFint32[int32](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzSizeFint64_uint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		w := &Writer{}
		AppendFint64(w, v)
		require.Len(t, w.b, SizeFint64)
	})
}

func TestReadFint64_uint64(t *testing.T) {
	validTests := []validTest[uint64]{
		{"0000000000000000", 0},
		{"0100000000000000", 1},
		{"ffffffffffffffff", math.MaxUint64},
		{"8796a5b4c3d2e1f0", 0xf0e1d2c3b4a59687},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadFint64[uint64](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"00", io.ErrUnexpectedEOF},
		{"0000", io.ErrUnexpectedEOF},
		{"000000", io.ErrUnexpectedEOF},
		{"00000000", io.ErrUnexpectedEOF},
		{"0000000000", io.ErrUnexpectedEOF},
		{"000000000000", io.ErrUnexpectedEOF},
		{"00000000000000", io.ErrUnexpectedEOF},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadFint64[uint64](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendFint64_uint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		require := require.New(t)

		w := &Writer{}
		AppendFint64(w, v)

		r := &Reader{b: w.b}
		got, err := ReadFint64[uint64](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzSizeFint64_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		w := &Writer{}
		AppendFint64(w, v)
		require.Len(t, w.b, SizeFint64)
	})
}

func TestReadFint64_int64(t *testing.T) {
	validTests := []validTest[int64]{
		{"0000000000000080", math.MinInt64},
		{"ffffffffffffffff", -1},
		{"0000000000000000", 0},
		{"0100000000000000", 1},
		{"ffffffffffffff7f", math.MaxInt64},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadFint64[int64](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"00", io.ErrUnexpectedEOF},
		{"0000", io.ErrUnexpectedEOF},
		{"000000", io.ErrUnexpectedEOF},
		{"00000000", io.ErrUnexpectedEOF},
		{"0000000000", io.ErrUnexpectedEOF},
		{"000000000000", io.ErrUnexpectedEOF},
		{"00000000000000", io.ErrUnexpectedEOF},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadFint64[int64](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendFint64_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		require := require.New(t)

		w := &Writer{}
		AppendFint64(w, v)

		r := &Reader{b: w.b}
		got, err := ReadFint64[int64](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzSizeBool(f *testing.F) {
	f.Fuzz(func(t *testing.T, v bool) {
		w := &Writer{}
		AppendBool(w, v)
		require.Len(t, w.b, SizeBool)
	})
}

func TestReadBool(t *testing.T) {
	validTests := []validTest[bool]{
		{"00", false},
		{"01", true},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadBool(r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"02", errInvalidBool},
		{"ff", errInvalidBool},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadBool(r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendBool(f *testing.F) {
	f.Fuzz(func(t *testing.T, v bool) {
		require := require.New(t)

		w := &Writer{}
		AppendBool(w, v)

		r := &Reader{b: w.b}
		got, err := ReadBool(r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzSizeBytes_string(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		w := &Writer{}
		AppendBytes(w, v)

		size := SizeBytes(v)
		require.Len(t, w.b, size)
	})
}

func FuzzSizeBytes_bytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, v []byte) {
		w := &Writer{}
		AppendBytes(w, v)

		size := SizeBytes(v)
		require.Len(t, w.b, size)
	})
}

func TestReadBytes_string(t *testing.T) {
	validTests := []validTest[string]{
		{"00", ""},
		{"0774657374696e67", "testing"},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadBytes[string](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"870074657374696e67", errPaddedZeroes},
		{"ffffffffffffffffff01", errInvalidLength},
		{"01", io.ErrUnexpectedEOF},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadBytes[string](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReadBytes_bytes(t *testing.T) {
	validTests := []validTest[[]byte]{
		{"00", []byte{}},
		{"0774657374696e67", []byte("testing")},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := ReadBytes[[]byte](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.b)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"870074657374696e67", errPaddedZeroes},
		{"ffffffffffffffffff01", errInvalidLength},
		{"01", io.ErrUnexpectedEOF},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{b: test.Bytes(t)}
			_, err := ReadBytes[[]byte](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendBytes_string(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		require := require.New(t)

		w := &Writer{}
		AppendBytes(w, v)

		r := &Reader{b: w.b}
		got, err := ReadBytes[string](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzAppendBytes_bytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, v []byte) {
		require := require.New(t)

		w := &Writer{}
		AppendBytes(w, v)

		r := &Reader{b: w.b}
		got, err := ReadBytes[[]byte](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func TestAppend_ProtoCompatibility(t *testing.T) {
	tests := []struct {
		name  string
		proto protoreflect.ProtoMessage
		f     func(*Writer)
	}{
		{
			name: "int32",
			proto: &pb.Scalars{
				Int32: 128,
			},
			f: func(w *Writer) {
				AppendTag(w, 1, Varint)
				AppendInt[int32](w, 128)
			},
		},
		{
			name: "int64",
			proto: &pb.Scalars{
				Int64: 259,
			},
			f: func(w *Writer) {
				AppendTag(w, 2, Varint)
				AppendInt[int64](w, 259)
			},
		},
		{
			name: "uint32",
			proto: &pb.Scalars{
				Uint32: 1234,
			},
			f: func(w *Writer) {
				AppendTag(w, 3, Varint)
				AppendInt[uint32](w, 1234)
			},
		},
		{
			name: "uint64",
			proto: &pb.Scalars{
				Uint64: 2938567,
			},
			f: func(w *Writer) {
				AppendTag(w, 4, Varint)
				AppendInt[uint64](w, 2938567)
			},
		},
		{
			name: "sint32",
			proto: &pb.Scalars{
				Sint32: -2136745,
			},
			f: func(w *Writer) {
				AppendTag(w, 5, Varint)
				AppendSint[int32](w, -2136745)
			},
		},
		{
			name: "sint64",
			proto: &pb.Scalars{
				Sint64: -9287364,
			},
			f: func(w *Writer) {
				AppendTag(w, 6, Varint)
				AppendSint[int64](w, -9287364)
			},
		},
		{
			name: "fixed32",
			proto: &pb.Scalars{
				Fixed32: 876254,
			},
			f: func(w *Writer) {
				AppendTag(w, 7, I32)
				AppendFint32[uint32](w, 876254)
			},
		},
		{
			name: "fixed64",
			proto: &pb.Scalars{
				Fixed64: 328137645632,
			},
			f: func(w *Writer) {
				AppendTag(w, 8, I64)
				AppendFint64[uint64](w, 328137645632)
			},
		},
		{
			name: "sfixed32",
			proto: &pb.Scalars{
				Sfixed32: -123463246,
			},
			f: func(w *Writer) {
				AppendTag(w, 9, I32)
				AppendFint32[int32](w, -123463246)
			},
		},
		{
			name: "sfixed64",
			proto: &pb.Scalars{
				Sfixed64: -8762135423,
			},
			f: func(w *Writer) {
				AppendTag(w, 10, I64)
				AppendFint64[int64](w, -8762135423)
			},
		},
		{
			name: "bool",
			proto: &pb.Scalars{
				Bool: true,
			},
			f: func(w *Writer) {
				AppendTag(w, 11, Varint)
				AppendBool(w, true)
			},
		},
		{
			name: "string",
			proto: &pb.Scalars{
				String_: "hi mom!",
			},
			f: func(w *Writer) {
				AppendTag(w, 12, Len)
				AppendBytes(w, "hi mom!")
			},
		},
		{
			name: "bytes",
			proto: &pb.Scalars{
				Bytes: []byte("hi dad!"),
			},
			f: func(w *Writer) {
				AppendTag(w, 13, Len)
				AppendBytes(w, []byte("hi dad!"))
			},
		},
		{
			name: "largest field number",
			proto: &pb.LargestFieldNumber{
				Int32: 1,
			},
			f: func(w *Writer) {
				AppendTag(w, MaxFieldNumber, Varint)
				AppendInt[int32](w, 1)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pbBytes, err := proto.Marshal(test.proto)
			require.NoError(t, err)

			w := &Writer{}
			test.f(w)
			require.Equal(t, pbBytes, w.b)
		})
	}
}
