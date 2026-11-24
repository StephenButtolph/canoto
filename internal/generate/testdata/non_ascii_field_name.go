package testdata

type nonASCIIFieldName struct {
	Größe int64 `canoto:"int,1"`

	canotoData canotoData_nonASCIIFieldName
}
