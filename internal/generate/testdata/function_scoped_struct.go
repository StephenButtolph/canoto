package testdata

func functionScopedStruct() {
	type functionScopedStruct struct {
		Int int64 `canoto:"int,1"`
	}
	_ = functionScopedStruct{}
}

var _ = func() {
	type functionLiteralScopedStruct struct {
		Int int64 `canoto:"int,1"`
	}
	_ = functionLiteralScopedStruct{}
}
