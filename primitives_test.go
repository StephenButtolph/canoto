package canoto_test

import (
	"encoding/hex"
	"io"
	"math"
	"strconv"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
	"github.com/thepudds/fzgen/fuzzer"

	. "github.com/StephenButtolph/canoto"
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
		{"00", tag{fieldNumber: 0, wireType: Varint}},
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
		{"8080808080808080808080", ErrOverflow},
		{"ffffffffffffffffff02", ErrOverflow},
		{"8180808080808080808000", ErrOverflow},
		{"8100", ErrPaddedZeroes},
		{"818000", ErrPaddedZeroes},
		{"81808000", ErrPaddedZeroes},
		{"8180808000", ErrPaddedZeroes},
		{"818080808000", ErrPaddedZeroes},
		{"81808080808000", ErrPaddedZeroes},
		{"8180808080808000", ErrPaddedZeroes},
		{"818080808080808000", ErrPaddedZeroes},
		{"81808080808080808000", ErrPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			_, _, err := ReadTag(r)
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzSizeInt_int8(f *testing.F)   { f.Fuzz(testSizeInt[int8]) }
func FuzzSizeInt_uint8(f *testing.F)  { f.Fuzz(testSizeInt[uint8]) }
func FuzzSizeInt_int16(f *testing.F)  { f.Fuzz(testSizeInt[int16]) }
func FuzzSizeInt_uint16(f *testing.F) { f.Fuzz(testSizeInt[uint16]) }
func FuzzSizeInt_int32(f *testing.F)  { f.Fuzz(testSizeInt[int32]) }
func FuzzSizeInt_uint32(f *testing.F) { f.Fuzz(testSizeInt[uint32]) }
func FuzzSizeInt_int64(f *testing.F)  { f.Fuzz(testSizeInt[int64]) }
func FuzzSizeInt_uint64(f *testing.F) { f.Fuzz(testSizeInt[uint64]) }

func testSizeInt[T Int](t *testing.T, v T) {
	w := &Writer{}
	AppendInt(w, v)

	size := SizeInt(v)
	require.Len(t, w.B, size)
}

func FuzzCountInts_int8(f *testing.F)   { f.Fuzz(testCountInts[int8]) }
func FuzzCountInts_uint8(f *testing.F)  { f.Fuzz(testCountInts[uint8]) }
func FuzzCountInts_int16(f *testing.F)  { f.Fuzz(testCountInts[int16]) }
func FuzzCountInts_uint16(f *testing.F) { f.Fuzz(testCountInts[uint16]) }
func FuzzCountInts_int32(f *testing.F)  { f.Fuzz(testCountInts[int32]) }
func FuzzCountInts_uint32(f *testing.F) { f.Fuzz(testCountInts[uint32]) }
func FuzzCountInts_int64(f *testing.F)  { f.Fuzz(testCountInts[int64]) }
func FuzzCountInts_uint64(f *testing.F) { f.Fuzz(testCountInts[uint64]) }

func testCountInts[T Int](t *testing.T, data []byte) {
	require := require.New(t)

	var nums []T
	fz := fuzzer.NewFuzzer(data)
	fz.Fill(&nums)

	w := &Writer{}
	for _, num := range nums {
		AppendInt(w, num)
	}

	count := CountInts(w.B)
	require.Len(nums, count)
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

			r := &Reader{B: test.Bytes(t)}
			var got int32
			require.NoError(ReadInt(r, &got))
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
		{"ffffffff7f", ErrOverflow},
		{"808080808001", ErrOverflow},
		{"ffffffffff7f", ErrOverflow},
		{"80808080808001", ErrOverflow},
		{"ffffffffffff7f", ErrOverflow},
		{"8080808080808001", ErrOverflow},
		{"ffffffffffffff7f", ErrOverflow},
		{"808080808080808001", ErrOverflow},
		{"ffffffffffffffff7f", ErrOverflow},
		{"8080808080808080808080", ErrOverflow},
		{"ffffffffffffffffff02", ErrOverflow},
		{"8180808080808080808000", ErrOverflow},
		{"8100", ErrPaddedZeroes},
		{"818000", ErrPaddedZeroes},
		{"81808000", ErrPaddedZeroes},
		{"8180808000", ErrPaddedZeroes},
		{"818080808000", ErrPaddedZeroes},
		{"81808080808000", ErrPaddedZeroes},
		{"8180808080808000", ErrPaddedZeroes},
		{"818080808080808000", ErrPaddedZeroes},
		{"81808080808080808000", ErrPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadInt(r, new(int32))
			require.ErrorIs(t, err, test.want)
		})
	}
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

			r := &Reader{B: test.Bytes(t)}
			var got uint32
			require.NoError(ReadInt(r, &got))
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
		{"8080808080808080808080", ErrOverflow},
		{"8080808080808080808080", ErrOverflow},
		{"ffffffff10", ErrOverflow},
		{"8180808080808080808000", ErrOverflow},
		{"8100", ErrPaddedZeroes},
		{"818000", ErrPaddedZeroes},
		{"81808000", ErrPaddedZeroes},
		{"8180808000", ErrPaddedZeroes},
		{"818080808000", ErrPaddedZeroes},
		{"81808080808000", ErrPaddedZeroes},
		{"8180808080808000", ErrPaddedZeroes},
		{"818080808080808000", ErrPaddedZeroes},
		{"81808080808080808000", ErrPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadInt(r, new(uint32))
			require.ErrorIs(t, err, test.want)
		})
	}
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

			r := &Reader{B: test.Bytes(t)}
			var got int64
			require.NoError(ReadInt(r, &got))
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
		{"8080808080808080808080", ErrOverflow},
		{"ffffffffffffffffff02", ErrOverflow},
		{"8180808080808080808000", ErrOverflow},
		{"8100", ErrPaddedZeroes},
		{"818000", ErrPaddedZeroes},
		{"81808000", ErrPaddedZeroes},
		{"8180808000", ErrPaddedZeroes},
		{"818080808000", ErrPaddedZeroes},
		{"81808080808000", ErrPaddedZeroes},
		{"8180808080808000", ErrPaddedZeroes},
		{"818080808080808000", ErrPaddedZeroes},
		{"81808080808080808000", ErrPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadInt(r, new(int64))
			require.ErrorIs(t, err, test.want)
		})
	}
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

			r := &Reader{B: test.Bytes(t)}
			var got uint64
			require.NoError(ReadInt(r, &got))
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
		{"8080808080808080808080", ErrOverflow},
		{"ffffffffffffffffff02", ErrOverflow},
		{"8180808080808080808000", ErrOverflow},
		{"8100", ErrPaddedZeroes},
		{"818000", ErrPaddedZeroes},
		{"81808000", ErrPaddedZeroes},
		{"8180808000", ErrPaddedZeroes},
		{"818080808000", ErrPaddedZeroes},
		{"81808080808000", ErrPaddedZeroes},
		{"8180808080808000", ErrPaddedZeroes},
		{"818080808080808000", ErrPaddedZeroes},
		{"81808080808080808000", ErrPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadInt(r, new(uint64))
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendInt_int8(f *testing.F)   { f.Fuzz(testAppendInt[int8]) }
func FuzzAppendInt_uint8(f *testing.F)  { f.Fuzz(testAppendInt[uint8]) }
func FuzzAppendInt_int16(f *testing.F)  { f.Fuzz(testAppendInt[int16]) }
func FuzzAppendInt_uint16(f *testing.F) { f.Fuzz(testAppendInt[uint16]) }
func FuzzAppendInt_int32(f *testing.F)  { f.Fuzz(testAppendInt[int32]) }
func FuzzAppendInt_uint32(f *testing.F) { f.Fuzz(testAppendInt[uint32]) }
func FuzzAppendInt_int64(f *testing.F)  { f.Fuzz(testAppendInt[int64]) }
func FuzzAppendInt_uint64(f *testing.F) { f.Fuzz(testAppendInt[uint64]) }

func testAppendInt[T Int](t *testing.T, v T) {
	require := require.New(t)

	w := &Writer{}
	AppendInt(w, v)

	r := &Reader{B: w.B}
	var got T
	require.NoError(ReadInt(r, &got))
	require.Equal(v, got)
	require.Empty(r.B)
}

func FuzzSizeSint_int8(f *testing.F)  { f.Fuzz(testSizeSint[int8]) }
func FuzzSizeSint_int16(f *testing.F) { f.Fuzz(testSizeSint[int16]) }
func FuzzSizeSint_int32(f *testing.F) { f.Fuzz(testSizeSint[int32]) }
func FuzzSizeSint_int64(f *testing.F) { f.Fuzz(testSizeSint[int64]) }

func testSizeSint[T Sint](t *testing.T, v T) {
	w := &Writer{}
	AppendSint(w, v)

	size := SizeSint(v)
	require.Len(t, w.B, size)
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

			r := &Reader{B: test.Bytes(t)}
			var got int32
			require.NoError(ReadSint(r, &got))
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
		{"ffffffff10", ErrOverflow},
		{"8080808080808080808080", ErrOverflow},
		{"ffffffffffffffffff02", ErrOverflow},
		{"8180808080808080808000", ErrOverflow},
		{"8100", ErrPaddedZeroes},
		{"818000", ErrPaddedZeroes},
		{"81808000", ErrPaddedZeroes},
		{"8180808000", ErrPaddedZeroes},
		{"818080808000", ErrPaddedZeroes},
		{"81808080808000", ErrPaddedZeroes},
		{"8180808080808000", ErrPaddedZeroes},
		{"818080808080808000", ErrPaddedZeroes},
		{"81808080808080808000", ErrPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadSint(r, new(int32))
			require.ErrorIs(t, err, test.want)
		})
	}
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

			r := &Reader{B: test.Bytes(t)}
			var got int64
			require.NoError(ReadSint(r, &got))
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
		{"8080808080808080808080", ErrOverflow},
		{"ffffffffffffffffff02", ErrOverflow},
		{"8180808080808080808000", ErrOverflow},
		{"8100", ErrPaddedZeroes},
		{"818000", ErrPaddedZeroes},
		{"81808000", ErrPaddedZeroes},
		{"8180808000", ErrPaddedZeroes},
		{"818080808000", ErrPaddedZeroes},
		{"81808080808000", ErrPaddedZeroes},
		{"8180808080808000", ErrPaddedZeroes},
		{"818080808080808000", ErrPaddedZeroes},
		{"81808080808080808000", ErrPaddedZeroes},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadSint(r, new(int64))
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendSint_int8(f *testing.F)  { f.Fuzz(testAppendSint[int8]) }
func FuzzAppendSint_int16(f *testing.F) { f.Fuzz(testAppendSint[int16]) }
func FuzzAppendSint_int32(f *testing.F) { f.Fuzz(testAppendSint[int32]) }
func FuzzAppendSint_int64(f *testing.F) { f.Fuzz(testAppendSint[int64]) }

func testAppendSint[T Sint](t *testing.T, v T) {
	require := require.New(t)

	w := &Writer{}
	AppendSint(w, v)

	r := &Reader{B: w.B}
	var got T
	require.NoError(ReadSint(r, &got))
	require.Equal(v, got)
	require.Empty(r.B)
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

			r := &Reader{B: test.Bytes(t)}
			var got int32
			require.NoError(ReadFint32(r, &got))
			require.Equal(test.want, got)
			require.Empty(r.B)
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
			r := &Reader{B: test.Bytes(t)}
			err := ReadFint32(r, new(int32))
			require.ErrorIs(t, err, test.want)
		})
	}
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

			r := &Reader{B: test.Bytes(t)}
			var got uint32
			require.NoError(ReadFint32(r, &got))
			require.Equal(test.want, got)
			require.Empty(r.B)
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
			r := &Reader{B: test.Bytes(t)}
			err := ReadFint32(r, new(uint32))
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendFint32_int32(f *testing.F)  { f.Fuzz(testAppendFint32[int32]) }
func FuzzAppendFint32_uint32(f *testing.F) { f.Fuzz(testAppendFint32[uint32]) }

func testAppendFint32[T Int32](t *testing.T, v T) {
	require := require.New(t)

	w := &Writer{}
	AppendFint32(w, v)
	require.Len(w.B, SizeFint32)

	r := &Reader{B: w.B}
	var got T
	require.NoError(ReadFint32(r, &got))
	require.Equal(v, got)
	require.Empty(r.B)
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

			r := &Reader{B: test.Bytes(t)}
			var got int64
			require.NoError(ReadFint64(r, &got))
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
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadFint64(r, new(int64))
			require.ErrorIs(t, err, test.want)
		})
	}
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

			r := &Reader{B: test.Bytes(t)}
			var got uint64
			require.NoError(ReadFint64(r, &got))
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
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadFint64(r, new(uint64))
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendFint64_int64(f *testing.F)  { f.Fuzz(testAppendFint64[int64]) }
func FuzzAppendFint64_uint64(f *testing.F) { f.Fuzz(testAppendFint64[uint64]) }

func testAppendFint64[T Int64](t *testing.T, v T) {
	require := require.New(t)

	w := &Writer{}
	AppendFint64(w, v)
	require.Len(w.B, SizeFint64)

	r := &Reader{B: w.B}
	var got T
	require.NoError(ReadFint64(r, &got))
	require.Equal(v, got)
	require.Empty(r.B)
}

func TestReadBool(t *testing.T) {
	validTests := []validTest[bool]{
		{"00", false},
		{"01", true},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{B: test.Bytes(t)}
			var got bool
			require.NoError(ReadBool(r, &got))
			require.Equal(test.want, got)
			require.Empty(r.B)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"02", ErrInvalidBool},
		{"ff", ErrInvalidBool},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadBool(r, new(bool))
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestAppendBool(t *testing.T) {
	for _, b := range []bool{false, true} {
		t.Run(strconv.FormatBool(b), func(t *testing.T) {
			require := require.New(t)

			w := &Writer{}
			AppendBool(w, b)
			require.Len(w.B, SizeBool)

			r := &Reader{B: w.B}
			var got bool
			require.NoError(ReadBool(r, &got))
			require.Equal(b, got)
			require.Empty(r.B)
		})
	}
}

func FuzzSizeBytes_string(f *testing.F) { f.Fuzz(testSizeBytes[string]) }
func FuzzSizeBytes_bytes(f *testing.F)  { f.Fuzz(testSizeBytes[[]byte]) }

func testSizeBytes[T Bytes](t *testing.T, v T) {
	w := &Writer{}
	AppendBytes(w, v)

	size := SizeBytes(v)
	require.Len(t, w.B, size)
}

func FuzzCountBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		require := require.New(t)

		var (
			tag   string
			bytes [][]byte
		)
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&bytes)
		if len(tag) == 0 {
			return
		}

		w := &Writer{}
		for _, v := range bytes {
			Append(w, tag)
			AppendBytes(w, v)
		}

		count, err := CountBytes(w.B, tag)
		require.NoError(err)
		require.Len(bytes, count)
	})
}

func TestReadString(t *testing.T) {
	validTests := []validTest[string]{
		{"00", ""},
		{"0774657374696e67", "testing"},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{B: test.Bytes(t)}
			var got string
			require.NoError(ReadString(r, &got))
			require.Equal(test.want, got)
			require.Empty(r.B)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"870074657374696e67", ErrPaddedZeroes},
		{"ffffffffffffffffff01", ErrInvalidLength},
		{"01", io.ErrUnexpectedEOF},
		{"01C2", ErrStringNotUTF8},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadString(r, new(string))
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReadBytes(t *testing.T) {
	validTests := []validTest[[]byte]{
		{"00", []byte{}},
		{"0774657374696e67", []byte("testing")},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{B: test.Bytes(t)}
			var got []byte
			require.NoError(ReadBytes(r, &got))
			require.Equal(test.want, got)
			require.Empty(r.B)
		})
	}

	invalidTests := []invalidTest{
		{"", io.ErrUnexpectedEOF},
		{"870074657374696e67", ErrPaddedZeroes},
		{"ffffffffffffffffff01", ErrInvalidLength},
		{"01", io.ErrUnexpectedEOF},
	}
	for _, test := range invalidTests {
		t.Run(test.hex, func(t *testing.T) {
			r := &Reader{B: test.Bytes(t)}
			err := ReadBytes(r, new([]byte))
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendBytes_string(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		if !utf8.ValidString(v) {
			return
		}

		require := require.New(t)

		w := &Writer{}
		AppendBytes(w, v)

		r := &Reader{B: w.B}
		var got string
		require.NoError(ReadString(r, &got))
		require.Equal(v, got)
		require.Empty(r.B)
	})
}

func FuzzAppendBytes_bytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, v []byte) {
		require := require.New(t)

		w := &Writer{}
		AppendBytes(w, v)

		r := &Reader{B: w.B}
		var got []byte
		require.NoError(ReadBytes(r, &got))
		require.Equal(v, got)
		require.Empty(r.B)
	})
}
