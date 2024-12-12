// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: types.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Scalars struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int8                       int32                 `protobuf:"varint,1,opt,name=int8,proto3" json:"int8,omitempty"`
	Int16                      int32                 `protobuf:"varint,2,opt,name=int16,proto3" json:"int16,omitempty"`
	Int32                      int32                 `protobuf:"varint,3,opt,name=int32,proto3" json:"int32,omitempty"`
	Int64                      int64                 `protobuf:"varint,4,opt,name=int64,proto3" json:"int64,omitempty"`
	Uint8                      uint32                `protobuf:"varint,5,opt,name=uint8,proto3" json:"uint8,omitempty"`
	Uint16                     uint32                `protobuf:"varint,6,opt,name=uint16,proto3" json:"uint16,omitempty"`
	Uint32                     uint32                `protobuf:"varint,7,opt,name=uint32,proto3" json:"uint32,omitempty"`
	Uint64                     uint64                `protobuf:"varint,8,opt,name=uint64,proto3" json:"uint64,omitempty"`
	Sint8                      int32                 `protobuf:"zigzag32,9,opt,name=sint8,proto3" json:"sint8,omitempty"`
	Sint16                     int32                 `protobuf:"zigzag32,10,opt,name=sint16,proto3" json:"sint16,omitempty"`
	Sint32                     int32                 `protobuf:"zigzag32,11,opt,name=sint32,proto3" json:"sint32,omitempty"`
	Sint64                     int64                 `protobuf:"zigzag64,12,opt,name=sint64,proto3" json:"sint64,omitempty"`
	Fixed32                    uint32                `protobuf:"fixed32,13,opt,name=fixed32,proto3" json:"fixed32,omitempty"`
	Fixed64                    uint64                `protobuf:"fixed64,14,opt,name=fixed64,proto3" json:"fixed64,omitempty"`
	Sfixed32                   int32                 `protobuf:"fixed32,15,opt,name=sfixed32,proto3" json:"sfixed32,omitempty"`
	Sfixed64                   int64                 `protobuf:"fixed64,16,opt,name=sfixed64,proto3" json:"sfixed64,omitempty"`
	Bool                       bool                  `protobuf:"varint,17,opt,name=bool,proto3" json:"bool,omitempty"`
	String_                    string                `protobuf:"bytes,18,opt,name=string,proto3" json:"string,omitempty"`
	Bytes                      []byte                `protobuf:"bytes,19,opt,name=bytes,proto3" json:"bytes,omitempty"`
	LargestFieldNumber         *LargestFieldNumber   `protobuf:"bytes,20,opt,name=largest_field_number,json=largestFieldNumber,proto3" json:"largest_field_number,omitempty"`
	RepeatedInt8               []int32               `protobuf:"varint,21,rep,packed,name=repeated_int8,json=repeatedInt8,proto3" json:"repeated_int8,omitempty"`
	RepeatedInt16              []int32               `protobuf:"varint,22,rep,packed,name=repeated_int16,json=repeatedInt16,proto3" json:"repeated_int16,omitempty"`
	RepeatedInt32              []int32               `protobuf:"varint,23,rep,packed,name=repeated_int32,json=repeatedInt32,proto3" json:"repeated_int32,omitempty"`
	RepeatedInt64              []int64               `protobuf:"varint,24,rep,packed,name=repeated_int64,json=repeatedInt64,proto3" json:"repeated_int64,omitempty"`
	RepeatedUint8              []uint32              `protobuf:"varint,25,rep,packed,name=repeated_uint8,json=repeatedUint8,proto3" json:"repeated_uint8,omitempty"`
	RepeatedUint16             []uint32              `protobuf:"varint,26,rep,packed,name=repeated_uint16,json=repeatedUint16,proto3" json:"repeated_uint16,omitempty"`
	RepeatedUint32             []uint32              `protobuf:"varint,27,rep,packed,name=repeated_uint32,json=repeatedUint32,proto3" json:"repeated_uint32,omitempty"`
	RepeatedUint64             []uint64              `protobuf:"varint,28,rep,packed,name=repeated_uint64,json=repeatedUint64,proto3" json:"repeated_uint64,omitempty"`
	RepeatedSint8              []int32               `protobuf:"zigzag32,29,rep,packed,name=repeated_sint8,json=repeatedSint8,proto3" json:"repeated_sint8,omitempty"`
	RepeatedSint16             []int32               `protobuf:"zigzag32,30,rep,packed,name=repeated_sint16,json=repeatedSint16,proto3" json:"repeated_sint16,omitempty"`
	RepeatedSint32             []int32               `protobuf:"zigzag32,31,rep,packed,name=repeated_sint32,json=repeatedSint32,proto3" json:"repeated_sint32,omitempty"`
	RepeatedSint64             []int64               `protobuf:"zigzag64,32,rep,packed,name=repeated_sint64,json=repeatedSint64,proto3" json:"repeated_sint64,omitempty"`
	RepeatedFixed32            []uint32              `protobuf:"fixed32,33,rep,packed,name=repeated_fixed32,json=repeatedFixed32,proto3" json:"repeated_fixed32,omitempty"`
	RepeatedFixed64            []uint64              `protobuf:"fixed64,34,rep,packed,name=repeated_fixed64,json=repeatedFixed64,proto3" json:"repeated_fixed64,omitempty"`
	RepeatedSfixed32           []int32               `protobuf:"fixed32,35,rep,packed,name=repeated_sfixed32,json=repeatedSfixed32,proto3" json:"repeated_sfixed32,omitempty"`
	RepeatedSfixed64           []int64               `protobuf:"fixed64,36,rep,packed,name=repeated_sfixed64,json=repeatedSfixed64,proto3" json:"repeated_sfixed64,omitempty"`
	RepeatedBool               []bool                `protobuf:"varint,37,rep,packed,name=repeated_bool,json=repeatedBool,proto3" json:"repeated_bool,omitempty"`
	RepeatedString             []string              `protobuf:"bytes,38,rep,name=repeated_string,json=repeatedString,proto3" json:"repeated_string,omitempty"`
	RepeatedBytes              [][]byte              `protobuf:"bytes,39,rep,name=repeated_bytes,json=repeatedBytes,proto3" json:"repeated_bytes,omitempty"`
	RepeatedLargestFieldNumber []*LargestFieldNumber `protobuf:"bytes,40,rep,name=repeated_largest_field_number,json=repeatedLargestFieldNumber,proto3" json:"repeated_largest_field_number,omitempty"`
}

func (x *Scalars) Reset() {
	*x = Scalars{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Scalars) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Scalars) ProtoMessage() {}

func (x *Scalars) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Scalars.ProtoReflect.Descriptor instead.
func (*Scalars) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *Scalars) GetInt8() int32 {
	if x != nil {
		return x.Int8
	}
	return 0
}

func (x *Scalars) GetInt16() int32 {
	if x != nil {
		return x.Int16
	}
	return 0
}

func (x *Scalars) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

func (x *Scalars) GetInt64() int64 {
	if x != nil {
		return x.Int64
	}
	return 0
}

func (x *Scalars) GetUint8() uint32 {
	if x != nil {
		return x.Uint8
	}
	return 0
}

func (x *Scalars) GetUint16() uint32 {
	if x != nil {
		return x.Uint16
	}
	return 0
}

func (x *Scalars) GetUint32() uint32 {
	if x != nil {
		return x.Uint32
	}
	return 0
}

func (x *Scalars) GetUint64() uint64 {
	if x != nil {
		return x.Uint64
	}
	return 0
}

func (x *Scalars) GetSint8() int32 {
	if x != nil {
		return x.Sint8
	}
	return 0
}

func (x *Scalars) GetSint16() int32 {
	if x != nil {
		return x.Sint16
	}
	return 0
}

func (x *Scalars) GetSint32() int32 {
	if x != nil {
		return x.Sint32
	}
	return 0
}

func (x *Scalars) GetSint64() int64 {
	if x != nil {
		return x.Sint64
	}
	return 0
}

func (x *Scalars) GetFixed32() uint32 {
	if x != nil {
		return x.Fixed32
	}
	return 0
}

func (x *Scalars) GetFixed64() uint64 {
	if x != nil {
		return x.Fixed64
	}
	return 0
}

func (x *Scalars) GetSfixed32() int32 {
	if x != nil {
		return x.Sfixed32
	}
	return 0
}

func (x *Scalars) GetSfixed64() int64 {
	if x != nil {
		return x.Sfixed64
	}
	return 0
}

func (x *Scalars) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *Scalars) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Scalars) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

func (x *Scalars) GetLargestFieldNumber() *LargestFieldNumber {
	if x != nil {
		return x.LargestFieldNumber
	}
	return nil
}

func (x *Scalars) GetRepeatedInt8() []int32 {
	if x != nil {
		return x.RepeatedInt8
	}
	return nil
}

func (x *Scalars) GetRepeatedInt16() []int32 {
	if x != nil {
		return x.RepeatedInt16
	}
	return nil
}

func (x *Scalars) GetRepeatedInt32() []int32 {
	if x != nil {
		return x.RepeatedInt32
	}
	return nil
}

func (x *Scalars) GetRepeatedInt64() []int64 {
	if x != nil {
		return x.RepeatedInt64
	}
	return nil
}

func (x *Scalars) GetRepeatedUint8() []uint32 {
	if x != nil {
		return x.RepeatedUint8
	}
	return nil
}

func (x *Scalars) GetRepeatedUint16() []uint32 {
	if x != nil {
		return x.RepeatedUint16
	}
	return nil
}

func (x *Scalars) GetRepeatedUint32() []uint32 {
	if x != nil {
		return x.RepeatedUint32
	}
	return nil
}

func (x *Scalars) GetRepeatedUint64() []uint64 {
	if x != nil {
		return x.RepeatedUint64
	}
	return nil
}

func (x *Scalars) GetRepeatedSint8() []int32 {
	if x != nil {
		return x.RepeatedSint8
	}
	return nil
}

func (x *Scalars) GetRepeatedSint16() []int32 {
	if x != nil {
		return x.RepeatedSint16
	}
	return nil
}

func (x *Scalars) GetRepeatedSint32() []int32 {
	if x != nil {
		return x.RepeatedSint32
	}
	return nil
}

func (x *Scalars) GetRepeatedSint64() []int64 {
	if x != nil {
		return x.RepeatedSint64
	}
	return nil
}

func (x *Scalars) GetRepeatedFixed32() []uint32 {
	if x != nil {
		return x.RepeatedFixed32
	}
	return nil
}

func (x *Scalars) GetRepeatedFixed64() []uint64 {
	if x != nil {
		return x.RepeatedFixed64
	}
	return nil
}

func (x *Scalars) GetRepeatedSfixed32() []int32 {
	if x != nil {
		return x.RepeatedSfixed32
	}
	return nil
}

func (x *Scalars) GetRepeatedSfixed64() []int64 {
	if x != nil {
		return x.RepeatedSfixed64
	}
	return nil
}

func (x *Scalars) GetRepeatedBool() []bool {
	if x != nil {
		return x.RepeatedBool
	}
	return nil
}

func (x *Scalars) GetRepeatedString() []string {
	if x != nil {
		return x.RepeatedString
	}
	return nil
}

func (x *Scalars) GetRepeatedBytes() [][]byte {
	if x != nil {
		return x.RepeatedBytes
	}
	return nil
}

func (x *Scalars) GetRepeatedLargestFieldNumber() []*LargestFieldNumber {
	if x != nil {
		return x.RepeatedLargestFieldNumber
	}
	return nil
}

type LargestFieldNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32 int32 `protobuf:"varint,536870911,opt,name=int32,proto3" json:"int32,omitempty"`
}

func (x *LargestFieldNumber) Reset() {
	*x = LargestFieldNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LargestFieldNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LargestFieldNumber) ProtoMessage() {}

func (x *LargestFieldNumber) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LargestFieldNumber.ProtoReflect.Descriptor instead.
func (*LargestFieldNumber) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *LargestFieldNumber) GetInt32() int32 {
	if x != nil {
		return x.Int32
	}
	return 0
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0xf1, 0x0a, 0x0a, 0x07, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x74, 0x38, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x69, 0x6e, 0x74,
	0x38, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x31, 0x36, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x69, 0x6e, 0x74, 0x31, 0x36, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x12, 0x14, 0x0a, 0x05, 0x75, 0x69, 0x6e, 0x74, 0x38, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x75, 0x69, 0x6e, 0x74, 0x38, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e,
	0x74, 0x31, 0x36, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x31,
	0x36, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x69, 0x6e, 0x74, 0x38, 0x18, 0x09, 0x20, 0x01, 0x28, 0x11,
	0x52, 0x05, 0x73, 0x69, 0x6e, 0x74, 0x38, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x31,
	0x36, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x11, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x31, 0x36, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x11, 0x52,
	0x06, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x12, 0x52, 0x06, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x07,
	0x52, 0x07, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x36, 0x34, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x10, 0x52, 0x08, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x6f, 0x6c, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x13, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x12, 0x48, 0x0a,
	0x14, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x12, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x38, 0x18, 0x15, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0c,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x38, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x31, 0x36, 0x18, 0x16,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e,
	0x74, 0x31, 0x36, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x17, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0d, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x18, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x69,
	0x6e, 0x74, 0x38, 0x18, 0x19, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x38, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x31, 0x36, 0x18, 0x1a, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x31,
	0x36, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x18, 0x1b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x1c, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x55, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x73, 0x69, 0x6e, 0x74, 0x38, 0x18, 0x1d, 0x20, 0x03, 0x28, 0x11, 0x52, 0x0d, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x74, 0x38, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x31, 0x36, 0x18, 0x1e, 0x20,
	0x03, 0x28, 0x11, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e,
	0x74, 0x31, 0x36, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x1f, 0x20, 0x03, 0x28, 0x11, 0x52, 0x0e, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x27, 0x0a, 0x0f,
	0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18,
	0x20, 0x20, 0x03, 0x28, 0x12, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x18, 0x21, 0x20, 0x03, 0x28, 0x07, 0x52,
	0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x78,
	0x65, 0x64, 0x36, 0x34, 0x18, 0x22, 0x20, 0x03, 0x28, 0x06, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x2b, 0x0a, 0x11, 0x72,
	0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x18, 0x23, 0x20, 0x03, 0x28, 0x0f, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x2b, 0x0a, 0x11, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x18, 0x24, 0x20,
	0x03, 0x28, 0x10, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x25, 0x20, 0x03, 0x28, 0x08, 0x52, 0x0c, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x26, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x27, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0d, 0x72, 0x65, 0x70,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x1d, 0x72, 0x65,
	0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x28, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x73, 0x74, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x1a, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x73, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x12, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x73, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x05, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x18, 0xff, 0xff, 0xff, 0xff, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x74, 0x65, 0x70, 0x68, 0x65, 0x6e, 0x42, 0x75, 0x74, 0x74, 0x6f,
	0x6c, 0x70, 0x68, 0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x6e, 0x6f, 0x74,
	0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_types_proto_goTypes = []interface{}{
	(*Scalars)(nil),            // 0: pb.Scalars
	(*LargestFieldNumber)(nil), // 1: pb.LargestFieldNumber
}
var file_types_proto_depIdxs = []int32{
	1, // 0: pb.Scalars.largest_field_number:type_name -> pb.LargestFieldNumber
	1, // 1: pb.Scalars.repeated_largest_field_number:type_name -> pb.LargestFieldNumber
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Scalars); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LargestFieldNumber); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
