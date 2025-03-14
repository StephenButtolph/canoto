//go:generate canoto --internal $GOFILE

package canoto

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"math"
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepudds/fzgen/fuzzer"
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

func FuzzSizeUint_uint8(f *testing.F)  { f.Fuzz(testSizeUint[uint8]) }
func FuzzSizeUint_uint16(f *testing.F) { f.Fuzz(testSizeUint[uint16]) }
func FuzzSizeUint_uint32(f *testing.F) { f.Fuzz(testSizeUint[uint32]) }
func FuzzSizeUint_uint64(f *testing.F) { f.Fuzz(testSizeUint[uint64]) }

func testSizeUint[T Uint](t *testing.T, v T) {
	w := &Writer{}
	AppendUint(w, v)

	size := SizeUint(v)
	require.Len(t, w.B, int(size)) //#nosec G115 // False positive
}

func FuzzCountInts_uint8(f *testing.F)  { f.Fuzz(testCountInts[uint8]) }
func FuzzCountInts_uint16(f *testing.F) { f.Fuzz(testCountInts[uint16]) }
func FuzzCountInts_uint32(f *testing.F) { f.Fuzz(testCountInts[uint32]) }
func FuzzCountInts_uint64(f *testing.F) { f.Fuzz(testCountInts[uint64]) }

func testCountInts[T Uint](t *testing.T, data []byte) {
	require := require.New(t)

	var nums []T
	fz := fuzzer.NewFuzzer(data)
	fz.Fill(&nums)

	w := &Writer{}
	for _, num := range nums {
		AppendUint(w, num)
	}

	count := CountInts(w.B)
	require.Len(nums, int(count)) //#nosec G115 // False positive
}

func TestReadUint_uint32(t *testing.T) {
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
			require.NoError(ReadUint(r, &got))
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
			err := ReadUint(r, new(uint32))
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReadUint_uint64(t *testing.T) {
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
			require.NoError(ReadUint(r, &got))
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
			err := ReadUint(r, new(uint64))
			require.ErrorIs(t, err, test.want)
		})
	}
}

func FuzzAppendInt_uint8(f *testing.F)  { f.Fuzz(testAppendUint[uint8]) }
func FuzzAppendInt_uint16(f *testing.F) { f.Fuzz(testAppendUint[uint16]) }
func FuzzAppendInt_uint32(f *testing.F) { f.Fuzz(testAppendUint[uint32]) }
func FuzzAppendInt_uint64(f *testing.F) { f.Fuzz(testAppendUint[uint64]) }

func testAppendUint[T Uint](t *testing.T, v T) {
	require := require.New(t)

	w := &Writer{}
	AppendUint(w, v)

	r := &Reader{B: w.B}
	var got T
	require.NoError(ReadUint(r, &got))
	require.Equal(v, got)
	require.Empty(r.B)
}

func FuzzSizeInt_int8(f *testing.F)  { f.Fuzz(testSizeInt[int8]) }
func FuzzSizeInt_int16(f *testing.F) { f.Fuzz(testSizeInt[int16]) }
func FuzzSizeInt_int32(f *testing.F) { f.Fuzz(testSizeInt[int32]) }
func FuzzSizeInt_int64(f *testing.F) { f.Fuzz(testSizeInt[int64]) }

func testSizeInt[T Int](t *testing.T, v T) {
	w := &Writer{}
	AppendInt(w, v)

	size := SizeInt(v)
	require.Len(t, w.B, int(size)) //#nosec G115 // False positive
}

func TestReadInt_int32(t *testing.T) {
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
			err := ReadInt(r, new(int32))
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReadInt_int64(t *testing.T) {
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

func FuzzAppendInt_int8(f *testing.F)  { f.Fuzz(testAppendInt[int8]) }
func FuzzAppendInt_int16(f *testing.F) { f.Fuzz(testAppendInt[int16]) }
func FuzzAppendInt_int32(f *testing.F) { f.Fuzz(testAppendInt[int32]) }
func FuzzAppendInt_int64(f *testing.F) { f.Fuzz(testAppendInt[int64]) }

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
	require.Len(t, w.B, int(size)) //#nosec G115 // False positive
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
		require.Len(bytes, int(count)) //#nosec G115 // False positive
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
		{"ffffffffffffffffff01", io.ErrUnexpectedEOF},
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
		{"ffffffffffffffffff01", io.ErrUnexpectedEOF},
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
		if !ValidString(v) {
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

func TestIsSigned(t *testing.T) {
	require := require.New(t)

	require.True(isSigned[int8]())
	require.True(isSigned[int16]())
	require.True(isSigned[int32]())
	require.True(isSigned[int64]())

	require.False(isSigned[uint8]())
	require.False(isSigned[uint16]())
	require.False(isSigned[uint32]())
	require.False(isSigned[uint64]())
}

func TestSizeOf(t *testing.T) {
	require := require.New(t)

	require.Equal(SizeEnum8, SizeOf[int8](0))
	require.Equal(SizeEnum8, SizeOf[uint8](0))

	require.Equal(SizeEnum16, SizeOf[int16](0))
	require.Equal(SizeEnum16, SizeOf[uint16](0))

	require.Equal(SizeEnum32, SizeOf[int32](0))
	require.Equal(SizeEnum32, SizeOf[uint32](0))

	require.Equal(SizeEnum64, SizeOf[int64](0))
	require.Equal(SizeEnum64, SizeOf[uint64](0))
}

func TestIsBytesEmpty(t *testing.T) {
	require := require.New(t)

	require.True(isBytesEmpty(make([]byte, 0)))
	require.True(isBytesEmpty(make([]byte, 10)))

	require.False(isBytesEmpty([]byte{0: 1}))
	require.False(isBytesEmpty([]byte{10: 1}))
}

type SpecFuzzer struct {
	Int8                       int8                         `canoto:"int,1"              json:"Int8,omitempty"`
	Int16                      int16                        `canoto:"int,2"              json:"Int16,omitempty"`
	Int32                      int32                        `canoto:"int,3"              json:"Int32,omitempty"`
	Int64                      int64                        `canoto:"int,4"              json:"Int64,omitempty"`
	Uint8                      uint8                        `canoto:"uint,5"             json:"Uint8,omitempty"`
	Uint16                     uint16                       `canoto:"uint,6"             json:"Uint16,omitempty"`
	Uint32                     uint32                       `canoto:"uint,7"             json:"Uint32,omitempty"`
	Uint64                     uint64                       `canoto:"uint,8"             json:"Uint64,omitempty"`
	Sfixed32                   int32                        `canoto:"fint32,9"           json:"Sfixed32,omitempty"`
	Fixed32                    uint32                       `canoto:"fint32,10"          json:"Fixed32,omitempty"`
	Sfixed64                   int64                        `canoto:"fint64,11"          json:"Sfixed64,omitempty"`
	Fixed64                    uint64                       `canoto:"fint64,12"          json:"Fixed64,omitempty"`
	Bool                       bool                         `canoto:"bool,13"            json:"Bool,omitempty"`
	String                     string                       `canoto:"string,14"          json:"String,omitempty"`
	Bytes                      []byte                       `canoto:"bytes,15"           json:"Bytes,omitempty"`
	LargestFieldNumber         *LargestFieldNumber[uint32]  `canoto:"pointer,16"         json:"LargestFieldNumber,omitempty"`
	RepeatedInt8               []int8                       `canoto:"repeated int,17"    json:"RepeatedInt8,omitempty"`
	RepeatedInt16              []int16                      `canoto:"repeated int,18"    json:"RepeatedInt16,omitempty"`
	RepeatedInt32              []int32                      `canoto:"repeated int,19"    json:"RepeatedInt32,omitempty"`
	RepeatedInt64              []int64                      `canoto:"repeated int,20"    json:"RepeatedInt64,omitempty"`
	RepeatedUint16             []uint16                     `canoto:"repeated uint,21"   json:"RepeatedUint16,omitempty"`
	RepeatedUint32             []uint32                     `canoto:"repeated uint,22"   json:"RepeatedUint32,omitempty"`
	RepeatedUint64             []uint64                     `canoto:"repeated uint,23"   json:"RepeatedUint64,omitempty"`
	RepeatedSfixed32           []int32                      `canoto:"repeated fint32,24" json:"RepeatedSfixed32,omitempty"`
	RepeatedFixed32            []uint32                     `canoto:"repeated fint32,25" json:"RepeatedFixed32,omitempty"`
	RepeatedSfixed64           []int64                      `canoto:"repeated fint64,26" json:"RepeatedSfixed64,omitempty"`
	RepeatedFixed64            []uint64                     `canoto:"repeated fint64,27" json:"RepeatedFixed64,omitempty"`
	RepeatedBool               []bool                       `canoto:"repeated bool,28"   json:"RepeatedBool,omitempty"`
	RepeatedString             []string                     `canoto:"repeated string,29" json:"RepeatedString,omitempty"`
	RepeatedBytes              [][]byte                     `canoto:"repeated bytes,30"  json:"RepeatedBytes,omitempty"`
	RepeatedLargestFieldNumber []LargestFieldNumber[uint32] `canoto:"repeated value,31"  json:"RepeatedLargestFieldNumber,omitempty"`
	OneOf                      *OneOf                       `canoto:"pointer,32"         json:"OneOf,omitempty"`
	Pointer                    *LargestFieldNumber[uint32]  `canoto:"pointer,33"         json:"Pointer,omitempty"`
	Field                      *LargestFieldNumber[uint32]  `canoto:"field,34"           json:"Field,omitempty"`
	Recursive                  *SpecFuzzer                  `canoto:"pointer,35"         json:"Recursive,omitempty"`

	canotoData canotoData_SpecFuzzer
}

type LargestFieldNumber[T Uint] struct {
	Uint T `canoto:"uint,536870911" json:"Uint,omitempty"`

	canotoData canotoData_LargestFieldNumber
}

type OneOf struct {
	A1 int32 `canoto:"int,1,A" json:"A1,omitempty"`
	B1 int32 `canoto:"int,3,B" json:"B1,omitempty"`
	B2 int64 `canoto:"int,4,B" json:"B2,omitempty"`
	C  int32 `canoto:"int,5"   json:"C,omitempty"`
	D  int64 `canoto:"int,6"   json:"D,omitempty"`
	A2 int64 `canoto:"int,7,A" json:"A2,omitempty"`

	canotoData canotoData_OneOf
}

func FuzzSpec(f *testing.F) {
	full := SpecFuzzer{
		Int8:     -31,
		Int16:    -2164,
		Int32:    -12786345,
		Int64:    98761243,
		Uint8:    254,
		Uint16:   21645,
		Uint32:   32485976,
		Uint64:   287634,
		Fixed32:  98765234,
		Fixed64:  1234576,
		Sfixed32: -21348976,
		Sfixed64: 98756432,
		Bool:     true,
		String:   "hi my name is Bob",
		Bytes:    []byte("hi my name is Bob too"),
		LargestFieldNumber: &LargestFieldNumber[uint32]{
			Uint: 216457,
		},

		RepeatedInt8:     []int8{1, 2, 3},
		RepeatedInt16:    []int16{1, 2, 3},
		RepeatedInt32:    []int32{1, 2, 3},
		RepeatedInt64:    []int64{1, 2, 3},
		RepeatedUint16:   []uint16{1, 2, 3},
		RepeatedUint32:   []uint32{1, 2, 3},
		RepeatedUint64:   []uint64{1, 2, 3},
		RepeatedFixed32:  []uint32{1, 2, 3},
		RepeatedFixed64:  []uint64{1, 2, 3},
		RepeatedSfixed32: []int32{1, 2, 3},
		RepeatedSfixed64: []int64{1, 2, 3},
		RepeatedBool:     []bool{true, false, true},
		RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
		RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
		RepeatedLargestFieldNumber: []LargestFieldNumber[uint32]{
			{Uint: 123455},
			{Uint: 876523},
		},

		OneOf: &OneOf{
			A1: 1,
			B2: 2,
			C:  3,
			D:  4,
		},
	}
	fullBytes := full.MarshalCanoto()
	f.Add(fullBytes)

	full.Recursive = new(SpecFuzzer)
	require.NoError(f, full.Recursive.UnmarshalCanoto(fullBytes))

	recursiveFullBytes := full.MarshalCanoto()
	f.Add(recursiveFullBytes)

	spec := (*SpecFuzzer)(nil).CanotoSpec()
	f.Fuzz(func(t *testing.T, b []byte) {
		// Standardize the empty input for reflect based equality.
		if len(b) == 0 {
			b = nil
		}
		originalBytes := slices.Clone(b)

		require := require.New(t)

		// Verify that unmarshalling the message using [Unmarshal] returns the
		// same error as the unmarshalling the message directly.
		var msg SpecFuzzer
		expectedErr := msg.UnmarshalCanoto(b)
		anyMSG, actualErr := Unmarshal(spec, b)
		require.Equal(expectedErr, actualErr)

		if expectedErr != nil {
			return
		}

		// Modify the original bytes to ensure that [Unmarshal] does not hold a
		// reference to the originally passed in slice.
		for i := range b {
			b[i]++
		}

		// Verify that the unmarshalled messages have the same json
		// representation.
		expectedJSON, err := json.Marshal(&msg)
		require.NoError(err)
		actualJSON, err := json.Marshal(anyMSG)
		require.NoError(err)
		require.JSONEq(string(expectedJSON), string(actualJSON))

		// Verify that re-marshalling the unmarshalled message returns the same
		// bytes as the original message.
		actualBytes, err := Marshal(spec, anyMSG)
		require.NoError(err)
		require.Equal(originalBytes, actualBytes)
	})
}
