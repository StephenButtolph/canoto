//go:generate go run github.com/StephenButtolph/canoto/canoto $GOFILE

package examples

import "github.com/StephenButtolph/canoto"

var _ canoto.Message = (*genericBool[bool])(nil)

// genericBool ensures that the generated zero checks compile for bool fields
// whose type is a type parameter, which can't be used as a condition.
type genericBool[B ~bool] struct {
	Flag  B `canoto:"bool,1"`
	OneOf B `canoto:"bool,2,A"`
	Other B `canoto:"bool,3,A"`

	canotoData canotoData_genericBool
}
