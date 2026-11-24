package testdata

type nonASCIIOneOfName struct {
	A int64 `canoto:"int,1,Größe"`
	B int64 `canoto:"int,2,Größe"`

	canotoData canotoData_nonASCIIOneOfName
}
