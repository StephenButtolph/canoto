//go:generate go run github.com/StephenButtolph/canoto/canoto $GOFILE

// Package imported declares messages whose names collide with identifiers in
// the examples package, to verify that imported types are not confused with
// local messages or type parameters.
package imported

import "github.com/StephenButtolph/canoto"

var (
	_ canoto.Message = (*A)(nil)
	_ canoto.Message = (*V)(nil)
)

// A collides with the examples.A message.
type A struct {
	Int int64 `canoto:"int,1"`

	canotoData canotoData_A
}

// V collides with the FieldPointer type parameter of examples.ImportedName.
type V struct {
	Int int64 `canoto:"int,1"`

	canotoData canotoData_V
}

// FixedBytes is a fixed bytes type declared in another package.
type FixedBytes [4]byte
