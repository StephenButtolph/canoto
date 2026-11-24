package testdata

type duplicateFieldNumber struct {
	IntA int64 `canoto:"int,1"`
	IntB int64 `canoto:"int,1"`

	canotoData canotoData_duplicateFieldNumber
}
