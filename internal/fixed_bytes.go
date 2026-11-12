//go:generate go run github.com/StephenButtolph/canoto/canoto $GOFILE

package examples

import (
	"github.com/StephenButtolph/canoto"
	"github.com/StephenButtolph/canoto/internal/imported"
)

var _ canoto.Message = (*fixedBytes)(nil)

// fixedBytes ensures that the generated zero checks compile for fixed bytes
// fields whose types are imported, and for fixed bytes fields in a oneOf.
type fixedBytes struct {
	Imported imported.FixedBytes `canoto:"fixed bytes,1"`
	OneOfA   [4]byte             `canoto:"fixed bytes,2,A"`
	OneOfB   customFixedBytes    `canoto:"fixed bytes,3,A"`

	canotoData canotoData_fixedBytes
}
