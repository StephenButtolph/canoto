package testdata

type leadingDigitOneOfName struct {
	A int64 `canoto:"int,1,1OneOf"`
	B int64 `canoto:"int,2,1OneOf"`

	canotoData canotoData_leadingDigitOneOfName
}
