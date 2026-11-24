package testdata

type quotedStructTag struct {
	Int int64 "canoto:\"int,0\""

	canotoData canotoData_quotedStructTag
}
