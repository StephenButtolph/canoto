package testdata

type fieldNumberTooLarge struct {
	Int int64 `canoto:"int,536870912"`
}
