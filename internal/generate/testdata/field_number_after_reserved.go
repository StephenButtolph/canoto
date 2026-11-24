package testdata

type fieldNumberAfterReserved struct {
	Int int64 `canoto:"int,20000"`

	canotoData canotoData_fieldNumberAfterReserved
}
