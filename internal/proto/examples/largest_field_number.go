package examples

type LargestFieldNumber struct {
	Int32 int32 `canoto:"int,536870911"`

	canotoData canotoData_LargestFieldNumber //nolint // needed for codegen
}
