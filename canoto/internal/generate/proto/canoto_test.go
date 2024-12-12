package proto

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/StephenButtolph/canoto"
	"github.com/StephenButtolph/canoto/canoto/internal/generate/proto/pb"
)

func TestAppend_ProtoCompatibility(t *testing.T) {
	tests := []struct {
		name  string
		proto protoreflect.ProtoMessage
		f     func(*canoto.Writer)
	}{
		{
			name: "int8",
			proto: &pb.Scalars{
				Int8: 52,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(1, canoto.Varint))
				canoto.AppendInt[int8](w, 52)
			},
		},
		{
			name: "int16",
			proto: &pb.Scalars{
				Int16: 1234,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(2, canoto.Varint))
				canoto.AppendInt[int16](w, 1234)
			},
		},
		{
			name: "int32",
			proto: &pb.Scalars{
				Int32: 121234,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(3, canoto.Varint))
				canoto.AppendInt[int32](w, 121234)
			},
		},
		{
			name: "int64",
			proto: &pb.Scalars{
				Int64: 259,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(4, canoto.Varint))
				canoto.AppendInt[int64](w, 259)
			},
		},
		{
			name: "uint8",
			proto: &pb.Scalars{
				Uint8: 9,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(5, canoto.Varint))
				canoto.AppendInt[uint8](w, 9)
			},
		},
		{
			name: "uint16",
			proto: &pb.Scalars{
				Uint16: 1234,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(6, canoto.Varint))
				canoto.AppendInt[uint16](w, 1234)
			},
		},
		{
			name: "uint32",
			proto: &pb.Scalars{
				Uint32: 1234,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(7, canoto.Varint))
				canoto.AppendInt[uint32](w, 1234)
			},
		},
		{
			name: "uint64",
			proto: &pb.Scalars{
				Uint64: 2938567,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(8, canoto.Varint))
				canoto.AppendInt[uint64](w, 2938567)
			},
		},
		{
			name: "sint8",
			proto: &pb.Scalars{
				Sint8: -52,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(9, canoto.Varint))
				canoto.AppendSint[int8](w, -52)
			},
		},
		{
			name: "sint16",
			proto: &pb.Scalars{
				Sint16: -1234,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(10, canoto.Varint))
				canoto.AppendSint[int16](w, -1234)
			},
		},
		{
			name: "sint32",
			proto: &pb.Scalars{
				Sint32: -2136745,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(11, canoto.Varint))
				canoto.AppendSint[int32](w, -2136745)
			},
		},
		{
			name: "sint64",
			proto: &pb.Scalars{
				Sint64: -9287364,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(12, canoto.Varint))
				canoto.AppendSint[int64](w, -9287364)
			},
		},
		{
			name: "fixed32",
			proto: &pb.Scalars{
				Fixed32: 876254,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(13, canoto.I32))
				canoto.AppendFint32[uint32](w, 876254)
			},
		},
		{
			name: "fixed64",
			proto: &pb.Scalars{
				Fixed64: 328137645632,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(14, canoto.I64))
				canoto.AppendFint64[uint64](w, 328137645632)
			},
		},
		{
			name: "sfixed32",
			proto: &pb.Scalars{
				Sfixed32: -123463246,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(15, canoto.I32))
				canoto.AppendFint32[int32](w, -123463246)
			},
		},
		{
			name: "sfixed64",
			proto: &pb.Scalars{
				Sfixed64: -8762135423,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(16, canoto.I64))
				canoto.AppendFint64[int64](w, -8762135423)
			},
		},
		{
			name: "bool",
			proto: &pb.Scalars{
				Bool: true,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(17, canoto.Varint))
				canoto.AppendBool(w, true)
			},
		},
		{
			name: "string",
			proto: &pb.Scalars{
				String_: "hi mom!",
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(18, canoto.Len))
				canoto.AppendBytes(w, "hi mom!")
			},
		},
		{
			name: "bytes",
			proto: &pb.Scalars{
				Bytes: []byte("hi dad!"),
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(19, canoto.Len))
				canoto.AppendBytes(w, []byte("hi dad!"))
			},
		},
		{
			name: "largest field number",
			proto: &pb.LargestFieldNumber{
				Int32: 1,
			},
			f: func(w *canoto.Writer) {
				canoto.Append(w, canoto.Tag(canoto.MaxFieldNumber, canoto.Varint))
				canoto.AppendInt[int32](w, 1)
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pbBytes, err := proto.Marshal(test.proto)
			require.NoError(t, err)

			w := &canoto.Writer{}
			test.f(w)
			require.Equal(t, pbBytes, w.B)
		})
	}
}
