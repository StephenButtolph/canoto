//go:generate canoto $GOFILE

package examples

type Scalars struct {
	Int32                      int32                `canoto:"int,1"`
	Int64                      int64                `canoto:"int,2"`
	Uint32                     uint32               `canoto:"int,3"`
	Uint64                     uint64               `canoto:"int,4"`
	Sint32                     int32                `canoto:"sint,5"`
	Sint64                     int64                `canoto:"sint,6"`
	Fixed32                    uint32               `canoto:"fint,7"`
	Fixed64                    uint64               `canoto:"fint,8"`
	Sfixed32                   int32                `canoto:"fint,9"`
	Sfixed64                   int64                `canoto:"fint,10"`
	Bool                       bool                 `canoto:"bool,11"`
	String                     string               `canoto:"bytes,12"`
	Bytes                      []byte               `canoto:"bytes,13"`
	LargestFieldNumber         LargestFieldNumber   `canoto:"bytes,14"`
	RepeatedInt32              []int32              `canoto:"int,15"`
	RepeatedInt64              []int64              `canoto:"int,16"`
	RepeatedUint32             []uint32             `canoto:"int,17"`
	RepeatedUint64             []uint64             `canoto:"int,18"`
	RepeatedSint32             []int32              `canoto:"sint,19"`
	RepeatedSint64             []int64              `canoto:"sint,20"`
	RepeatedFixed32            []uint32             `canoto:"fint,21"`
	RepeatedFixed64            []uint64             `canoto:"fint,22"`
	RepeatedSfixed32           []int32              `canoto:"fint,23"`
	RepeatedSfixed64           []int64              `canoto:"fint,24"`
	RepeatedBool               []bool               `canoto:"bool,25"`
	RepeatedString             []string             `canoto:"bytes,26"`
	RepeatedBytes              [][]byte             `canoto:"bytes,27"`
	RepeatedLargestFieldNumber []LargestFieldNumber `canoto:"bytes,28"`

	canotoData canotoData_Scalars
}
