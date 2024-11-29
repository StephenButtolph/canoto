package canoto

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/StephenButtolph/canoto/internal/proto/pb"
)

func TestWriter_ProtoCompatibility(t *testing.T) {
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
				w.Tag(1, Varint)
				w.Int32(128)
			},
		},
		{
			name: "int64",
			proto: &pb.Scalars{
				Int64: 259,
			},
			f: func(w *Writer) {
				w.Tag(2, Varint)
				w.Int64(259)
			},
		},
		{
			name: "uint32",
			proto: &pb.Scalars{
				Uint32: 1234,
			},
			f: func(w *Writer) {
				w.Tag(3, Varint)
				w.Uint32(1234)
			},
		},
		{
			name: "uint64",
			proto: &pb.Scalars{
				Uint64: 2938567,
			},
			f: func(w *Writer) {
				w.Tag(4, Varint)
				w.Uint64(2938567)
			},
		},
		{
			name: "sint32",
			proto: &pb.Scalars{
				Sint32: -2136745,
			},
			f: func(w *Writer) {
				w.Tag(5, Varint)
				w.Sint32(-2136745)
			},
		},
		{
			name: "sint64",
			proto: &pb.Scalars{
				Sint64: -9287364,
			},
			f: func(w *Writer) {
				w.Tag(6, Varint)
				w.Sint64(-9287364)
			},
		},
		{
			name: "fixed32",
			proto: &pb.Scalars{
				Fixed32: 876254,
			},
			f: func(w *Writer) {
				w.Tag(7, I32)
				w.Fixed32(876254)
			},
		},
		{
			name: "fixed64",
			proto: &pb.Scalars{
				Fixed64: 328137645632,
			},
			f: func(w *Writer) {
				w.Tag(8, I64)
				w.Fixed64(328137645632)
			},
		},
		{
			name: "sfixed32",
			proto: &pb.Scalars{
				Sfixed32: -123463246,
			},
			f: func(w *Writer) {
				w.Tag(9, I32)
				w.Sfixed32(-123463246)
			},
		},
		{
			name: "sfixed64",
			proto: &pb.Scalars{
				Sfixed64: -8762135423,
			},
			f: func(w *Writer) {
				w.Tag(10, I64)
				w.Sfixed64(-8762135423)
			},
		},
		{
			name: "bool",
			proto: &pb.Scalars{
				Bool: true,
			},
			f: func(w *Writer) {
				w.Tag(11, Varint)
				w.Bool(true)
			},
		},
		{
			name: "string",
			proto: &pb.Scalars{
				String_: "hi mom!",
			},
			f: func(w *Writer) {
				w.Tag(12, Len)
				w.String("hi mom!")
			},
		},
		{
			name: "bytes",
			proto: &pb.Scalars{
				Bytes: []byte("hi dad!"),
			},
			f: func(w *Writer) {
				w.Tag(13, Len)
				w.Bytes([]byte("hi dad!"))
			},
		},
		{
			name: "largest field number",
			proto: &pb.LargestFieldNumber{
				Int32: 1,
			},
			f: func(w *Writer) {
				w.Tag(MaxFieldNumber, Varint)
				w.Int32(1)
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

func FuzzWriter_Tag(f *testing.F) {
	f.Fuzz(func(t *testing.T, fieldNumber uint32, wireTypeByte byte) {
		wireType := WireType(wireTypeByte)
		if fieldNumber > MaxFieldNumber || !wireType.IsValid() {
			return
		}

		require := require.New(t)

		w := &Writer{}
		w.Tag(fieldNumber, wireType)

		r := &Reader{b: w.b}
		gotFieldNumber, gotWireType, err := r.Tag()
		require.NoError(err)
		require.Equal(fieldNumber, gotFieldNumber)
		require.Equal(wireType, gotWireType)
		require.Empty(r.b)
	})
}

func FuzzWriter_Int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		require := require.New(t)

		w := &Writer{}
		w.Int32(v)

		r := &Reader{b: w.b}
		got, err := r.Int32()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_Int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		require := require.New(t)

		w := &Writer{}
		w.Int64(v)

		r := &Reader{b: w.b}
		got, err := r.Int64()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_Uint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		require := require.New(t)

		w := &Writer{}
		w.Uint32(v)

		r := &Reader{b: w.b}
		got, err := r.Uint32()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_Uint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		require := require.New(t)

		w := &Writer{}
		w.Uint64(v)

		r := &Reader{b: w.b}
		got, err := r.Uint64()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_Sint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		require := require.New(t)

		w := &Writer{}
		w.Sint32(v)

		r := &Reader{b: w.b}
		got, err := r.Sint32()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_Sint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		require := require.New(t)

		w := &Writer{}
		w.Sint64(v)

		r := &Reader{b: w.b}
		got, err := r.Sint64()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_Fixed32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		require := require.New(t)

		w := &Writer{}
		w.Fixed32(v)

		r := &Reader{b: w.b}
		got, err := r.Fixed32()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_Fixed64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		require := require.New(t)

		w := &Writer{}
		w.Fixed64(v)

		r := &Reader{b: w.b}
		got, err := r.Fixed64()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_Sfixed32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		require := require.New(t)

		w := &Writer{}
		w.Sfixed32(v)

		r := &Reader{b: w.b}
		got, err := r.Sfixed32()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_Sfixed64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		require := require.New(t)

		w := &Writer{}
		w.Sfixed64(v)

		r := &Reader{b: w.b}
		got, err := r.Sfixed64()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_Bool(f *testing.F) {
	f.Fuzz(func(t *testing.T, v bool) {
		require := require.New(t)

		w := &Writer{}
		w.Bool(v)

		r := &Reader{b: w.b}
		got, err := r.Bool()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_String(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		require := require.New(t)

		w := &Writer{}
		w.String(v)

		r := &Reader{b: w.b}
		got, err := r.String()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}

func FuzzWriter_Bytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, v []byte) {
		require := require.New(t)

		w := &Writer{}
		w.Bytes(v)

		r := &Reader{b: w.b}
		got, err := r.Bytes()
		require.NoError(err)
		require.Equal(v, got)
		require.Empty(r.b)
	})
}
