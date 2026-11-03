package testdata

type fieldNumberBeforeReserved struct {
	Int int64 `canoto:"int,18999"`
}
