//go:generate canoto $GOFILE

package examples

type Scalars struct {
	Int8                       int8                 `canoto:"int,1"`
	Int16                      int16                `canoto:"int,2"`
	Int32                      int32                `canoto:"int,3"`
	Int64                      int64                `canoto:"int,4"`
	Uint8                      uint8                `canoto:"int,5"`
	Uint16                     uint16               `canoto:"int,6"`
	Uint32                     uint32               `canoto:"int,7"`
	Uint64                     uint64               `canoto:"int,8"`
	Sint8                      int8                 `canoto:"sint,9"`
	Sint16                     int16                `canoto:"sint,10"`
	Sint32                     int32                `canoto:"sint,11"`
	Sint64                     int64                `canoto:"sint,12"`
	Fixed32                    uint32               `canoto:"fint,13"`
	Fixed64                    uint64               `canoto:"fint,14"`
	Sfixed32                   int32                `canoto:"fint,15"`
	Sfixed64                   int64                `canoto:"fint,16"`
	Bool                       bool                 `canoto:"bool,17"`
	String                     string               `canoto:"bytes,18"`
	Bytes                      []byte               `canoto:"bytes,19"`
	LargestFieldNumber         LargestFieldNumber   `canoto:"bytes,20"`
	RepeatedInt8               []int8               `canoto:"int,21"`
	RepeatedInt16              []int16              `canoto:"int,22"`
	RepeatedInt32              []int32              `canoto:"int,23"`
	RepeatedInt64              []int64              `canoto:"int,24"`
	RepeatedUint8              []uint8              `canoto:"int,25"`
	RepeatedUint16             []uint16             `canoto:"int,26"`
	RepeatedUint32             []uint32             `canoto:"int,27"`
	RepeatedUint64             []uint64             `canoto:"int,28"`
	RepeatedSint8              []int8               `canoto:"sint,29"`
	RepeatedSint16             []int16              `canoto:"sint,30"`
	RepeatedSint32             []int32              `canoto:"sint,31"`
	RepeatedSint64             []int64              `canoto:"sint,32"`
	RepeatedFixed32            []uint32             `canoto:"fint,33"`
	RepeatedFixed64            []uint64             `canoto:"fint,34"`
	RepeatedSfixed32           []int32              `canoto:"fint,35"`
	RepeatedSfixed64           []int64              `canoto:"fint,36"`
	RepeatedBool               []bool               `canoto:"bool,37"`
	RepeatedString             []string             `canoto:"bytes,38"`
	RepeatedBytes              [][]byte             `canoto:"bytes,39"`
	RepeatedLargestFieldNumber []LargestFieldNumber `canoto:"bytes,40"`

	canotoData canotoData_Scalars
}
