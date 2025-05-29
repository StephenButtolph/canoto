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
	Int8 int8 `canoto:"int,1"`

	canotoData canotoData_justAnInt `canoto:"nocopy"`
}
