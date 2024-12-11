package examples

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepudds/fzgen/fuzzer"
	"google.golang.org/protobuf/proto"

	"github.com/StephenButtolph/canoto"
	"github.com/StephenButtolph/canoto/canoto/internal/generate/proto/pb"
)

func canonicalizeSlice[T any](s []T) []T {
	if len(s) == 0 {
		return nil
	}
	return s
}

func canonicalizeCanotoScalars(s Scalars) Scalars {
	s.Bytes = canonicalizeSlice(s.Bytes)
	s.RepeatedInt32 = canonicalizeSlice(s.RepeatedInt32)
	s.RepeatedInt64 = canonicalizeSlice(s.RepeatedInt64)
	s.RepeatedUint32 = canonicalizeSlice(s.RepeatedUint32)
	s.RepeatedUint64 = canonicalizeSlice(s.RepeatedUint64)
	s.RepeatedSint32 = canonicalizeSlice(s.RepeatedSint32)
	s.RepeatedSint64 = canonicalizeSlice(s.RepeatedSint64)
	s.RepeatedFixed32 = canonicalizeSlice(s.RepeatedFixed32)
	s.RepeatedFixed64 = canonicalizeSlice(s.RepeatedFixed64)
	s.RepeatedSfixed32 = canonicalizeSlice(s.RepeatedSfixed32)
	s.RepeatedSfixed64 = canonicalizeSlice(s.RepeatedSfixed64)
	s.RepeatedBool = canonicalizeSlice(s.RepeatedBool)
	s.RepeatedString = canonicalizeSlice(s.RepeatedString)
	s.RepeatedBytes = canonicalizeSlice(s.RepeatedBytes)
	s.canotoData = canotoData_Scalars{}
	return s
}

func canonicalizeProtoScalars(s *pb.Scalars) pb.Scalars {
	var largestFieldNumber *pb.LargestFieldNumber
	if s.LargestFieldNumber != nil {
		largestFieldNumber = &pb.LargestFieldNumber{
			Int32: s.LargestFieldNumber.Int32,
		}
	}
	return pb.Scalars{
		Int32:              s.Int32,
		Int64:              s.Int64,
		Uint32:             s.Uint32,
		Uint64:             s.Uint64,
		Sint32:             s.Sint32,
		Sint64:             s.Sint64,
		Fixed32:            s.Fixed32,
		Fixed64:            s.Fixed64,
		Sfixed32:           s.Sfixed32,
		Sfixed64:           s.Sfixed64,
		Bool:               s.Bool,
		String_:            s.String_,
		Bytes:              s.Bytes,
		LargestFieldNumber: largestFieldNumber,
		RepeatedInt32:      s.RepeatedInt32,
		RepeatedInt64:      s.RepeatedInt64,
		RepeatedUint32:     s.RepeatedUint32,
		RepeatedUint64:     s.RepeatedUint64,
		RepeatedSint32:     s.RepeatedSint32,
		RepeatedSint64:     s.RepeatedSint64,
		RepeatedFixed32:    s.RepeatedFixed32,
		RepeatedFixed64:    s.RepeatedFixed64,
		RepeatedSfixed32:   s.RepeatedSfixed32,
		RepeatedSfixed64:   s.RepeatedSfixed64,
		RepeatedBool:       s.RepeatedBool,
		RepeatedString:     s.RepeatedString,
		RepeatedBytes:      s.RepeatedBytes,
	}
}

func canotoScalarsToProto(s Scalars) pb.Scalars {
	var largestFieldNumber *pb.LargestFieldNumber
	if s.LargestFieldNumber.Int32 != 0 {
		largestFieldNumber = &pb.LargestFieldNumber{
			Int32: s.LargestFieldNumber.Int32,
		}
	}
	return pb.Scalars{
		Int32:              s.Int32,
		Int64:              s.Int64,
		Uint32:             s.Uint32,
		Uint64:             s.Uint64,
		Sint32:             s.Sint32,
		Sint64:             s.Sint64,
		Fixed32:            s.Fixed32,
		Fixed64:            s.Fixed64,
		Sfixed32:           s.Sfixed32,
		Sfixed64:           s.Sfixed64,
		Bool:               s.Bool,
		String_:            s.String,
		Bytes:              s.Bytes,
		LargestFieldNumber: largestFieldNumber,
		RepeatedInt32:      s.RepeatedInt32,
		RepeatedInt64:      s.RepeatedInt64,
		RepeatedUint32:     s.RepeatedUint32,
		RepeatedUint64:     s.RepeatedUint64,
		RepeatedSint32:     s.RepeatedSint32,
		RepeatedSint64:     s.RepeatedSint64,
		RepeatedFixed32:    s.RepeatedFixed32,
		RepeatedFixed64:    s.RepeatedFixed64,
		RepeatedSfixed32:   s.RepeatedSfixed32,
		RepeatedSfixed64:   s.RepeatedSfixed64,
		RepeatedBool:       s.RepeatedBool,
		RepeatedString:     s.RepeatedString,
		RepeatedBytes:      s.RepeatedBytes,
	}
}

func FuzzScalars_UnmarshalCanoto(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		require := require.New(t)

		var canotoScalars Scalars
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&canotoScalars)
		canotoScalars = canonicalizeCanotoScalars(canotoScalars)

		pbScalars := canotoScalarsToProto(canotoScalars)
		pbScalarsBytes, err := proto.Marshal(&pbScalars)
		if err != nil {
			return
		}

		var canotoScalarsFromProto Scalars
		require.NoError(canotoScalarsFromProto.UnmarshalCanoto(pbScalarsBytes))
		require.Equal(canotoScalars, canotoScalarsFromProto)
	})
}

func FuzzScalars_MarshalCanoto(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		require := require.New(t)

		var canotoScalars Scalars
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&canotoScalars)
		canotoScalars = canonicalizeCanotoScalars(canotoScalars)
		if !canotoScalars.ValidCanoto() {
			return
		}

		size := canotoScalars.CalculateCanotoSize()
		w := canoto.Writer{
			B: make([]byte, 0, size),
		}
		canotoScalars.MarshalCanotoInto(&w)
		require.Len(w.B, size)

		var pbScalars pb.Scalars
		require.NoError(proto.Unmarshal(w.B, &pbScalars))
		require.Equal(
			canotoScalarsToProto(canotoScalars),
			canonicalizeProtoScalars(&pbScalars),
		)
	})
}

func FuzzScalars_Canonical(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		require := require.New(t)

		var scalars Scalars
		err := scalars.UnmarshalCanoto(b)
		if err != nil {
			return
		}

		size := scalars.CalculateCanotoSize()
		require.Len(b, size)

		w := canoto.Writer{
			B: make([]byte, 0, size),
		}
		scalars.MarshalCanotoInto(&w)
		require.Equal(b, w.B)
	})
}

func BenchmarkScalars_MarshalCanoto(b *testing.B) {
	b.Run("full stack", func(b *testing.B) {
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
				String:   "hi my name is Bob",
				Bytes:    []byte("hi my name is Bob too"),
				LargestFieldNumber: LargestFieldNumber{
					Int32: 216457,
				},
				RepeatedInt32:    []int32{1, 2, 3},
				RepeatedInt64:    []int64{1, 2, 3},
				RepeatedUint32:   []uint32{1, 2, 3},
				RepeatedUint64:   []uint64{1, 2, 3},
				RepeatedSint32:   []int32{1, 2, 3},
				RepeatedSint64:   []int64{1, 2, 3},
				RepeatedFixed32:  []uint32{1, 2, 3},
				RepeatedFixed64:  []uint64{1, 2, 3},
				RepeatedSfixed32: []int32{1, 2, 3},
				RepeatedSfixed64: []int64{1, 2, 3},
				RepeatedBool:     []bool{true, false, true},
				RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
				RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
			}
			cbScalars.MarshalCanoto()
		}
	})
	b.Run("full heap", func(b *testing.B) {
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
			String:   "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
			LargestFieldNumber: LargestFieldNumber{
				Int32: 216457,
			},
			RepeatedInt32:    []int32{1, 2, 3},
			RepeatedInt64:    []int64{1, 2, 3},
			RepeatedUint32:   []uint32{1, 2, 3},
			RepeatedUint64:   []uint64{1, 2, 3},
			RepeatedSint32:   []int32{1, 2, 3},
			RepeatedSint64:   []int64{1, 2, 3},
			RepeatedFixed32:  []uint32{1, 2, 3},
			RepeatedFixed64:  []uint64{1, 2, 3},
			RepeatedSfixed32: []int32{1, 2, 3},
			RepeatedSfixed64: []int64{1, 2, 3},
			RepeatedBool:     []bool{true, false, true},
			RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
			RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
		}
		for range b.N {
			cbScalars.MarshalCanoto()
		}
	})
	b.Run("primitives stack", func(b *testing.B) {
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
				String:   "hi my name is Bob",
				Bytes:    []byte("hi my name is Bob too"),
				LargestFieldNumber: LargestFieldNumber{
					Int32: 216457,
				},
			}
			cbScalars.MarshalCanoto()
		}
	})
	b.Run("primitives heap", func(b *testing.B) {
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
			String:   "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
			LargestFieldNumber: LargestFieldNumber{
				Int32: 216457,
			},
		}
		for range b.N {
			cbScalars.MarshalCanoto()
		}
	})
}

func BenchmarkScalars_UnmarshalCanoto(b *testing.B) {
	b.Run("full", func(b *testing.B) {
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
			String:   "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
			LargestFieldNumber: LargestFieldNumber{
				Int32: 216457,
			},
			RepeatedInt32:    []int32{1, 2, 3},
			RepeatedInt64:    []int64{1, 2, 3},
			RepeatedUint32:   []uint32{1, 2, 3},
			RepeatedUint64:   []uint64{1, 2, 3},
			RepeatedSint32:   []int32{1, 2, 3},
			RepeatedSint64:   []int64{1, 2, 3},
			RepeatedFixed32:  []uint32{1, 2, 3},
			RepeatedFixed64:  []uint64{1, 2, 3},
			RepeatedSfixed32: []int32{1, 2, 3},
			RepeatedSfixed64: []int64{1, 2, 3},
			RepeatedBool:     []bool{true, false, true},
			RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
			RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
		}
		bytes := cbScalars.MarshalCanoto()

		for _, unsafe := range []bool{false, true} {
			b.Run("unsafe="+strconv.FormatBool(unsafe), func(b *testing.B) {
				for range b.N {
					var (
						scalars Scalars
						reader  = canoto.Reader{
							B:      bytes,
							Unsafe: unsafe,
						}
					)
					_ = scalars.UnmarshalCanotoFrom(&reader)
				}
			})
		}
	})
	b.Run("primitives", func(b *testing.B) {
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
			String:   "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
			LargestFieldNumber: LargestFieldNumber{
				Int32: 216457,
			},
		}
		bytes := cbScalars.MarshalCanoto()

		for _, unsafe := range []bool{false, true} {
			b.Run("unsafe="+strconv.FormatBool(unsafe), func(b *testing.B) {
				for range b.N {
					var (
						scalars Scalars
						reader  = canoto.Reader{
							B:      bytes,
							Unsafe: unsafe,
						}
					)
					_ = scalars.UnmarshalCanotoFrom(&reader)
				}
			})
		}
	})
}

func BenchmarkScalars_MarshalProto(b *testing.B) {
	b.Run("full stack", func(b *testing.B) {
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
				String_:  "hi my name is Bob",
				Bytes:    []byte("hi my name is Bob too"),
				LargestFieldNumber: &pb.LargestFieldNumber{
					Int32: 216457,
				},
				RepeatedInt32:    []int32{1, 2, 3},
				RepeatedInt64:    []int64{1, 2, 3},
				RepeatedUint32:   []uint32{1, 2, 3},
				RepeatedUint64:   []uint64{1, 2, 3},
				RepeatedSint32:   []int32{1, 2, 3},
				RepeatedSint64:   []int64{1, 2, 3},
				RepeatedFixed32:  []uint32{1, 2, 3},
				RepeatedFixed64:  []uint64{1, 2, 3},
				RepeatedSfixed32: []int32{1, 2, 3},
				RepeatedSfixed64: []int64{1, 2, 3},
				RepeatedBool:     []bool{true, false, true},
				RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
				RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
			}
			_, _ = proto.Marshal(&pbScalars)
		}
	})
	b.Run("full heap", func(b *testing.B) {
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
			String_:  "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
			LargestFieldNumber: &pb.LargestFieldNumber{
				Int32: 216457,
			},
			RepeatedInt32:    []int32{1, 2, 3},
			RepeatedInt64:    []int64{1, 2, 3},
			RepeatedUint32:   []uint32{1, 2, 3},
			RepeatedUint64:   []uint64{1, 2, 3},
			RepeatedSint32:   []int32{1, 2, 3},
			RepeatedSint64:   []int64{1, 2, 3},
			RepeatedFixed32:  []uint32{1, 2, 3},
			RepeatedFixed64:  []uint64{1, 2, 3},
			RepeatedSfixed32: []int32{1, 2, 3},
			RepeatedSfixed64: []int64{1, 2, 3},
			RepeatedBool:     []bool{true, false, true},
			RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
			RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
		}
		for range b.N {
			_, _ = proto.Marshal(&pbScalars)
		}
	})
	b.Run("primitives stack", func(b *testing.B) {
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
				String_:  "hi my name is Bob",
				Bytes:    []byte("hi my name is Bob too"),
				LargestFieldNumber: &pb.LargestFieldNumber{
					Int32: 216457,
				},
			}
			_, _ = proto.Marshal(&pbScalars)
		}
	})
	b.Run("primitives heap", func(b *testing.B) {
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
			String_:  "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
			LargestFieldNumber: &pb.LargestFieldNumber{
				Int32: 216457,
			},
		}
		for range b.N {
			_, _ = proto.Marshal(&pbScalars)
		}
	})
}

func BenchmarkScalars_UnmarshalProto(b *testing.B) {
	b.Run("full", func(b *testing.B) {
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
			String_:  "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
			LargestFieldNumber: &pb.LargestFieldNumber{
				Int32: 216457,
			},
			RepeatedInt32:    []int32{1, 2, 3},
			RepeatedInt64:    []int64{1, 2, 3},
			RepeatedUint32:   []uint32{1, 2, 3},
			RepeatedUint64:   []uint64{1, 2, 3},
			RepeatedSint32:   []int32{1, 2, 3},
			RepeatedSint64:   []int64{1, 2, 3},
			RepeatedFixed32:  []uint32{1, 2, 3},
			RepeatedFixed64:  []uint64{1, 2, 3},
			RepeatedSfixed32: []int32{1, 2, 3},
			RepeatedSfixed64: []int64{1, 2, 3},
			RepeatedBool:     []bool{true, false, true},
			RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
			RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
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
	})
	b.Run("primitives", func(b *testing.B) {
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
			String_:  "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
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
	})
}
