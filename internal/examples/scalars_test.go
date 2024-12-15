package examples

import (
	"slices"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thepudds/fzgen/fuzzer"
	"google.golang.org/protobuf/proto"

	"github.com/StephenButtolph/canoto"
	"github.com/StephenButtolph/canoto/internal/proto/pb"
)

func canonicalizeSlice[T any](s []T) []T {
	if len(s) == 0 {
		return nil
	}
	return s
}

func castSlice[I, O canoto.Int](s []I) []O {
	if len(s) == 0 {
		return nil
	}
	newS := make([]O, len(s))
	for i, v := range s {
		newS[i] = O(v)
	}
	return newS
}

func arrayToSlice[T any](s [][32]T) [][]T {
	if len(s) == 0 {
		return nil
	}
	newS := make([][]T, len(s))
	for i, v := range s {
		newS[i] = slices.Clone(v[:])
	}
	return newS
}

func canonicalizeCanotoScalars(s *Scalars) *Scalars {
	s.Bytes = canonicalizeSlice(s.Bytes)
	s.RepeatedInt8 = canonicalizeSlice(s.RepeatedInt8)
	s.RepeatedInt16 = canonicalizeSlice(s.RepeatedInt16)
	s.RepeatedInt32 = canonicalizeSlice(s.RepeatedInt32)
	s.RepeatedInt64 = canonicalizeSlice(s.RepeatedInt64)
	s.RepeatedUint8 = canonicalizeSlice(s.RepeatedUint8)
	s.RepeatedUint16 = canonicalizeSlice(s.RepeatedUint16)
	s.RepeatedUint32 = canonicalizeSlice(s.RepeatedUint32)
	s.RepeatedUint64 = canonicalizeSlice(s.RepeatedUint64)
	s.RepeatedSint8 = canonicalizeSlice(s.RepeatedSint8)
	s.RepeatedSint16 = canonicalizeSlice(s.RepeatedSint16)
	s.RepeatedSint32 = canonicalizeSlice(s.RepeatedSint32)
	s.RepeatedSint64 = canonicalizeSlice(s.RepeatedSint64)
	s.RepeatedFixed32 = canonicalizeSlice(s.RepeatedFixed32)
	s.RepeatedFixed64 = canonicalizeSlice(s.RepeatedFixed64)
	s.RepeatedSfixed32 = canonicalizeSlice(s.RepeatedSfixed32)
	s.RepeatedSfixed64 = canonicalizeSlice(s.RepeatedSfixed64)
	s.RepeatedBool = canonicalizeSlice(s.RepeatedBool)
	s.RepeatedString = canonicalizeSlice(s.RepeatedString)
	s.RepeatedBytes = canonicalizeSlice(s.RepeatedBytes)
	s.RepeatedLargestFieldNumber = canonicalizeSlice(s.RepeatedLargestFieldNumber)
	s.RepeatedFixedBytes = canonicalizeSlice(s.RepeatedFixedBytes)
	for i := range s.FixedRepeatedBytes {
		s.FixedRepeatedBytes[i] = canonicalizeSlice(s.FixedRepeatedBytes[i])
	}
	if s.CustomType.CalculateCanotoSize() == 0 {
		s.CustomType.Int = nil
	}
	s.CustomBytes = canonicalizeSlice(s.CustomBytes)
	s.CustomRepeatedBytes = canonicalizeSlice(s.CustomRepeatedBytes)
	s.CustomRepeatedFixedBytes = canonicalizeSlice(s.CustomRepeatedFixedBytes)
	for i := range s.CustomFixedRepeatedBytes {
		s.CustomFixedRepeatedBytes[i] = canonicalizeSlice(s.CustomFixedRepeatedBytes[i])
	}
	s.canotoData = canotoData_Scalars{}
	return s
}

func canonicalizeProtoScalars(s *pb.Scalars) *pb.Scalars {
	var largestFieldNumber *pb.LargestFieldNumber
	if s.LargestFieldNumber != nil {
		largestFieldNumber = &pb.LargestFieldNumber{
			Int32: s.LargestFieldNumber.Int32,
		}
	}
	repeatedLargestFieldNumbers := make([]*pb.LargestFieldNumber, len(s.RepeatedLargestFieldNumber))
	for i, v := range s.RepeatedLargestFieldNumber {
		var largestFieldNumber *pb.LargestFieldNumber
		if v != nil {
			largestFieldNumber = &pb.LargestFieldNumber{
				Int32: v.Int32,
			}
		}
		repeatedLargestFieldNumbers[i] = largestFieldNumber
	}
	fixedRepeatedLargestFieldNumber := make([]*pb.LargestFieldNumber, len(s.FixedRepeatedLargestFieldNumber))
	for i, v := range s.FixedRepeatedLargestFieldNumber {
		var largestFieldNumber *pb.LargestFieldNumber
		if v != nil {
			largestFieldNumber = &pb.LargestFieldNumber{
				Int32: v.Int32,
			}
		}
		fixedRepeatedLargestFieldNumber[i] = largestFieldNumber
	}
	fixedRepeatedBytes := make([][]byte, len(s.FixedRepeatedBytes))
	for i, v := range s.FixedRepeatedBytes {
		fixedRepeatedBytes[i] = canonicalizeSlice(v)
	}
	customFixedRepeatedBytes := make([][]byte, len(s.CustomFixedRepeatedBytes))
	for i, v := range s.CustomFixedRepeatedBytes {
		customFixedRepeatedBytes[i] = canonicalizeSlice(v)
	}
	return &pb.Scalars{
		Int8:               s.Int8,
		Int16:              s.Int16,
		Int32:              s.Int32,
		Int64:              s.Int64,
		Uint8:              s.Uint8,
		Uint16:             s.Uint16,
		Uint32:             s.Uint32,
		Uint64:             s.Uint64,
		Sint8:              s.Sint8,
		Sint16:             s.Sint16,
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

		RepeatedInt8:               s.RepeatedInt8,
		RepeatedInt16:              s.RepeatedInt16,
		RepeatedInt32:              s.RepeatedInt32,
		RepeatedInt64:              s.RepeatedInt64,
		RepeatedUint8:              s.RepeatedUint8,
		RepeatedUint16:             s.RepeatedUint16,
		RepeatedUint32:             s.RepeatedUint32,
		RepeatedUint64:             s.RepeatedUint64,
		RepeatedSint8:              s.RepeatedSint8,
		RepeatedSint16:             s.RepeatedSint16,
		RepeatedSint32:             s.RepeatedSint32,
		RepeatedSint64:             s.RepeatedSint64,
		RepeatedFixed32:            s.RepeatedFixed32,
		RepeatedFixed64:            s.RepeatedFixed64,
		RepeatedSfixed32:           s.RepeatedSfixed32,
		RepeatedSfixed64:           s.RepeatedSfixed64,
		RepeatedBool:               s.RepeatedBool,
		RepeatedString:             s.RepeatedString,
		RepeatedBytes:              s.RepeatedBytes,
		RepeatedLargestFieldNumber: canonicalizeSlice(repeatedLargestFieldNumbers),

		FixedRepeatedInt8:               s.FixedRepeatedInt8,
		FixedRepeatedInt16:              s.FixedRepeatedInt16,
		FixedRepeatedInt32:              s.FixedRepeatedInt32,
		FixedRepeatedInt64:              s.FixedRepeatedInt64,
		FixedRepeatedUint8:              s.FixedRepeatedUint8,
		FixedRepeatedUint16:             s.FixedRepeatedUint16,
		FixedRepeatedUint32:             s.FixedRepeatedUint32,
		FixedRepeatedUint64:             s.FixedRepeatedUint64,
		FixedRepeatedSint8:              s.FixedRepeatedSint8,
		FixedRepeatedSint16:             s.FixedRepeatedSint16,
		FixedRepeatedSint32:             s.FixedRepeatedSint32,
		FixedRepeatedSint64:             s.FixedRepeatedSint64,
		FixedRepeatedFixed32:            s.FixedRepeatedFixed32,
		FixedRepeatedFixed64:            s.FixedRepeatedFixed64,
		FixedRepeatedSfixed32:           s.FixedRepeatedSfixed32,
		FixedRepeatedSfixed64:           s.FixedRepeatedSfixed64,
		FixedRepeatedBool:               s.FixedRepeatedBool,
		FixedRepeatedString:             s.FixedRepeatedString,
		FixedBytes:                      s.FixedBytes,
		RepeatedFixedBytes:              s.RepeatedFixedBytes,
		FixedRepeatedBytes:              canonicalizeSlice(fixedRepeatedBytes),
		FixedRepeatedFixedBytes:         s.FixedRepeatedFixedBytes,
		FixedRepeatedLargestFieldNumber: canonicalizeSlice(fixedRepeatedLargestFieldNumber),

		ConstRepeatedUint64:           s.ConstRepeatedUint64,
		CustomType:                    s.CustomType,
		CustomUint32:                  s.CustomUint32,
		CustomBytes:                   s.CustomBytes,
		CustomFixedBytes:              s.CustomFixedBytes,
		CustomRepeatedBytes:           s.CustomRepeatedBytes,
		CustomRepeatedFixedBytes:      s.CustomRepeatedFixedBytes,
		CustomFixedRepeatedBytes:      canonicalizeSlice(customFixedRepeatedBytes),
		CustomFixedRepeatedFixedBytes: s.CustomFixedRepeatedFixedBytes,
	}
}

func canotoScalarsToProto(s *Scalars) *pb.Scalars {
	var largestFieldNumber *pb.LargestFieldNumber
	if s.LargestFieldNumber.Int32 != 0 {
		largestFieldNumber = &pb.LargestFieldNumber{
			Int32: s.LargestFieldNumber.Int32,
		}
	}
	repeatedLargestFieldNumbers := make([]*pb.LargestFieldNumber, len(s.RepeatedLargestFieldNumber))
	for i := range s.RepeatedLargestFieldNumber {
		v := &s.RepeatedLargestFieldNumber[i]

		repeatedLargestFieldNumbers[i] = &pb.LargestFieldNumber{
			Int32: v.Int32,
		}
	}
	var (
		fixedLargestFieldNumbers = make([]*pb.LargestFieldNumber, len(s.FixedRepeatedLargestFieldNumber))
		isZero                   = true
	)
	for i := range s.FixedRepeatedLargestFieldNumber {
		v := &s.FixedRepeatedLargestFieldNumber[i]

		fixedLargestFieldNumbers[i] = &pb.LargestFieldNumber{
			Int32: v.Int32,
		}
		isZero = isZero && v.Int32 == 0
	}
	if isZero {
		fixedLargestFieldNumbers = nil
	}

	var customType []byte
	if s.CustomType.CalculateCanotoSize() != 0 {
		customType = s.CustomType.Int.Bytes()
	}
	pbs := pb.Scalars{
		Int8:               int32(s.Int8),
		Int16:              int32(s.Int16),
		Int32:              s.Int32,
		Int64:              s.Int64,
		Uint8:              uint32(s.Uint8),
		Uint16:             uint32(s.Uint16),
		Uint32:             s.Uint32,
		Uint64:             s.Uint64,
		Sint8:              int32(s.Sint8),
		Sint16:             int32(s.Sint16),
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

		RepeatedInt8:               castSlice[int8, int32](s.RepeatedInt8),
		RepeatedInt16:              castSlice[int16, int32](s.RepeatedInt16),
		RepeatedInt32:              s.RepeatedInt32,
		RepeatedInt64:              s.RepeatedInt64,
		RepeatedUint8:              castSlice[uint8, uint32](s.RepeatedUint8),
		RepeatedUint16:             castSlice[uint16, uint32](s.RepeatedUint16),
		RepeatedUint32:             s.RepeatedUint32,
		RepeatedUint64:             s.RepeatedUint64,
		RepeatedSint8:              castSlice[int8, int32](s.RepeatedSint8),
		RepeatedSint16:             castSlice[int16, int32](s.RepeatedSint16),
		RepeatedSint32:             s.RepeatedSint32,
		RepeatedSint64:             s.RepeatedSint64,
		RepeatedFixed32:            s.RepeatedFixed32,
		RepeatedFixed64:            s.RepeatedFixed64,
		RepeatedSfixed32:           s.RepeatedSfixed32,
		RepeatedSfixed64:           s.RepeatedSfixed64,
		RepeatedBool:               s.RepeatedBool,
		RepeatedString:             s.RepeatedString,
		RepeatedBytes:              s.RepeatedBytes,
		RepeatedLargestFieldNumber: canonicalizeSlice(repeatedLargestFieldNumbers),

		RepeatedFixedBytes:              arrayToSlice(s.RepeatedFixedBytes),
		FixedRepeatedLargestFieldNumber: fixedLargestFieldNumbers,
		CustomType:                      customType,
		CustomUint32:                    uint32(s.CustomUint32),
		CustomBytes:                     s.CustomBytes,
		CustomRepeatedBytes:             s.CustomRepeatedBytes,
		CustomRepeatedFixedBytes:        arrayToSlice(s.CustomRepeatedFixedBytes),
	}
	if !canoto.IsZero(s.FixedRepeatedInt8) {
		pbs.FixedRepeatedInt8 = castSlice[int8, int32](s.FixedRepeatedInt8[:])
	}
	if !canoto.IsZero(s.FixedRepeatedInt16) {
		pbs.FixedRepeatedInt16 = castSlice[int16, int32](s.FixedRepeatedInt16[:])
	}
	if !canoto.IsZero(s.FixedRepeatedInt32) {
		pbs.FixedRepeatedInt32 = slices.Clone(s.FixedRepeatedInt32[:])
	}
	if !canoto.IsZero(s.FixedRepeatedInt64) {
		pbs.FixedRepeatedInt64 = slices.Clone(s.FixedRepeatedInt64[:])
	}
	if !canoto.IsZero(s.FixedRepeatedUint8) {
		pbs.FixedRepeatedUint8 = castSlice[uint8, uint32](s.FixedRepeatedUint8[:])
	}
	if !canoto.IsZero(s.FixedRepeatedUint16) {
		pbs.FixedRepeatedUint16 = castSlice[uint16, uint32](s.FixedRepeatedUint16[:])
	}
	if !canoto.IsZero(s.FixedRepeatedUint32) {
		pbs.FixedRepeatedUint32 = slices.Clone(s.FixedRepeatedUint32[:])
	}
	if !canoto.IsZero(s.FixedRepeatedUint64) {
		pbs.FixedRepeatedUint64 = slices.Clone(s.FixedRepeatedUint64[:])
	}
	if !canoto.IsZero(s.FixedRepeatedSint8) {
		pbs.FixedRepeatedSint8 = castSlice[int8, int32](s.FixedRepeatedSint8[:])
	}
	if !canoto.IsZero(s.FixedRepeatedSint16) {
		pbs.FixedRepeatedSint16 = castSlice[int16, int32](s.FixedRepeatedSint16[:])
	}
	if !canoto.IsZero(s.FixedRepeatedSint32) {
		pbs.FixedRepeatedSint32 = slices.Clone(s.FixedRepeatedSint32[:])
	}
	if !canoto.IsZero(s.FixedRepeatedSint64) {
		pbs.FixedRepeatedSint64 = slices.Clone(s.FixedRepeatedSint64[:])
	}
	if !canoto.IsZero(s.FixedRepeatedFixed32) {
		pbs.FixedRepeatedFixed32 = slices.Clone(s.FixedRepeatedFixed32[:])
	}
	if !canoto.IsZero(s.FixedRepeatedFixed64) {
		pbs.FixedRepeatedFixed64 = slices.Clone(s.FixedRepeatedFixed64[:])
	}
	if !canoto.IsZero(s.FixedRepeatedSfixed32) {
		pbs.FixedRepeatedSfixed32 = slices.Clone(s.FixedRepeatedSfixed32[:])
	}
	if !canoto.IsZero(s.FixedRepeatedSfixed64) {
		pbs.FixedRepeatedSfixed64 = slices.Clone(s.FixedRepeatedSfixed64[:])
	}
	if !canoto.IsZero(s.FixedRepeatedBool) {
		pbs.FixedRepeatedBool = slices.Clone(s.FixedRepeatedBool[:])
	}
	if !canoto.IsZero(s.FixedRepeatedString) {
		pbs.FixedRepeatedString = slices.Clone(s.FixedRepeatedString[:])
	}
	if !canoto.IsZero(s.FixedBytes) {
		pbs.FixedBytes = slices.Clone(s.FixedBytes[:])
	}
	{
		isZero := true
		for _, v := range s.FixedRepeatedBytes {
			isZero = isZero && len(v) == 0
		}
		if !isZero {
			for _, v := range s.FixedRepeatedBytes {
				pbs.FixedRepeatedBytes = append(pbs.FixedRepeatedBytes, canonicalizeSlice(v))
			}
		}
	}
	if !canoto.IsZero(s.FixedRepeatedFixedBytes) {
		for _, v := range s.FixedRepeatedFixedBytes {
			pbs.FixedRepeatedFixedBytes = append(pbs.FixedRepeatedFixedBytes, slices.Clone(v[:]))
		}
	}
	if !canoto.IsZero(s.ConstRepeatedUint64) {
		pbs.ConstRepeatedUint64 = slices.Clone(s.ConstRepeatedUint64[:])
	}
	if !canoto.IsZero(s.CustomFixedBytes) {
		pbs.CustomFixedBytes = slices.Clone(s.CustomFixedBytes[:])
	}
	{
		isZero := true
		for _, v := range s.CustomFixedRepeatedBytes {
			isZero = isZero && len(v) == 0
		}
		if !isZero {
			for _, v := range s.CustomFixedRepeatedBytes {
				pbs.CustomFixedRepeatedBytes = append(pbs.CustomFixedRepeatedBytes, canonicalizeSlice(v))
			}
		}
	}
	if !canoto.IsZero(s.CustomFixedRepeatedFixedBytes) {
		for _, v := range s.CustomFixedRepeatedFixedBytes {
			pbs.CustomFixedRepeatedFixedBytes = append(pbs.CustomFixedRepeatedFixedBytes, slices.Clone(v[:]))
		}
	}
	return &pbs
}

func FuzzScalars_UnmarshalCanoto(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		require := require.New(t)

		canotoScalars := &Scalars{}
		fz := fuzzer.NewFuzzer(data)
		fz.Fill(&canotoScalars)
		canotoScalars = canonicalizeCanotoScalars(canotoScalars)

		pbScalars := canotoScalarsToProto(canotoScalars)
		pbScalarsBytes, err := proto.Marshal(pbScalars)
		if err != nil {
			return
		}

		canotoScalarsFromProto := &Scalars{}
		require.NoError(canotoScalarsFromProto.UnmarshalCanoto(pbScalarsBytes))
		require.Equal(
			canotoScalars,
			canonicalizeCanotoScalars(canotoScalarsFromProto),
		)
	})
}

func FuzzScalars_MarshalCanoto(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		require := require.New(t)

		canotoScalars := &Scalars{}
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
		if err := scalars.UnmarshalCanoto(b); err != nil {
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
				Int8:     31,
				Int16:    2164,
				Int32:    216457,
				Int64:    -2138746,
				Uint8:    254,
				Uint16:   21645,
				Uint32:   32485976,
				Uint64:   287634,
				Sint8:    -31,
				Sint16:   -2164,
				Sint32:   -12786345,
				Sint64:   98761243,
				Fixed32:  98765234,
				Fixed64:  1234576,
				Sfixed32: -21348976,
				Sfixed64: 98756432,
				Bool:     true,
				String:   "hi my name is Bob",
				Bytes:    []byte("hi my name is Bob too"),
				LargestFieldNumber: LargestFieldNumber[int32]{
					Int32: 216457,
				},

				RepeatedInt8:     []int8{1, 2, 3},
				RepeatedInt16:    []int16{1, 2, 3},
				RepeatedInt32:    []int32{1, 2, 3},
				RepeatedInt64:    []int64{1, 2, 3},
				RepeatedUint8:    []uint8{1, 2, 3},
				RepeatedUint16:   []uint16{1, 2, 3},
				RepeatedUint32:   []uint32{1, 2, 3},
				RepeatedUint64:   []uint64{1, 2, 3},
				RepeatedSint8:    []int8{1, 2, 3},
				RepeatedSint16:   []int16{1, 2, 3},
				RepeatedSint32:   []int32{1, 2, 3},
				RepeatedSint64:   []int64{1, 2, 3},
				RepeatedFixed32:  []uint32{1, 2, 3},
				RepeatedFixed64:  []uint64{1, 2, 3},
				RepeatedSfixed32: []int32{1, 2, 3},
				RepeatedSfixed64: []int64{1, 2, 3},
				RepeatedBool:     []bool{true, false, true},
				RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
				RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
				RepeatedLargestFieldNumber: []LargestFieldNumber[int32]{
					{Int32: 123455},
					{Int32: 876523},
				},

				FixedRepeatedInt8:       [3]int8{1, 2, 3},
				FixedRepeatedInt16:      [3]int16{1, 2, 3},
				FixedRepeatedInt32:      [3]int32{1, 2, 3},
				FixedRepeatedInt64:      [3]int64{1, 2, 3},
				FixedRepeatedUint8:      [3]uint8{1, 2, 3},
				FixedRepeatedUint16:     [3]uint16{1, 2, 3},
				FixedRepeatedUint32:     [3]uint32{1, 2, 3},
				FixedRepeatedUint64:     [3]uint64{1, 2, 3},
				FixedRepeatedSint8:      [3]int8{1, 2, 3},
				FixedRepeatedSint16:     [3]int16{1, 2, 3},
				FixedRepeatedSint32:     [3]int32{1, 2, 3},
				FixedRepeatedSint64:     [3]int64{1, 2, 3},
				FixedRepeatedFixed32:    [3]uint32{1, 2, 3},
				FixedRepeatedFixed64:    [3]uint64{1, 2, 3},
				FixedRepeatedSfixed32:   [3]int32{1, 2, 3},
				FixedRepeatedSfixed64:   [3]int64{1, 2, 3},
				FixedRepeatedBool:       [3]bool{true, false, true},
				FixedRepeatedString:     [3]string{"hi", "my", "name"},
				FixedBytes:              [32]byte{1},
				RepeatedFixedBytes:      [][32]byte{{1}, {2}, {3}},
				FixedRepeatedBytes:      [3][]byte{{1}, {2}, {3}},
				FixedRepeatedFixedBytes: [3][32]byte{{1}, {2}, {3}},
				FixedRepeatedLargestFieldNumber: [3]LargestFieldNumber[int32]{
					{Int32: 123455},
					{Int32: 876523},
					{Int32: -576214},
				},

				ConstRepeatedUint64: [constRepeatedUint64Len]uint64{1, 2, 3},
			}
			cbScalars.MarshalCanoto()
		}
	})
	b.Run("full heap", func(b *testing.B) {
		cbScalars := Scalars{
			Int8:     31,
			Int16:    2164,
			Int32:    216457,
			Int64:    -2138746,
			Uint8:    254,
			Uint16:   21645,
			Uint32:   32485976,
			Uint64:   287634,
			Sint8:    -31,
			Sint16:   -2164,
			Sint32:   -12786345,
			Sint64:   98761243,
			Fixed32:  98765234,
			Fixed64:  1234576,
			Sfixed32: -21348976,
			Sfixed64: 98756432,
			Bool:     true,
			String:   "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
			LargestFieldNumber: LargestFieldNumber[int32]{
				Int32: 216457,
			},

			RepeatedInt8:     []int8{1, 2, 3},
			RepeatedInt16:    []int16{1, 2, 3},
			RepeatedInt32:    []int32{1, 2, 3},
			RepeatedInt64:    []int64{1, 2, 3},
			RepeatedUint8:    []uint8{1, 2, 3},
			RepeatedUint16:   []uint16{1, 2, 3},
			RepeatedUint32:   []uint32{1, 2, 3},
			RepeatedUint64:   []uint64{1, 2, 3},
			RepeatedSint8:    []int8{1, 2, 3},
			RepeatedSint16:   []int16{1, 2, 3},
			RepeatedSint32:   []int32{1, 2, 3},
			RepeatedSint64:   []int64{1, 2, 3},
			RepeatedFixed32:  []uint32{1, 2, 3},
			RepeatedFixed64:  []uint64{1, 2, 3},
			RepeatedSfixed32: []int32{1, 2, 3},
			RepeatedSfixed64: []int64{1, 2, 3},
			RepeatedBool:     []bool{true, false, true},
			RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
			RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
			RepeatedLargestFieldNumber: []LargestFieldNumber[int32]{
				{Int32: 123455},
				{Int32: 876523},
			},

			FixedRepeatedInt8:       [3]int8{1, 2, 3},
			FixedRepeatedInt16:      [3]int16{1, 2, 3},
			FixedRepeatedInt32:      [3]int32{1, 2, 3},
			FixedRepeatedInt64:      [3]int64{1, 2, 3},
			FixedRepeatedUint8:      [3]uint8{1, 2, 3},
			FixedRepeatedUint16:     [3]uint16{1, 2, 3},
			FixedRepeatedUint32:     [3]uint32{1, 2, 3},
			FixedRepeatedUint64:     [3]uint64{1, 2, 3},
			FixedRepeatedSint8:      [3]int8{1, 2, 3},
			FixedRepeatedSint16:     [3]int16{1, 2, 3},
			FixedRepeatedSint32:     [3]int32{1, 2, 3},
			FixedRepeatedSint64:     [3]int64{1, 2, 3},
			FixedRepeatedFixed32:    [3]uint32{1, 2, 3},
			FixedRepeatedFixed64:    [3]uint64{1, 2, 3},
			FixedRepeatedSfixed32:   [3]int32{1, 2, 3},
			FixedRepeatedSfixed64:   [3]int64{1, 2, 3},
			FixedRepeatedBool:       [3]bool{true, false, true},
			FixedRepeatedString:     [3]string{"hi", "my", "name"},
			FixedBytes:              [32]byte{1},
			RepeatedFixedBytes:      [][32]byte{{1}, {2}, {3}},
			FixedRepeatedBytes:      [3][]byte{{1}, {2}, {3}},
			FixedRepeatedFixedBytes: [3][32]byte{{1}, {2}, {3}},
			FixedRepeatedLargestFieldNumber: [3]LargestFieldNumber[int32]{
				{Int32: 123455},
				{Int32: 876523},
				{Int32: -576214},
			},

			ConstRepeatedUint64: [constRepeatedUint64Len]uint64{1, 2, 3},
		}
		for range b.N {
			cbScalars.MarshalCanoto()
		}
	})
	b.Run("primitives stack", func(b *testing.B) {
		for range b.N {
			cbScalars := Scalars{
				Int8:     31,
				Int16:    2164,
				Int32:    216457,
				Int64:    -2138746,
				Uint8:    254,
				Uint16:   21645,
				Uint32:   32485976,
				Uint64:   287634,
				Sint8:    -31,
				Sint16:   -2164,
				Sint32:   -12786345,
				Sint64:   98761243,
				Fixed32:  98765234,
				Fixed64:  1234576,
				Sfixed32: -21348976,
				Sfixed64: 98756432,
				Bool:     true,
				String:   "hi my name is Bob",
				Bytes:    []byte("hi my name is Bob too"),
				LargestFieldNumber: LargestFieldNumber[int32]{
					Int32: 216457,
				},
			}
			cbScalars.MarshalCanoto()
		}
	})
	b.Run("primitives heap", func(b *testing.B) {
		cbScalars := Scalars{
			Int8:     31,
			Int16:    2164,
			Int32:    216457,
			Int64:    -2138746,
			Uint8:    254,
			Uint16:   21645,
			Uint32:   32485976,
			Uint64:   287634,
			Sint8:    -31,
			Sint16:   -2164,
			Sint32:   -12786345,
			Sint64:   98761243,
			Fixed32:  98765234,
			Fixed64:  1234576,
			Sfixed32: -21348976,
			Sfixed64: 98756432,
			Bool:     true,
			String:   "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
			LargestFieldNumber: LargestFieldNumber[int32]{
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
			Int8:     31,
			Int16:    2164,
			Int32:    216457,
			Int64:    -2138746,
			Uint8:    254,
			Uint16:   21645,
			Uint32:   32485976,
			Uint64:   287634,
			Sint8:    -31,
			Sint16:   -2164,
			Sint32:   -12786345,
			Sint64:   98761243,
			Fixed32:  98765234,
			Fixed64:  1234576,
			Sfixed32: -21348976,
			Sfixed64: 98756432,
			Bool:     true,
			String:   "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
			LargestFieldNumber: LargestFieldNumber[int32]{
				Int32: 216457,
			},

			RepeatedInt8:     []int8{1, 2, 3},
			RepeatedInt16:    []int16{1, 2, 3},
			RepeatedInt32:    []int32{1, 2, 3},
			RepeatedInt64:    []int64{1, 2, 3},
			RepeatedUint8:    []uint8{1, 2, 3},
			RepeatedUint16:   []uint16{1, 2, 3},
			RepeatedUint32:   []uint32{1, 2, 3},
			RepeatedUint64:   []uint64{1, 2, 3},
			RepeatedSint8:    []int8{1, 2, 3},
			RepeatedSint16:   []int16{1, 2, 3},
			RepeatedSint32:   []int32{1, 2, 3},
			RepeatedSint64:   []int64{1, 2, 3},
			RepeatedFixed32:  []uint32{1, 2, 3},
			RepeatedFixed64:  []uint64{1, 2, 3},
			RepeatedSfixed32: []int32{1, 2, 3},
			RepeatedSfixed64: []int64{1, 2, 3},
			RepeatedBool:     []bool{true, false, true},
			RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
			RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
			RepeatedLargestFieldNumber: []LargestFieldNumber[int32]{
				{Int32: 123455},
				{Int32: 876523},
			},

			FixedRepeatedInt8:       [3]int8{1, 2, 3},
			FixedRepeatedInt16:      [3]int16{1, 2, 3},
			FixedRepeatedInt32:      [3]int32{1, 2, 3},
			FixedRepeatedInt64:      [3]int64{1, 2, 3},
			FixedRepeatedUint8:      [3]uint8{1, 2, 3},
			FixedRepeatedUint16:     [3]uint16{1, 2, 3},
			FixedRepeatedUint32:     [3]uint32{1, 2, 3},
			FixedRepeatedUint64:     [3]uint64{1, 2, 3},
			FixedRepeatedSint8:      [3]int8{1, 2, 3},
			FixedRepeatedSint16:     [3]int16{1, 2, 3},
			FixedRepeatedSint32:     [3]int32{1, 2, 3},
			FixedRepeatedSint64:     [3]int64{1, 2, 3},
			FixedRepeatedFixed32:    [3]uint32{1, 2, 3},
			FixedRepeatedFixed64:    [3]uint64{1, 2, 3},
			FixedRepeatedSfixed32:   [3]int32{1, 2, 3},
			FixedRepeatedSfixed64:   [3]int64{1, 2, 3},
			FixedRepeatedBool:       [3]bool{true, false, true},
			FixedRepeatedString:     [3]string{"hi", "my", "name"},
			FixedBytes:              [32]byte{1},
			RepeatedFixedBytes:      [][32]byte{{1}, {2}, {3}},
			FixedRepeatedBytes:      [3][]byte{{1}, {2}, {3}},
			FixedRepeatedFixedBytes: [3][32]byte{{1}, {2}, {3}},
			FixedRepeatedLargestFieldNumber: [3]LargestFieldNumber[int32]{
				{Int32: 123455},
				{Int32: 876523},
				{Int32: -576214},
			},

			ConstRepeatedUint64: [constRepeatedUint64Len]uint64{1, 2, 3},
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
			Int8:     31,
			Int16:    2164,
			Int32:    216457,
			Int64:    -2138746,
			Uint8:    254,
			Uint16:   21645,
			Uint32:   32485976,
			Uint64:   287634,
			Sint8:    -31,
			Sint16:   -2164,
			Sint32:   -12786345,
			Sint64:   98761243,
			Fixed32:  98765234,
			Fixed64:  1234576,
			Sfixed32: -21348976,
			Sfixed64: 98756432,
			Bool:     true,
			String:   "hi my name is Bob",
			Bytes:    []byte("hi my name is Bob too"),
			LargestFieldNumber: LargestFieldNumber[int32]{
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
				Int8:     31,
				Int16:    2164,
				Int32:    216457,
				Int64:    -2138746,
				Uint8:    254,
				Uint16:   21645,
				Uint32:   32485976,
				Uint64:   287634,
				Sint8:    -31,
				Sint16:   -2164,
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

				RepeatedInt8:     []int32{1, 2, 3},
				RepeatedInt16:    []int32{1, 2, 3},
				RepeatedInt32:    []int32{1, 2, 3},
				RepeatedInt64:    []int64{1, 2, 3},
				RepeatedUint8:    []uint32{1, 2, 3},
				RepeatedUint16:   []uint32{1, 2, 3},
				RepeatedUint32:   []uint32{1, 2, 3},
				RepeatedUint64:   []uint64{1, 2, 3},
				RepeatedSint8:    []int32{1, 2, 3},
				RepeatedSint16:   []int32{1, 2, 3},
				RepeatedSint32:   []int32{1, 2, 3},
				RepeatedSint64:   []int64{1, 2, 3},
				RepeatedFixed32:  []uint32{1, 2, 3},
				RepeatedFixed64:  []uint64{1, 2, 3},
				RepeatedSfixed32: []int32{1, 2, 3},
				RepeatedSfixed64: []int64{1, 2, 3},
				RepeatedBool:     []bool{true, false, true},
				RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
				RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
				RepeatedLargestFieldNumber: []*pb.LargestFieldNumber{
					{Int32: 123455},
					{Int32: 876523},
				},

				FixedRepeatedInt8:     []int32{1, 2, 3},
				FixedRepeatedInt16:    []int32{1, 2, 3},
				FixedRepeatedInt32:    []int32{1, 2, 3},
				FixedRepeatedInt64:    []int64{1, 2, 3},
				FixedRepeatedUint8:    []uint32{1, 2, 3},
				FixedRepeatedUint16:   []uint32{1, 2, 3},
				FixedRepeatedUint32:   []uint32{1, 2, 3},
				FixedRepeatedUint64:   []uint64{1, 2, 3},
				FixedRepeatedSint8:    []int32{1, 2, 3},
				FixedRepeatedSint16:   []int32{1, 2, 3},
				FixedRepeatedSint32:   []int32{1, 2, 3},
				FixedRepeatedSint64:   []int64{1, 2, 3},
				FixedRepeatedFixed32:  []uint32{1, 2, 3},
				FixedRepeatedFixed64:  []uint64{1, 2, 3},
				FixedRepeatedSfixed32: []int32{1, 2, 3},
				FixedRepeatedSfixed64: []int64{1, 2, 3},
				FixedRepeatedBool:     []bool{true, false, true},
				FixedRepeatedString:   []string{"hi", "my", "name"},
				FixedBytes:            []byte{0: 1, 31: 0},
				RepeatedFixedBytes: [][]byte{
					{0: 1, 31: 0},
					{0: 2, 31: 0},
					{0: 3, 31: 0},
				},
				FixedRepeatedBytes: [][]byte{{1}, {2}, {3}},
				FixedRepeatedFixedBytes: [][]byte{
					{0: 1, 31: 0},
					{0: 2, 31: 0},
					{0: 3, 31: 0},
				},
				FixedRepeatedLargestFieldNumber: []*pb.LargestFieldNumber{
					{Int32: 123455},
					{Int32: 876523},
					{Int32: -576214},
				},

				ConstRepeatedUint64: []uint64{1, 2, 3},
			}
			_, _ = proto.Marshal(&pbScalars)
		}
	})
	b.Run("full heap", func(b *testing.B) {
		pbScalars := pb.Scalars{
			Int8:     31,
			Int16:    2164,
			Int32:    216457,
			Int64:    -2138746,
			Uint8:    254,
			Uint16:   21645,
			Uint32:   32485976,
			Uint64:   287634,
			Sint8:    -31,
			Sint16:   -2164,
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

			RepeatedInt8:     []int32{1, 2, 3},
			RepeatedInt16:    []int32{1, 2, 3},
			RepeatedInt32:    []int32{1, 2, 3},
			RepeatedInt64:    []int64{1, 2, 3},
			RepeatedUint8:    []uint32{1, 2, 3},
			RepeatedUint16:   []uint32{1, 2, 3},
			RepeatedUint32:   []uint32{1, 2, 3},
			RepeatedUint64:   []uint64{1, 2, 3},
			RepeatedSint8:    []int32{1, 2, 3},
			RepeatedSint16:   []int32{1, 2, 3},
			RepeatedSint32:   []int32{1, 2, 3},
			RepeatedSint64:   []int64{1, 2, 3},
			RepeatedFixed32:  []uint32{1, 2, 3},
			RepeatedFixed64:  []uint64{1, 2, 3},
			RepeatedSfixed32: []int32{1, 2, 3},
			RepeatedSfixed64: []int64{1, 2, 3},
			RepeatedBool:     []bool{true, false, true},
			RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
			RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
			RepeatedLargestFieldNumber: []*pb.LargestFieldNumber{
				{Int32: 123455},
				{Int32: 876523},
			},

			FixedRepeatedInt8:     []int32{1, 2, 3},
			FixedRepeatedInt16:    []int32{1, 2, 3},
			FixedRepeatedInt32:    []int32{1, 2, 3},
			FixedRepeatedInt64:    []int64{1, 2, 3},
			FixedRepeatedUint8:    []uint32{1, 2, 3},
			FixedRepeatedUint16:   []uint32{1, 2, 3},
			FixedRepeatedUint32:   []uint32{1, 2, 3},
			FixedRepeatedUint64:   []uint64{1, 2, 3},
			FixedRepeatedSint8:    []int32{1, 2, 3},
			FixedRepeatedSint16:   []int32{1, 2, 3},
			FixedRepeatedSint32:   []int32{1, 2, 3},
			FixedRepeatedSint64:   []int64{1, 2, 3},
			FixedRepeatedFixed32:  []uint32{1, 2, 3},
			FixedRepeatedFixed64:  []uint64{1, 2, 3},
			FixedRepeatedSfixed32: []int32{1, 2, 3},
			FixedRepeatedSfixed64: []int64{1, 2, 3},
			FixedRepeatedBool:     []bool{true, false, true},
			FixedRepeatedString:   []string{"hi", "my", "name"},
			FixedBytes:            []byte{0: 1, 31: 0},
			RepeatedFixedBytes: [][]byte{
				{0: 1, 31: 0},
				{0: 2, 31: 0},
				{0: 3, 31: 0},
			},
			FixedRepeatedBytes: [][]byte{{1}, {2}, {3}},
			FixedRepeatedFixedBytes: [][]byte{
				{0: 1, 31: 0},
				{0: 2, 31: 0},
				{0: 3, 31: 0},
			},
			FixedRepeatedLargestFieldNumber: []*pb.LargestFieldNumber{
				{Int32: 123455},
				{Int32: 876523},
				{Int32: -576214},
			},

			ConstRepeatedUint64: []uint64{1, 2, 3},
		}
		for range b.N {
			_, _ = proto.Marshal(&pbScalars)
		}
	})
	b.Run("primitives stack", func(b *testing.B) {
		for range b.N {
			pbScalars := pb.Scalars{
				Int8:     31,
				Int16:    2164,
				Int32:    216457,
				Int64:    -2138746,
				Uint8:    254,
				Uint16:   21645,
				Uint32:   32485976,
				Uint64:   287634,
				Sint8:    -31,
				Sint16:   -2164,
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
			Int8:     31,
			Int16:    2164,
			Int32:    216457,
			Int64:    -2138746,
			Uint8:    254,
			Uint16:   21645,
			Uint32:   32485976,
			Uint64:   287634,
			Sint8:    -31,
			Sint16:   -2164,
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
			Int8:     31,
			Int16:    2164,
			Int32:    216457,
			Int64:    -2138746,
			Uint8:    254,
			Uint16:   21645,
			Uint32:   32485976,
			Uint64:   287634,
			Sint8:    -31,
			Sint16:   -2164,
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

			RepeatedInt8:     []int32{1, 2, 3},
			RepeatedInt16:    []int32{1, 2, 3},
			RepeatedInt32:    []int32{1, 2, 3},
			RepeatedInt64:    []int64{1, 2, 3},
			RepeatedUint8:    []uint32{1, 2, 3},
			RepeatedUint16:   []uint32{1, 2, 3},
			RepeatedUint32:   []uint32{1, 2, 3},
			RepeatedUint64:   []uint64{1, 2, 3},
			RepeatedSint8:    []int32{1, 2, 3},
			RepeatedSint16:   []int32{1, 2, 3},
			RepeatedSint32:   []int32{1, 2, 3},
			RepeatedSint64:   []int64{1, 2, 3},
			RepeatedFixed32:  []uint32{1, 2, 3},
			RepeatedFixed64:  []uint64{1, 2, 3},
			RepeatedSfixed32: []int32{1, 2, 3},
			RepeatedSfixed64: []int64{1, 2, 3},
			RepeatedBool:     []bool{true, false, true},
			RepeatedString:   []string{"hi", "my", "name", "is", "Bob"},
			RepeatedBytes:    [][]byte{{1, 2, 3}, {4, 5, 6}},
			RepeatedLargestFieldNumber: []*pb.LargestFieldNumber{
				{Int32: 123455},
				{Int32: 876523},
			},

			FixedRepeatedInt8:     []int32{1, 2, 3},
			FixedRepeatedInt16:    []int32{1, 2, 3},
			FixedRepeatedInt32:    []int32{1, 2, 3},
			FixedRepeatedInt64:    []int64{1, 2, 3},
			FixedRepeatedUint8:    []uint32{1, 2, 3},
			FixedRepeatedUint16:   []uint32{1, 2, 3},
			FixedRepeatedUint32:   []uint32{1, 2, 3},
			FixedRepeatedUint64:   []uint64{1, 2, 3},
			FixedRepeatedSint8:    []int32{1, 2, 3},
			FixedRepeatedSint16:   []int32{1, 2, 3},
			FixedRepeatedSint32:   []int32{1, 2, 3},
			FixedRepeatedSint64:   []int64{1, 2, 3},
			FixedRepeatedFixed32:  []uint32{1, 2, 3},
			FixedRepeatedFixed64:  []uint64{1, 2, 3},
			FixedRepeatedSfixed32: []int32{1, 2, 3},
			FixedRepeatedSfixed64: []int64{1, 2, 3},
			FixedRepeatedBool:     []bool{true, false, true},
			FixedRepeatedString:   []string{"hi", "my", "name"},
			FixedBytes:            []byte{0: 1, 31: 0},
			RepeatedFixedBytes: [][]byte{
				{0: 1, 31: 0},
				{0: 2, 31: 0},
				{0: 3, 31: 0},
			},
			FixedRepeatedBytes: [][]byte{{1}, {2}, {3}},
			FixedRepeatedFixedBytes: [][]byte{
				{0: 1, 31: 0},
				{0: 2, 31: 0},
				{0: 3, 31: 0},
			},
			FixedRepeatedLargestFieldNumber: []*pb.LargestFieldNumber{
				{Int32: 123455},
				{Int32: 876523},
				{Int32: -576214},
			},

			ConstRepeatedUint64: []uint64{1, 2, 3},
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
			Int8:     31,
			Int16:    2164,
			Int32:    216457,
			Int64:    -2138746,
			Uint8:    254,
			Uint16:   21645,
			Uint32:   32485976,
			Uint64:   287634,
			Sint8:    -31,
			Sint16:   -2164,
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
