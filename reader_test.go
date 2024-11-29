package canoto

import (
	"encoding/hex"
	"io"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
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

func TestReader_Tag(t *testing.T) {
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
			gotField, gotType, err := r.Tag()
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
			_, _, err := r.Tag()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Int32(t *testing.T) {
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
			got, err := r.Int32()
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
			_, err := r.Int32()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Int64(t *testing.T) {
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
			got, err := r.Int64()
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
			_, err := r.Int64()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Uint32(t *testing.T) {
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
			got, err := r.Uint32()
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
			_, err := r.Uint32()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Uint64(t *testing.T) {
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
			got, err := r.Uint64()
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
			_, err := r.Uint64()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Sint32(t *testing.T) {
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
			got, err := r.Sint32()
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
			_, err := r.Sint32()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Sint64(t *testing.T) {
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
			got, err := r.Sint64()
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
			_, err := r.Sint64()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Fixed32(t *testing.T) {
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
			got, err := r.Fixed32()
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
			_, err := r.Fixed32()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Fixed64(t *testing.T) {
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
			got, err := r.Fixed64()
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
			_, err := r.Fixed64()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Sfixed32(t *testing.T) {
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
			got, err := r.Sfixed32()
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
			_, err := r.Sfixed32()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Sfixed64(t *testing.T) {
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
			got, err := r.Sfixed64()
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
			_, err := r.Sfixed64()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Bool(t *testing.T) {
	validTests := []validTest[bool]{
		{"00", false},
		{"01", true},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := r.Bool()
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
			_, err := r.Bool()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_String(t *testing.T) {
	validTests := []validTest[string]{
		{"00", ""},
		{"0774657374696e67", "testing"},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := r.String()
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
			_, err := r.String()
			require.ErrorIs(t, err, test.want)
		})
	}
}

func TestReader_Bytes(t *testing.T) {
	validTests := []validTest[[]byte]{
		{"00", []byte{}},
		{"0774657374696e67", []byte("testing")},
	}
	for _, test := range validTests {
		t.Run(test.hex, func(t *testing.T) {
			require := require.New(t)

			r := &Reader{b: test.Bytes(t)}
			got, err := r.Bytes()
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
			_, err := r.Bytes()
			require.ErrorIs(t, err, test.want)
		})
	}
}
