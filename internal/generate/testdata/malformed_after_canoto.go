package testdata

type malformedAfterCanoto struct {
	Int int64 `canoto:"int,1" json:int`

	canotoData canotoData_malformedAfterCanoto
}
