package testdata

type duplicateFieldNumber struct {
	IntA int64 `canoto:"sint,1"`
	IntB int64 `canoto:"sint,1"`
}
