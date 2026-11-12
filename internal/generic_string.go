//go:generate go run github.com/StephenButtolph/canoto/canoto $GOFILE

package examples

import "github.com/StephenButtolph/canoto"

var _ canoto.Message = (*genericString[string])(nil)

// genericString ensures that the generated zero checks compile for string
// fields in a oneOf whose type is a type parameter.
type genericString[S ~string] struct {
	OneOf S `canoto:"string,1,A"`
	Other S `canoto:"string,2,A"`

	canotoData canotoData_genericString
}
