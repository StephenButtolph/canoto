package testdata

// promotedValueCanotoData does not declare canotoData, but one is promoted
// from the embedded promotedValueCanotoDataBase.
type promotedValueCanotoData struct {
	promotedValueCanotoDataBase

	Other int64 `canoto:"int,1"`
}

type promotedValueCanotoDataBase struct {
	Int int64 `canoto:"int,1"`

	canotoData canotoData_promotedValueCanotoDataBase
}
