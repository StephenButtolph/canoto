package canoto

import (
	"encoding/hex"
	"io"
	"math"
	"testing"
	"unicode/utf8"

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

func TestReadTag(t *testing.T) {
	type tag struct {
		fieldNumber uint32
		wireType    WireType
	}
	validTests := []validTest[tag]{
		{"01", tag{fieldNumber: 0, wireType: I64}},
		{"02", tag{fieldNumber: 0, wireType: Len}},
		{"05", tag{fieldNumber: 0, wireType: I32}},
		{"08", tag{fieldNumber: 1, wireType: Varint}},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{B: test.Bytes(t)}
			gotField, gotType, err := ReadTag(r)
			require.NoError(err)
			require.Equal(test.want, tag{fieldNumber: gotField, wireType: gotType})
			require.Empty(r.B)
		})
	}

	invalidTests := []invalidTest{
		{"00", errPaddedZeroes},
		{"03", ErrInvalidWireType},
		{"04", ErrInvalidWireType},
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
			r := &Reader{B: test.Bytes(t)}
			_, _, err := ReadTag(r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzSizeInt_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		if v == 0 {
			return
		}

		w := &Writer{}
		AppendInt(w, v)

		size := SizeInt(v)
		require.Len(t, w.B, size)
	})
}

func TestReadInt_int32(t *testing.T) {
	validTests := []validTest[int32]{
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

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadInt[int32](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
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
		{"00", errPaddedZeroes},
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
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadInt[int32](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendInt_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		if v == 0 {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendInt(w, v)

		r := &Reader{B: w.B}
		got, err := ReadInt[int32](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func FuzzSizeInt_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		if v == 0 {
			return
		}

		w := &Writer{}
		AppendInt(w, v)

		size := SizeInt(v)
		require.Len(t, w.B, size)
	})
}

func TestReadInt_int64(t *testing.T) {
	validTests := []validTest[int64]{
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

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadInt[int64](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
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
		{"00", errPaddedZeroes},
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
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadInt[int64](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendInt_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		if v == 0 {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendInt(w, v)

		r := &Reader{B: w.B}
		got, err := ReadInt[int64](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func FuzzSizeInt_uint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		if v == 0 {
			return
		}

		w := &Writer{}
		AppendInt(w, v)

		size := SizeInt(v)
		require.Len(t, w.B, size)
	})
}

func TestReadInt_uint32(t *testing.T) {
	validTests := []validTest[uint32]{
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

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadInt[uint32](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
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
		{"00", errPaddedZeroes},
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
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadInt[uint32](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendInt_uint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		if v == 0 {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendInt(w, v)

		r := &Reader{B: w.B}
		got, err := ReadInt[uint32](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func FuzzSizeInt_uint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		if v == 0 {
			return
		}

		w := &Writer{}
		AppendInt(w, v)

		size := SizeInt(v)
		require.Len(t, w.B, size)
	})
}

func TestReadInt_uint64(t *testing.T) {
	validTests := []validTest[uint64]{
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

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadInt[uint64](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
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
		{"00", errPaddedZeroes},
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
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadInt[uint64](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendInt_uint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		if v == 0 {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendInt(w, v)

		r := &Reader{B: w.B}
		got, err := ReadInt[uint64](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func FuzzSizeSint_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		if v == 0 {
			return
		}

		w := &Writer{}
		AppendSint(w, v)

		size := SizeSint(v)
		require.Len(t, w.B, size)
	})
}

func TestReadSint_int32(t *testing.T) {
	validTests := []validTest[int32]{
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

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadSint[int32](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
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
		{"00", errPaddedZeroes},
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
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadSint[int32](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendSint_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		if v == 0 {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendSint(w, v)

		r := &Reader{B: w.B}
		got, err := ReadSint[int32](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func FuzzSizeSint_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		if v == 0 {
			return
		}

		w := &Writer{}
		AppendSint(w, v)

		size := SizeSint(v)
		require.Len(t, w.B, size)
	})
}

func TestReadSint_int64(t *testing.T) {
	validTests := []validTest[int64]{
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

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadSint[int64](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
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
		{"00", errPaddedZeroes},
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
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadSint[int64](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendSint_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		if v == 0 {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendSint(w, v)

		r := &Reader{B: w.B}
		got, err := ReadSint[int64](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func FuzzSizeFint32_uint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		if v == 0 {
			return
		}

		w := &Writer{}
		AppendFint32(w, v)
		require.Len(t, w.B, SizeFint32)
	})
}

func TestReadFint32_uint32(t *testing.T) {
	validTests := []validTest[uint32]{
		{"01000000", 1},
		{"ffffffff", math.MaxUint32},
		{"c3d2e1f0", 0xf0e1d2c3},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadFint32[uint32](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"00", io.ErrUnexpectedEOF},
		{"0000", io.ErrUnexpectedEOF},
		{"000000", io.ErrUnexpectedEOF},
		{"00000000", errZeroValue},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadFint32[uint32](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendFint32_uint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		if v == 0 {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendFint32(w, v)

		r := &Reader{B: w.B}
		got, err := ReadFint32[uint32](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func FuzzSizeFint32_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		if v == 0 {
			return
		}

		w := &Writer{}
		AppendFint32(w, v)
		require.Len(t, w.B, SizeFint32)
	})
}

func TestReadFint32_int32(t *testing.T) {
	validTests := []validTest[int32]{
		{"00000080", math.MinInt32},
		{"ffffffff", -1},
		{"01000000", 1},
		{"ffffff7f", math.MaxInt32},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadFint32[int32](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"00", io.ErrUnexpectedEOF},
		{"0000", io.ErrUnexpectedEOF},
		{"000000", io.ErrUnexpectedEOF},
		{"00000000", errZeroValue},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadFint32[int32](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendFint32_int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		if v == 0 {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendFint32(w, v)

		r := &Reader{B: w.B}
		got, err := ReadFint32[int32](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func FuzzSizeFint64_uint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		if v == 0 {
			return
		}

		w := &Writer{}
		AppendFint64(w, v)
		require.Len(t, w.B, SizeFint64)
	})
}

func TestReadFint64_uint64(t *testing.T) {
	validTests := []validTest[uint64]{
		{"0100000000000000", 1},
		{"ffffffffffffffff", math.MaxUint64},
		{"8796a5b4c3d2e1f0", 0xf0e1d2c3b4a59687},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadFint64[uint64](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
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
		{"0000000000000000", errZeroValue},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadFint64[uint64](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendFint64_uint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		if v == 0 {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendFint64(w, v)

		r := &Reader{B: w.B}
		got, err := ReadFint64[uint64](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func FuzzSizeFint64_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		if v == 0 {
			return
		}

		w := &Writer{}
		AppendFint64(w, v)
		require.Len(t, w.B, SizeFint64)
	})
}

func TestReadFint64_int64(t *testing.T) {
	validTests := []validTest[int64]{
		{"0000000000000080", math.MinInt64},
		{"ffffffffffffffff", -1},
		{"0100000000000000", 1},
		{"ffffffffffffff7f", math.MaxInt64},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadFint64[int64](r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
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
		{"0000000000000000", errZeroValue},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadFint64[int64](r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendFint64_int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		if v == 0 {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendFint64(w, v)

		r := &Reader{B: w.B}
		got, err := ReadFint64[int64](r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func TestSizeBool(t *testing.T) {
	w := &Writer{}
	AppendTrue(w)
	require.Len(t, w.B, SizeBool)
}

func TestReadTrue(t *testing.T) {
	t.Run("01", func(t *testing.T) {
		require := require.New(t)

		r := &Reader{B: []byte{trueByte}}
		err := ReadTrue(r)
		require.NoError(err)
		require.Empty(r.B)
	})

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"00", errInvalidBool},
		{"02", errInvalidBool},
		{"ff", errInvalidBool},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadTrue(r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestAppendTrue(t *testing.T) {
	require := require.New(t)

	w := &Writer{}
	AppendTrue(w)

	r := &Reader{B: w.B}
	require.NoError(ReadTrue(r))
	require.Empty(r.B)
}

func FuzzSizeBytes_string(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		if len(v) == 0 {
			return
		}

		w := &Writer{}
		AppendBytes(w, v)

		size := SizeBytes(v)
		require.Len(t, w.B, size)
	})
}

func FuzzSizeBytes_bytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, v []byte) {
		if len(v) == 0 {
			return
		}

		w := &Writer{}
		AppendBytes(w, v)

		size := SizeBytes(v)
		require.Len(t, w.B, size)
	})
}

func TestReadString(t *testing.T) {
	validTests := []validTest[string]{
		{"0774657374696e67", "testing"},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadString(r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"00", errPaddedZeroes},
		{"870074657374696e67", errPaddedZeroes},
		{"ffffffffffffffffff01", errInvalidLength},
		{"01", io.ErrUnexpectedEOF},
		{"01C2", errStringNotUTF8},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadString(r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReadBytes(t *testing.T) {
	validTests := []validTest[[]byte]{
		{"0774657374696e67", []byte("testing")},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{B: test.Bytes(t)}
			got, err := ReadBytes(r)
			require.NoError(err)
			require.Equal(test.want, got)
			require.Empty(r.B)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"00", errPaddedZeroes},
		{"870074657374696e67", errPaddedZeroes},
		{"ffffffffffffffffff01", errInvalidLength},
		{"01", io.ErrUnexpectedEOF},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			_, err := ReadBytes(r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendBytes_string(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		if len(v) == 0 || !utf8.ValidString(v) {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendBytes(w, v)

		r := &Reader{B: w.B}
		got, err := ReadString(r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func FuzzAppendBytes_bytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, v []byte) {
		if len(v) == 0 {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendBytes(w, v)

		r := &Reader{B: w.B}
		got, err := ReadBytes(r)
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.B)
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
				Append(w, Tag(1, Varint))
				AppendInt[int32](w, 128)
			},
		},
		{
			name: "int64",
			proto: &pb.Scalars{
				Int64: 259,
			},
			f: func(w *Writer) {
				Append(w, Tag(2, Varint))
				AppendInt[int64](w, 259)
			},
		},
		{
			name: "uint32",
			proto: &pb.Scalars{
				Uint32: 1234,
			},
			f: func(w *Writer) {
				Append(w, Tag(3, Varint))
				AppendInt[uint32](w, 1234)
			},
		},
		{
			name: "uint64",
			proto: &pb.Scalars{
				Uint64: 2938567,
			},
			f: func(w *Writer) {
				Append(w, Tag(4, Varint))
				AppendInt[uint64](w, 2938567)
			},
		},
		{
			name: "sint32",
			proto: &pb.Scalars{
				Sint32: -2136745,
			},
			f: func(w *Writer) {
				Append(w, Tag(5, Varint))
				AppendSint[int32](w, -2136745)
			},
		},
		{
			name: "sint64",
			proto: &pb.Scalars{
				Sint64: -9287364,
			},
			f: func(w *Writer) {
				Append(w, Tag(6, Varint))
				AppendSint[int64](w, -9287364)
			},
		},
		{
			name: "fixed32",
			proto: &pb.Scalars{
				Fixed32: 876254,
			},
			f: func(w *Writer) {
				Append(w, Tag(7, I32))
				AppendFint32[uint32](w, 876254)
			},
		},
		{
			name: "fixed64",
			proto: &pb.Scalars{
				Fixed64: 328137645632,
			},
			f: func(w *Writer) {
				Append(w, Tag(8, I64))
				AppendFint64[uint64](w, 328137645632)
			},
		},
		{
			name: "sfixed32",
			proto: &pb.Scalars{
				Sfixed32: -123463246,
			},
			f: func(w *Writer) {
				Append(w, Tag(9, I32))
				AppendFint32[int32](w, -123463246)
			},
		},
		{
			name: "sfixed64",
			proto: &pb.Scalars{
				Sfixed64: -8762135423,
			},
			f: func(w *Writer) {
				Append(w, Tag(10, I64))
				AppendFint64[int64](w, -8762135423)
			},
		},
		{
			name: "bool",
			proto: &pb.Scalars{
				Bool: true,
			},
			f: func(w *Writer) {
				Append(w, Tag(11, Varint))
				AppendTrue(w)
			},
		},
		{
			name: "string",
			proto: &pb.Scalars{
				String_: "hi mom!",
			},
			f: func(w *Writer) {
				Append(w, Tag(12, Len))
				AppendBytes(w, "hi mom!")
			},
		},
		{
			name: "bytes",
			proto: &pb.Scalars{
				Bytes: []byte("hi dad!"),
			},
			f: func(w *Writer) {
				Append(w, Tag(13, Len))
				AppendBytes(w, []byte("hi dad!"))
			},
		},
		{
			name: "largest field number",
			proto: &pb.LargestFieldNumber{
				Int32: 1,
			},
			f: func(w *Writer) {
				Append(w, Tag(MaxFieldNumber, Varint))
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
			require.Equal(t, pbBytes, w.B)
		})
	}
}
