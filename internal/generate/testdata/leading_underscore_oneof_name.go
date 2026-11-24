package testdata

type leadingUnderscoreOneOfName struct {
	A int64 `canoto:"int,1,_OneOf"`
	B int64 `canoto:"int,2,_OneOf"`

	canotoData canotoData_leadingUnderscoreOneOfName
}
