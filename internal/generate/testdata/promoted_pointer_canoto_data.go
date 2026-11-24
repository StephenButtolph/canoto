package testdata

// promotedPointerCanotoData does not declare canotoData, but one is promoted
// from the embedded *promotedPointerCanotoDataBase.
type promotedPointerCanotoData struct {
	*promotedPointerCanotoDataBase

	Other int64 `canoto:"int,1"`
}

type promotedPointerCanotoDataBase struct {
	Int int64 `canoto:"int,1"`

	canotoData canotoData_promotedPointerCanotoDataBase
}
