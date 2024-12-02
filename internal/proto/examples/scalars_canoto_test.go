package examples

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/StephenButtolph/canoto"
	"github.com/StephenButtolph/canoto/internal/proto/pb"
)

func FuzzScalars_UnmarshalCanoto(f *testing.F) {
	f.Fuzz(func(
		t *testing.T,
		i32 int32,
		i64 int64,
		u32 uint32,
		u64 uint64,
		s32 int32,
		s64 int64,
		f32 uint32,
		f64 uint64,
		sf32 int32,
		sf64 int64,
		b bool,
		s string,
		bs []byte,
	) {
		if len(bs) == 0 {
			bs = nil
		}

		require := require.New(t)

		var largestFieldNumber *pb.LargestFieldNumber
		if i32 != 0 {
			largestFieldNumber = &pb.LargestFieldNumber{
				Int32: i32,
			}
		}

		pbScalars := pb.Scalars{
			Int32:              i32,
			Int64:              i64,
			Uint32:             u32,
			Uint64:             u64,
			Sint32:             s32,
			Sint64:             s64,
			Fixed32:            f32,
			Fixed64:            f64,
			Sfixed32:           sf32,
			Sfixed64:           sf64,
			Bool:               b,
			String_:            s,
			Bytes:              bs,
			LargestFieldNumber: largestFieldNumber,
		}
		pbScalarsBytes, err := proto.Marshal(&pbScalars)
		if err != nil {
			return
		}

		var (
			canotoScalars Scalars
			reader        = canoto.Reader{B: pbScalarsBytes}
		)
		require.NoError(canotoScalars.UnmarshalCanoto(&reader))

		require.Equal(
			Scalars{
				Int32:    i32,
				Int64:    i64,
				Uint32:   u32,
				Uint64:   u64,
				Sint32:   s32,
				Sint64:   s64,
				Fixed32:  f32,
				Fixed64:  f64,
				Sfixed32: sf32,
				Sfixed64: sf64,
				Bool:     b,
				String:   s,
				Bytes:    bs,
				LargestFieldNumber: LargestFieldNumber{
					Int32: i32,
				},
			},
			canotoScalars,
		)
	})
}

func FuzzScalars_MarshalCanoto(f *testing.F) {
	f.Fuzz(func(
		t *testing.T,
		i32 int32,
		i64 int64,
		u32 uint32,
		u64 uint64,
		s32 int32,
		s64 int64,
		f32 uint32,
		f64 uint64,
		sf32 int32,
		sf64 int64,
		b bool,
		s string,
		bs []byte,
	) {
		if len(bs) == 0 {
			bs = nil
		}

		require := require.New(t)

		cbScalars := Scalars{
			Int32:    i32,
			Int64:    i64,
			Uint32:   u32,
			Uint64:   u64,
			Sint32:   s32,
			Sint64:   s64,
			Fixed32:  f32,
			Fixed64:  f64,
			Sfixed32: sf32,
			Sfixed64: sf64,
			Bool:     b,
			String:   s,
			Bytes:    bs,
			LargestFieldNumber: LargestFieldNumber{
				Int32: i32,
			},
		}
		if !cbScalars.ValidCanoto() {
			return
		}

		size := cbScalars.SizeCanoto()
		w := canoto.Writer{
			B: make([]byte, 0, size),
		}
		cbScalars.MarshalCanoto(&w)
		require.Len(w.B, size)

		var pbScalars pb.Scalars
		require.NoError(proto.Unmarshal(w.B, &pbScalars))

		var expectedLargestFieldNumber *pb.LargestFieldNumber
		if i32 != 0 {
			expectedLargestFieldNumber = &pb.LargestFieldNumber{
				Int32: i32,
			}
		}

		var actualLargestFieldNumber *pb.LargestFieldNumber
		if pbScalars.LargestFieldNumber != nil {
			actualLargestFieldNumber = &pb.LargestFieldNumber{
				Int32: pbScalars.LargestFieldNumber.Int32,
			}
		}

		require.Equal(
			pb.Scalars{
				Int32:              i32,
				Int64:              i64,
				Uint32:             u32,
				Uint64:             u64,
				Sint32:             s32,
				Sint64:             s64,
				Fixed32:            f32,
				Fixed64:            f64,
				Sfixed32:           sf32,
				Sfixed64:           sf64,
				Bool:               b,
				String_:            s,
				Bytes:              bs,
				LargestFieldNumber: expectedLargestFieldNumber,
			},
			pb.Scalars{
				Int32:              pbScalars.Int32,
				Int64:              pbScalars.Int64,
				Uint32:             pbScalars.Uint32,
				Uint64:             pbScalars.Uint64,
				Sint32:             pbScalars.Sint32,
				Sint64:             pbScalars.Sint64,
				Fixed32:            pbScalars.Fixed32,
				Fixed64:            pbScalars.Fixed64,
				Sfixed32:           pbScalars.Sfixed32,
				Sfixed64:           pbScalars.Sfixed64,
				Bool:               pbScalars.Bool,
				String_:            pbScalars.String_,
				Bytes:              pbScalars.Bytes,
				LargestFieldNumber: actualLargestFieldNumber,
			},
		)
	})
}

func FuzzScalars_Canonical(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		require := require.New(t)

		var (
			scalars Scalars
			reader  = canoto.Reader{B: b}
		)
		err := scalars.UnmarshalCanoto(&reader)
		if err != nil {
			return
		}

		size := scalars.SizeCanoto()
		require.Len(b, size)

		w := canoto.Writer{
			B: make([]byte, 0, size),
		}
		scalars.MarshalCanoto(&w)
		require.Equal(b, w.B)
	})
}

func BenchmarkScalars_MarshalCanoto(b *testing.B) {
	for range b.N {
		cbScalars := Scalars{
			Int32:    216457,
			Int64:    -2138746,
			Uint32:   32485976,
			Uint64:   287634,
			Sint32:   -12786345,
			Sint64:   98761243,
			Fixed32:  98765234,
			Fixed64:  1234576,
			Sfixed32: -21348976,
			Sfixed64: 98756432,
			Bool:     true,
			String:   "asdkcjhbsakjdhcbasdc",
			Bytes:    []byte("kjsadhbcaskjdvhb"),
			LargestFieldNumber: LargestFieldNumber{
				Int32: 216457,
			},
		}

		w := canoto.Writer{
			B: make([]byte, 0, cbScalars.SizeCanoto()),
		}
		cbScalars.MarshalCanoto(&w)
	}
}

func BenchmarkScalars_UnmarshalCanoto(b *testing.B) {
	cbScalars := Scalars{
		Int32:    216457,
		Int64:    -2138746,
		Uint32:   32485976,
		Uint64:   287634,
		Sint32:   -12786345,
		Sint64:   98761243,
		Fixed32:  98765234,
		Fixed64:  1234576,
		Sfixed32: -21348976,
		Sfixed64: 98756432,
		Bool:     true,
		String:   "asdkcjhbsakjdhcbasdc",
		Bytes:    []byte("kjsadhbcaskjdvhb"),
		LargestFieldNumber: LargestFieldNumber{
			Int32: 216457,
		},
	}
	w := canoto.Writer{
		B: make([]byte, 0, cbScalars.SizeCanoto()),
	}
	cbScalars.MarshalCanoto(&w)

	for _, unsafe := range []bool{false, true} {
		b.Run("unsafe="+strconv.FormatBool(unsafe), func(b *testing.B) {
			for range b.N {
				var (
					scalars Scalars
					reader  = canoto.Reader{
						B:      w.B,
						Unsafe: unsafe,
					}
				)
				_ = scalars.UnmarshalCanoto(&reader)
			}
		})
	}
}

func BenchmarkScalars_MarshalProto(b *testing.B) {
	for range b.N {
		pbScalars := pb.Scalars{
			Int32:    216457,
			Int64:    -2138746,
			Uint32:   32485976,
			Uint64:   287634,
			Sint32:   -12786345,
			Sint64:   98761243,
			Fixed32:  98765234,
			Fixed64:  1234576,
			Sfixed32: -21348976,
			Sfixed64: 98756432,
			Bool:     true,
			String_:  "asdkcjhbsakjdhcbasdc",
			Bytes:    []byte("kjsadhbcaskjdvhb"),
			LargestFieldNumber: &pb.LargestFieldNumber{
				Int32: 216457,
			},
		}
		_, _ = proto.Marshal(&pbScalars)
	}
}

func BenchmarkScalars_UnmarshalProto(b *testing.B) {
	pbScalars := pb.Scalars{
		Int32:    216457,
		Int64:    -2138746,
		Uint32:   32485976,
		Uint64:   287634,
		Sint32:   -12786345,
		Sint64:   98761243,
		Fixed32:  98765234,
		Fixed64:  1234576,
		Sfixed32: -21348976,
		Sfixed64: 98756432,
		Bool:     true,
		String_:  "asdkcjhbsakjdhcbasdc",
		Bytes:    []byte("kjsadhbcaskjdvhb"),
		LargestFieldNumber: &pb.LargestFieldNumber{
			Int32: 216457,
		},
	}
	scalarsBytes, err := proto.Marshal(&pbScalars)
	require.NoError(b, err)

	b.ResetTimer()
	for range b.N {
		var (
			scalars pb.Scalars
			reader  = proto.UnmarshalOptions{
				Merge: true,
			}
		)
		_ = reader.Unmarshal(scalarsBytes, &scalars)
	}
}
