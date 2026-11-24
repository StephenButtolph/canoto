package testdata

type pointerCanotoData struct {
	Int int64 `canoto:"int,1"`

	canotoData *canotoData_pointerCanotoData
}
