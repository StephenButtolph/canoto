package testdata

type leadingUnderscoreFieldName struct {
	_Int int64 `canoto:"int,1"`

	canotoData canotoData_leadingUnderscoreFieldName
}
