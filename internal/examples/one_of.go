//go:generate canoto --proto $GOFILE

package examples

import "github.com/StephenButtolph/canoto"

var _ canoto.Message = (*OneOf)(nil)

type OneOf struct {
	A1 int32 `canoto:"int,1,a"`
	A2 int64 `canoto:"int,7,a"`
	B1 int32 `canoto:"int,3,b"`
	B2 int64 `canoto:"int,4,b"`
	C  int32 `canoto:"int,5"`
	D  int64 `canoto:"int,6"`

	canotoData canotoData_OneOf
}
