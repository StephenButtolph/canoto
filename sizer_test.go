package canoto

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func FuzzSizer_Tag(f *testing.F) {
	f.Fuzz(func(t *testing.T, fieldNumber uint32, wireTypeByte byte) {
		wireType := WireType(wireTypeByte)
		if fieldNumber > MaxFieldNumber || !wireType.IsValid() {
			return
		}

		w := &Writer{}
		w.Tag(fieldNumber, wireType)

		size := SizeTag(fieldNumber, wireType)
		require.Len(t, w.b, size)
	})
}

func FuzzSizer_Int32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		w := &Writer{}
		w.Int32(v)

		size := SizeInt32(v)
		require.Len(t, w.b, size)
	})
}

func FuzzSizer_Int64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		w := &Writer{}
		w.Int64(v)

		size := SizeInt64(v)
		require.Len(t, w.b, size)
	})
}

func FuzzSizer_Uint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		w := &Writer{}
		w.Uint32(v)

		size := SizeUint32(v)
		require.Len(t, w.b, size)
	})
}

func FuzzSizer_Uint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		w := &Writer{}
		w.Uint64(v)

		size := SizeUint64(v)
		require.Len(t, w.b, size)
	})
}

func FuzzSizer_Sint32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		w := &Writer{}
		w.Sint32(v)

		size := SizeSint32(v)
		require.Len(t, w.b, size)
	})
}

func FuzzSizer_Sint64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		w := &Writer{}
		w.Sint64(v)

		size := SizeSint64(v)
		require.Len(t, w.b, size)
	})
}

func FuzzSizer_Fixed32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint32) {
		w := &Writer{}
		w.Fixed32(v)
		require.Len(t, w.b, SizeFixed32)
	})
}

func FuzzSizer_Fixed64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		w := &Writer{}
		w.Fixed64(v)
		require.Len(t, w.b, SizeFixed64)
	})
}

func FuzzSizer_Sfixed32(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int32) {
		w := &Writer{}
		w.Sfixed32(v)
		require.Len(t, w.b, SizeSfixed32)
	})
}

func FuzzSizer_Sfixed64(f *testing.F) {
	f.Fuzz(func(t *testing.T, v int64) {
		w := &Writer{}
		w.Sfixed64(v)
		require.Len(t, w.b, SizeSfixed64)
	})
}

func FuzzSizer_Bool(f *testing.F) {
	f.Fuzz(func(t *testing.T, v bool) {
		w := &Writer{}
		w.Bool(v)
		require.Len(t, w.b, SizeBool)
	})
}

func FuzzSizer_String(f *testing.F) {
	f.Fuzz(func(t *testing.T, v string) {
		w := &Writer{}
		w.String(v)

		size := SizeString(v)
		require.Len(t, w.b, size)
	})
}

func FuzzSizer_Bytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, v []byte) {
		w := &Writer{}
		w.Bytes(v)

		size := SizeBytes(v)
		require.Len(t, w.b, size)
	})
}
