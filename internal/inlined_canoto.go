//go:generate canoto --library=./ --import=github.com/StephenButtolph/canoto/internal/canoto $GOFILE

package examples

import (
	"github.com/StephenButtolph/canoto/internal/canoto"
)

var (
	_ canoto.Message                = (*justAnInt)(nil)
	_ canoto.FieldMaker[*justAnInt] = (*justAnInt)(nil)
)

type justAnInt struct {
	Int8 uint8 `canoto:"uint,1"`

	canotoData canotoData_justAnInt `canoto:"noatomic"`
}
