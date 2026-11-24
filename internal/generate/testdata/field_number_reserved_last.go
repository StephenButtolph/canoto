package testdata

type fieldNumberReservedLast struct {
	Int int64 `canoto:"int,19999"`

	canotoData canotoData_fieldNumberReservedLast
}
