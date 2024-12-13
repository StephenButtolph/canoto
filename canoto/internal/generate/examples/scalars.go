//go:generate canoto $GOFILE

package examples

const constRepeatedUint64Len = 3

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
	Fixed32                         uint32                         `canoto:"fint,13"`
	Fixed64                         uint64                         `canoto:"fint,14"`
	Sfixed32                        int32                          `canoto:"fint,15"`
	Sfixed64                        int64                          `canoto:"fint,16"`
	Bool                            bool                           `canoto:"bool,17"`
	String                          string                         `canoto:"bytes,18"`
	Bytes                           []byte                         `canoto:"bytes,19"`
	LargestFieldNumber              LargestFieldNumber             `canoto:"bytes,20"`
	RepeatedInt8                    []int8                         `canoto:"int,21"`
	RepeatedInt16                   []int16                        `canoto:"int,22"`
	RepeatedInt32                   []int32                        `canoto:"int,23"`
	RepeatedInt64                   []int64                        `canoto:"int,24"`
	RepeatedUint8                   []uint8                        `canoto:"int,25"`
	RepeatedUint16                  []uint16                       `canoto:"int,26"`
	RepeatedUint32                  []uint32                       `canoto:"int,27"`
	RepeatedUint64                  []uint64                       `canoto:"int,28"`
	RepeatedSint8                   []int8                         `canoto:"sint,29"`
	RepeatedSint16                  []int16                        `canoto:"sint,30"`
	RepeatedSint32                  []int32                        `canoto:"sint,31"`
	RepeatedSint64                  []int64                        `canoto:"sint,32"`
	RepeatedFixed32                 []uint32                       `canoto:"fint,33"`
	RepeatedFixed64                 []uint64                       `canoto:"fint,34"`
	RepeatedSfixed32                []int32                        `canoto:"fint,35"`
	RepeatedSfixed64                []int64                        `canoto:"fint,36"`
	RepeatedBool                    []bool                         `canoto:"bool,37"`
	RepeatedString                  []string                       `canoto:"bytes,38"`
	RepeatedBytes                   [][]byte                       `canoto:"bytes,39"`
	RepeatedLargestFieldNumber      []LargestFieldNumber           `canoto:"bytes,40"`
	FixedRepeatedInt8               [3]int8                        `canoto:"int,41"`
	FixedRepeatedInt16              [3]int16                       `canoto:"int,42"`
	FixedRepeatedInt32              [3]int32                       `canoto:"int,43"`
	FixedRepeatedInt64              [3]int64                       `canoto:"int,44"`
	FixedRepeatedUint8              [3]uint8                       `canoto:"int,45"`
	FixedRepeatedUint16             [3]uint16                      `canoto:"int,46"`
	FixedRepeatedUint32             [3]uint32                      `canoto:"int,47"`
	FixedRepeatedUint64             [3]uint64                      `canoto:"int,48"`
	FixedRepeatedSint8              [3]int8                        `canoto:"sint,49"`
	FixedRepeatedSint16             [3]int16                       `canoto:"sint,50"`
	FixedRepeatedSint32             [3]int32                       `canoto:"sint,51"`
	FixedRepeatedSint64             [3]int64                       `canoto:"sint,52"`
	FixedRepeatedFixed32            [3]uint32                      `canoto:"fint,53"`
	FixedRepeatedFixed64            [3]uint64                      `canoto:"fint,54"`
	FixedRepeatedSfixed32           [3]int32                       `canoto:"fint,55"`
	FixedRepeatedSfixed64           [3]int64                       `canoto:"fint,56"`
	FixedRepeatedBool               [3]bool                        `canoto:"bool,57"`
	FixedRepeatedString             [3]string                      `canoto:"bytes,58"`
	FixedBytes                      [32]byte                       `canoto:"bytes,59"`
	RepeatedFixedBytes              [][32]byte                     `canoto:"bytes,60"`
	FixedRepeatedBytes              [3][]byte                      `canoto:"bytes,61"`
	FixedRepeatedFixedBytes         [3][32]byte                    `canoto:"bytes,62"`
	FixedRepeatedLargestFieldNumber [3]LargestFieldNumber          `canoto:"bytes,63"`
	ConstRepeatedUint64             [constRepeatedUint64Len]uint64 `canoto:"int,64"`

	canotoData canotoData_Scalars
}
