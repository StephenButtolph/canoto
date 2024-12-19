//go:generate canoto --proto $GOFILE

package examples

import (
	"github.com/StephenButtolph/canoto"
	"github.com/StephenButtolph/canoto/internal/big"
)

const constRepeatedUint64Len = 3

var (
	_ canoto.Message = (*LargestFieldNumber[int32])(nil)
	_ canoto.Message = (*OneOf)(nil)
	_ canoto.Message = (*Scalars)(nil)
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

type LargestFieldNumber[T canoto.Int] struct {
	Int32 T `canoto:"int,536870911"`

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

type Scalars struct {
	Int8                            int8                           `canoto:"int,1"`
	Int16                           int16                          `canoto:"int,2"`
	Int32                           int32                          `canoto:"int,3"`
	Int64                           int64                          `canoto:"int,4"`
	Uint8                           uint8                          `canoto:"int,5"`
	Uint16                          uint16                         `canoto:"int,6"`
	Uint32                          uint32                         `canoto:"int,7"`
	Uint64                          uint64                         `canoto:"int,8"`
	Sint8                           int8                           `canoto:"sint,9"`
	Sint16                          int16                          `canoto:"sint,10"`
	Sint32                          int32                          `canoto:"sint,11"`
	Sint64                          int64                          `canoto:"sint,12"`
	Fixed32                         uint32                         `canoto:"fint32,13"`
	Fixed64                         uint64                         `canoto:"fint64,14"`
	Sfixed32                        int32                          `canoto:"fint32,15"`
	Sfixed64                        int64                          `canoto:"fint64,16"`
	Bool                            bool                           `canoto:"bool,17"`
	String                          string                         `canoto:"string,18"`
	Bytes                           []byte                         `canoto:"bytes,19"`
	LargestFieldNumber              LargestFieldNumber[int32]      `canoto:"field,20"`
	RepeatedInt8                    []int8                         `canoto:"repeated int,21"`
	RepeatedInt16                   []int16                        `canoto:"repeated int,22"`
	RepeatedInt32                   []int32                        `canoto:"repeated int,23"`
	RepeatedInt64                   []int64                        `canoto:"repeated int,24"`
	RepeatedUint8                   []uint8                        `canoto:"repeated int,25"`
	RepeatedUint16                  []uint16                       `canoto:"repeated int,26"`
	RepeatedUint32                  []uint32                       `canoto:"repeated int,27"`
	RepeatedUint64                  []uint64                       `canoto:"repeated int,28"`
	RepeatedSint8                   []int8                         `canoto:"repeated sint,29"`
	RepeatedSint16                  []int16                        `canoto:"repeated sint,30"`
	RepeatedSint32                  []int32                        `canoto:"repeated sint,31"`
	RepeatedSint64                  []int64                        `canoto:"repeated sint,32"`
	RepeatedFixed32                 []uint32                       `canoto:"repeated fint32,33"`
	RepeatedFixed64                 []uint64                       `canoto:"repeated fint64,34"`
	RepeatedSfixed32                []int32                        `canoto:"repeated fint32,35"`
	RepeatedSfixed64                []int64                        `canoto:"repeated fint64,36"`
	RepeatedBool                    []bool                         `canoto:"repeated bool,37"`
	RepeatedString                  []string                       `canoto:"repeated string,38"`
	RepeatedBytes                   [][]byte                       `canoto:"repeated bytes,39"`
	RepeatedLargestFieldNumber      []LargestFieldNumber[int32]    `canoto:"repeated field,40"`
	FixedRepeatedInt8               [3]int8                        `canoto:"fixed repeated int,41"`
	FixedRepeatedInt16              [3]int16                       `canoto:"fixed repeated int,42"`
	FixedRepeatedInt32              [3]int32                       `canoto:"fixed repeated int,43"`
	FixedRepeatedInt64              [3]int64                       `canoto:"fixed repeated int,44"`
	FixedRepeatedUint8              [3]uint8                       `canoto:"fixed repeated int,45"`
	FixedRepeatedUint16             [3]uint16                      `canoto:"fixed repeated int,46"`
	FixedRepeatedUint32             [3]uint32                      `canoto:"fixed repeated int,47"`
	FixedRepeatedUint64             [3]uint64                      `canoto:"fixed repeated int,48"`
	FixedRepeatedSint8              [3]int8                        `canoto:"fixed repeated sint,49"`
	FixedRepeatedSint16             [3]int16                       `canoto:"fixed repeated sint,50"`
	FixedRepeatedSint32             [3]int32                       `canoto:"fixed repeated sint,51"`
	FixedRepeatedSint64             [3]int64                       `canoto:"fixed repeated sint,52"`
	FixedRepeatedFixed32            [3]uint32                      `canoto:"fixed repeated fint32,53"`
	FixedRepeatedFixed64            [3]uint64                      `canoto:"fixed repeated fint64,54"`
	FixedRepeatedSfixed32           [3]int32                       `canoto:"fixed repeated fint32,55"`
	FixedRepeatedSfixed64           [3]int64                       `canoto:"fixed repeated fint64,56"`
	FixedRepeatedBool               [3]bool                        `canoto:"fixed repeated bool,57"`
	FixedRepeatedString             [3]string                      `canoto:"fixed repeated string,58"`
	FixedBytes                      [32]byte                       `canoto:"fixed bytes,59"`
	RepeatedFixedBytes              [][32]byte                     `canoto:"repeated fixed bytes,60"`
	FixedRepeatedBytes              [3][]byte                      `canoto:"fixed repeated bytes,61"`
	FixedRepeatedFixedBytes         [3][32]byte                    `canoto:"fixed repeated fixed bytes,62"`
	FixedRepeatedLargestFieldNumber [3]LargestFieldNumber[int32]   `canoto:"fixed repeated field,63"`
	ConstRepeatedUint64             [constRepeatedUint64Len]uint64 `canoto:"fixed repeated int,64"`
	CustomType                      big.Int                        `canoto:"field,65"`
	CustomUint32                    customUint32                   `canoto:"fint32,66"`
	CustomString                    customString                   `canoto:"string,67"`
	CustomBytes                     customBytes                    `canoto:"bytes,68"`
	CustomFixedBytes                customFixedBytes               `canoto:"fixed bytes,69"`
	CustomRepeatedBytes             customRepeatedBytes            `canoto:"repeated bytes,70"`
	CustomRepeatedFixedBytes        customRepeatedFixedBytes       `canoto:"repeated fixed bytes,71"`
	CustomFixedRepeatedBytes        customFixedRepeatedBytes       `canoto:"fixed repeated bytes,72"`
	CustomFixedRepeatedFixedBytes   customFixedRepeatedFixedBytes  `canoto:"fixed repeated fixed bytes,73"`
	OneOf                           OneOf                          `canoto:"field,74"`

	canotoData canotoData_Scalars
}
