//go:generate canoto $GOFILE

package examples

type OneOf struct {
	Int8                            int8                           `canoto:"int,1,a"`
	Int16                           int16                          `canoto:"int,2,a"`
	Int32                           int32                          `canoto:"int,3,a"`
	Int64                           int64                          `canoto:"int,4,b"`
	Uint8                           uint8                          `canoto:"int,5,b"`
	Uint16                          uint16                         `canoto:"int,6,b"`
	Uint32                          uint32                         `canoto:"int,7,c"`
	Uint64                          uint64                         `canoto:"int,8,c"`
	Sint8                           int8                           `canoto:"sint,9,c"`
	Sint16                          int16                          `canoto:"sint,10,d"`
	Sint32                          int32                          `canoto:"sint,11,d"`
	Sint64                          int64                          `canoto:"sint,12,d"`
	Fixed32                         uint32                         `canoto:"fint32,13,e"`
	Fixed64                         uint64                         `canoto:"fint64,14,e"`
	Sfixed32                        int32                          `canoto:"fint32,15,e"`
	Sfixed64                        int64                          `canoto:"fint64,16,f"`
	Bool                            bool                           `canoto:"bool,17,f"`
	String                          string                         `canoto:"string,18,f"`
	Bytes                           []byte                         `canoto:"bytes,19,f"`
	LargestFieldNumber              LargestFieldNumber[int32]      `canoto:"field,20,f"`
	RepeatedInt8                    []int8                         `canoto:"repeated int,21,g"`
	RepeatedInt16                   []int16                        `canoto:"repeated int,22,g"`
	RepeatedInt32                   []int32                        `canoto:"repeated int,23,g"`
	RepeatedInt64                   []int64                        `canoto:"repeated int,24,g"`
	RepeatedUint8                   []uint8                        `canoto:"repeated int,25,g"`
	RepeatedUint16                  []uint16                       `canoto:"repeated int,26,h"`
	RepeatedUint32                  []uint32                       `canoto:"repeated int,27,h"`
	RepeatedUint64                  []uint64                       `canoto:"repeated int,28,h"`
	RepeatedSint8                   []int8                         `canoto:"repeated sint,29,h"`
	RepeatedSint16                  []int16                        `canoto:"repeated sint,30,h"`
	RepeatedSint32                  []int32                        `canoto:"repeated sint,31,h"`
	RepeatedSint64                  []int64                        `canoto:"repeated sint,32,i"`
	RepeatedFixed32                 []uint32                       `canoto:"repeated fint32,33,i"`
	RepeatedFixed64                 []uint64                       `canoto:"repeated fint64,34,i"`
	RepeatedSfixed32                []int32                        `canoto:"repeated fint32,35,i"`
	RepeatedSfixed64                []int64                        `canoto:"repeated fint64,36,i"`
	RepeatedBool                    []bool                         `canoto:"repeated bool,37,j"`
	RepeatedString                  []string                       `canoto:"repeated string,38,j"`
	RepeatedBytes                   [][]byte                       `canoto:"repeated bytes,39,j"`
	RepeatedLargestFieldNumber      []LargestFieldNumber[int32]    `canoto:"repeated field,40,j"`
	FixedRepeatedInt8               [3]int8                        `canoto:"fixed repeated int,41,k"`
	FixedRepeatedInt16              [3]int16                       `canoto:"fixed repeated int,42,k"`
	FixedRepeatedInt32              [3]int32                       `canoto:"fixed repeated int,43,k"`
	FixedRepeatedInt64              [3]int64                       `canoto:"fixed repeated int,44,k"`
	FixedRepeatedUint8              [3]uint8                       `canoto:"fixed repeated int,45,k"`
	FixedRepeatedUint16             [3]uint16                      `canoto:"fixed repeated int,46,k"`
	FixedRepeatedUint32             [3]uint32                      `canoto:"fixed repeated int,47,k"`
	FixedRepeatedUint64             [3]uint64                      `canoto:"fixed repeated int,48,l"`
	FixedRepeatedSint8              [3]int8                        `canoto:"fixed repeated sint,49,l"`
	FixedRepeatedSint16             [3]int16                       `canoto:"fixed repeated sint,50,l"`
	FixedRepeatedSint32             [3]int32                       `canoto:"fixed repeated sint,51,l"`
	FixedRepeatedSint64             [3]int64                       `canoto:"fixed repeated sint,52,l"`
	FixedRepeatedFixed32            [3]uint32                      `canoto:"fixed repeated fint32,53,m"`
	FixedRepeatedFixed64            [3]uint64                      `canoto:"fixed repeated fint64,54,m"`
	FixedRepeatedSfixed32           [3]int32                       `canoto:"fixed repeated fint32,55,m"`
	FixedRepeatedSfixed64           [3]int64                       `canoto:"fixed repeated fint64,56,m"`
	FixedRepeatedBool               [3]bool                        `canoto:"fixed repeated bool,57,n"`
	FixedRepeatedString             [3]string                      `canoto:"fixed repeated string,58,n"`
	FixedBytes                      [32]byte                       `canoto:"fixed bytes,59,n"`
	RepeatedFixedBytes              [][32]byte                     `canoto:"repeated fixed bytes,60,o"`
	FixedRepeatedBytes              [3][]byte                      `canoto:"fixed repeated bytes,61,o"`
	FixedRepeatedFixedBytes         [3][32]byte                    `canoto:"fixed repeated fixed bytes,62,o"`
	FixedRepeatedLargestFieldNumber [3]LargestFieldNumber[int32]   `canoto:"fixed repeated field,63,o"`
	ConstRepeatedUint64             [constRepeatedUint64Len]uint64 `canoto:"fixed repeated int,64,o"`
	CustomType                      CustomType                     `canoto:"field,65,p"`
	CustomUint32                    customUint32                   `canoto:"fint32,66,p"`
	CustomString                    customString                   `canoto:"string,67,p"`
	CustomBytes                     customBytes                    `canoto:"bytes,68,p"`
	CustomFixedBytes                customFixedBytes               `canoto:"fixed bytes,69,p"`
	CustomRepeatedBytes             customRepeatedBytes            `canoto:"repeated bytes,70,p"`
	CustomRepeatedFixedBytes        customRepeatedFixedBytes       `canoto:"repeated fixed bytes,71,q"`
	CustomFixedRepeatedBytes        customFixedRepeatedBytes       `canoto:"fixed repeated bytes,72,q"`
	CustomFixedRepeatedFixedBytes   customFixedRepeatedFixedBytes  `canoto:"fixed repeated fixed bytes,73,q"`

	canotoData canotoData_OneOf
}
