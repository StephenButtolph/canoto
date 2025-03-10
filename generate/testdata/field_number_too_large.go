package testdata

type fieldNumberTooLarge struct {
	Int int64 `canoto:"sint,536870912"`
}
