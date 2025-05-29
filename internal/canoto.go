//go:generate canoto --proto $GOFILE

package examples

import (
	"github.com/StephenButtolph/canoto"
	"github.com/StephenButtolph/canoto/internal/big"
)

const constRepeatedUint64Len = 3

var (
	_ canoto.Message = (*LargestFieldNumber[uint32])(nil)
	_ canoto.Message = (*OneOf)(nil)
	_ canoto.Message = (*GenericField[OneOf, *OneOf, *OneOf])(nil)
	_ canoto.Message = (*Scalars)(nil)

	_ canoto.FieldMaker[*LargestFieldNumber[uint32]]          = (*LargestFieldNumber[uint32])(nil)
	_ canoto.FieldMaker[*OneOf]                               = (*OneOf)(nil)
	_ canoto.FieldMaker[*GenericField[OneOf, *OneOf, *OneOf]] = (*GenericField[OneOf, *OneOf, *OneOf])(nil)
	_ canoto.FieldMaker[*Scalars]                             = (*Scalars)(nil)
)

type (
	customUint32                  uint32
	customString                  string
	customBytes                   []byte
	customFixedBytes              [3]byte
	customRepeatedBytes           [][]byte
	customRepeatedFixedBytes      [][32]byte
	customFixedRepeatedBytes      [3][]byte
	customFixedRepeatedFixedBytes [3][32]byte
)

type LargestFieldNumber[T canoto.Uint] struct {
	Uint T `canoto:"uint,536870911"`

	canotoData canotoData_LargestFieldNumber
}

type OneOf struct {
	A1 int32 `canoto:"int,1,A"`
	A2 int64 `canoto:"int,7,A"`
	B1 int32 `canoto:"int,3,B"`
	B2 int64 `canoto:"int,4,B"`
	C  int32 `canoto:"int,5"`
	D  int64 `canoto:"int,6"`

	canotoData canotoData_OneOf
}

type Node struct {
	Value int32 `canoto:"int,1"`
	Next  *Node `canoto:"pointer,2,OneOf"`

	canotoData canotoData_Node
}

type RecursiveA struct {
	Next *RecursiveB `canoto:"pointer,1"`

	canotoData canotoData_RecursiveA
}

type RecursiveB struct {
	Next *RecursiveA `canoto:"pointer,1"`

	canotoData canotoData_RecursiveB
}

type GenericField[V any, _ canoto.FieldPointer[V], T canoto.FieldMaker[T]] struct {
	Value                V     `canoto:"value,1"`
	RepeatedValue        []V   `canoto:"repeated value,2"`
	FixedRepeatedValue   [3]V  `canoto:"fixed repeated value,3"`
	Pointer              *V    `canoto:"pointer,4"`
	RepeatedPointer      []*V  `canoto:"repeated pointer,5"`
	FixedRepeatedPointer [3]*V `canoto:"fixed repeated pointer,6"`
	Field                T     `canoto:"field,7"`
	RepeatedField        []T   `canoto:"repeated field,8"`
	FixedRepeatedField   [3]T  `canoto:"fixed repeated field,9"`

	canotoData canotoData_GenericField
}

type NestedGenericField[V any, P canoto.FieldPointer[V], T canoto.FieldMaker[T]] struct {
	Value                GenericField[V, P, T]     `canoto:"value,1"`
	RepeatedValue        []GenericField[V, P, T]   `canoto:"repeated value,2"`
	FixedRepeatedValue   [3]GenericField[V, P, T]  `canoto:"fixed repeated value,3"`
	Pointer              *GenericField[V, P, T]    `canoto:"pointer,4"`
	RepeatedPointer      []*GenericField[V, P, T]  `canoto:"repeated pointer,5"`
	FixedRepeatedPointer [3]*GenericField[V, P, T] `canoto:"fixed repeated pointer,6"`
	Field                *GenericField[V, P, T]    `canoto:"field,7"`
	RepeatedField        []*GenericField[V, P, T]  `canoto:"repeated field,8"`
	FixedRepeatedField   [3]*GenericField[V, P, T] `canoto:"fixed repeated field,9"`

	canotoData canotoData_NestedGenericField
}

type Embedded struct {
	OneOf                                `canoto:"value,1"`
	*LargestFieldNumber[uint32]          `canoto:"pointer,2"`
	*GenericField[OneOf, *OneOf, *OneOf] `canoto:"field,3"`

	canotoData canotoData_Embedded
}

// Check for name collisions. Because we use "__" as a separator, the unescaped
// name would conflict as canoto__A__B__C__tag.
//
//nolint:stylecheck // This is checking for name collisions.
type A struct {
	B__C int32 `canoto:"int,1"`

	canotoData canotoData_A
}

//nolint:stylecheck // This is checking for name collisions.
type A__B struct {
	C int32 `canoto:"int,1"`

	canotoData canotoData_A__B
}

type Scalars struct {
	Int8                            int8                           `canoto:"int,1"`
	Int16                           int16                          `canoto:"int,2"`
	Int32                           int32                          `canoto:"int,3"`
	Int64                           int64                          `canoto:"int,4"`
	Uint8                           uint8                          `canoto:"uint,5"`
	Uint16                          uint16                         `canoto:"uint,6"`
	Uint32                          uint32                         `canoto:"uint,7"`
	Uint64                          uint64                         `canoto:"uint,8"`
	Sfixed32                        int32                          `canoto:"fint32,9"`
	Fixed32                         uint32                         `canoto:"fint32,10"`
	Sfixed64                        int64                          `canoto:"fint64,11"`
	Fixed64                         uint64                         `canoto:"fint64,12"`
	Bool                            bool                           `canoto:"bool,13"`
	String                          string                         `canoto:"string,14"`
	Bytes                           []byte                         `canoto:"bytes,15"`
	LargestFieldNumber              LargestFieldNumber[uint32]     `canoto:"value,16"`
	RepeatedInt8                    []int8                         `canoto:"repeated int,17"`
	RepeatedInt16                   []int16                        `canoto:"repeated int,18"`
	RepeatedInt32                   []int32                        `canoto:"repeated int,19"`
	RepeatedInt64                   []int64                        `canoto:"repeated int,20"`
	RepeatedUint8                   []uint8                        `canoto:"repeated uint,21"`
	RepeatedUint16                  []uint16                       `canoto:"repeated uint,22"`
	RepeatedUint32                  []uint32                       `canoto:"repeated uint,23"`
	RepeatedUint64                  []uint64                       `canoto:"repeated uint,24"`
	RepeatedSfixed32                []int32                        `canoto:"repeated fint32,25"`
	RepeatedFixed32                 []uint32                       `canoto:"repeated fint32,26"`
	RepeatedSfixed64                []int64                        `canoto:"repeated fint64,27"`
	RepeatedFixed64                 []uint64                       `canoto:"repeated fint64,28"`
	RepeatedBool                    []bool                         `canoto:"repeated bool,29"`
	RepeatedString                  []string                       `canoto:"repeated string,30"`
	RepeatedBytes                   [][]byte                       `canoto:"repeated bytes,31"`
	RepeatedLargestFieldNumber      []LargestFieldNumber[uint32]   `canoto:"repeated value,32"`
	FixedRepeatedInt8               [3]int8                        `canoto:"fixed repeated int,33"`
	FixedRepeatedInt16              [3]int16                       `canoto:"fixed repeated int,34"`
	FixedRepeatedInt32              [3]int32                       `canoto:"fixed repeated int,35"`
	FixedRepeatedInt64              [3]int64                       `canoto:"fixed repeated int,36"`
	FixedRepeatedUint8              [3]uint8                       `canoto:"fixed repeated uint,37"`
	FixedRepeatedUint16             [3]uint16                      `canoto:"fixed repeated uint,38"`
	FixedRepeatedUint32             [3]uint32                      `canoto:"fixed repeated uint,39"`
	FixedRepeatedUint64             [3]uint64                      `canoto:"fixed repeated uint,40"`
	FixedRepeatedSfixed32           [3]int32                       `canoto:"fixed repeated fint32,41"`
	FixedRepeatedFixed32            [3]uint32                      `canoto:"fixed repeated fint32,42"`
	FixedRepeatedSfixed64           [3]int64                       `canoto:"fixed repeated fint64,43"`
	FixedRepeatedFixed64            [3]uint64                      `canoto:"fixed repeated fint64,44"`
	FixedRepeatedBool               [3]bool                        `canoto:"fixed repeated bool,45"`
	FixedRepeatedString             [3]string                      `canoto:"fixed repeated string,46"`
	FixedBytes                      [32]byte                       `canoto:"fixed bytes,47"`
	RepeatedFixedBytes              [][32]byte                     `canoto:"repeated fixed bytes,48"`
	FixedRepeatedBytes              [3][]byte                      `canoto:"fixed repeated bytes,49"`
	FixedRepeatedFixedBytes         [3][32]byte                    `canoto:"fixed repeated fixed bytes,50"`
	FixedRepeatedLargestFieldNumber [3]LargestFieldNumber[uint32]  `canoto:"fixed repeated value,51"`
	ConstRepeatedUint64             [constRepeatedUint64Len]uint64 `canoto:"fixed repeated uint,52"`
	CustomType                      big.Int                        `canoto:"value,53"`
	CustomUint32                    customUint32                   `canoto:"fint32,54"`
	CustomString                    customString                   `canoto:"string,55"`
	CustomBytes                     customBytes                    `canoto:"bytes,56"`
	CustomFixedBytes                customFixedBytes               `canoto:"fixed bytes,57"`
	CustomRepeatedBytes             customRepeatedBytes            `canoto:"repeated bytes,58"`
	CustomRepeatedFixedBytes        customRepeatedFixedBytes       `canoto:"repeated fixed bytes,59"`
	CustomFixedRepeatedBytes        customFixedRepeatedBytes       `canoto:"fixed repeated bytes,60"`
	CustomFixedRepeatedFixedBytes   customFixedRepeatedFixedBytes  `canoto:"fixed repeated fixed bytes,61"`
	OneOf                           OneOf                          `canoto:"value,62"`
	Pointer                         *LargestFieldNumber[uint32]    `canoto:"pointer,63"`
	RepeatedPointer                 []*LargestFieldNumber[uint32]  `canoto:"repeated pointer,64"`
	FixedRepeatedPointer            [3]*LargestFieldNumber[uint32] `canoto:"fixed repeated pointer,65"`
	Field                           *LargestFieldNumber[uint32]    `canoto:"field,66"`
	RepeatedField                   []*LargestFieldNumber[uint32]  `canoto:"repeated field,67"`
	FixedRepeatedField              [3]*LargestFieldNumber[uint32] `canoto:"fixed repeated field,68"`

	canotoData canotoData_Scalars
}

type SpecUnusedZero struct {
	Bool           bool     `canoto:"bool,1"`
	RepeatedBool   []bool   `canoto:"repeated bool,2"`
	String         string   `canoto:"string,3"`
	RepeatedString []string `canoto:"repeated string,4"`
	Bytes          []byte   `canoto:"bytes,5"`
	RepeatedBytes  [][]byte `canoto:"repeated bytes,6"`

	canotoData canotoData_SpecUnusedZero
}
