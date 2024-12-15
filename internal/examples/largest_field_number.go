//go:generate canoto --proto=true $GOFILE

package examples

import "github.com/StephenButtolph/canoto"

type LargestFieldNumber[T canoto.Int] struct {
	Int32 T `canoto:"int,536870911"`

	canotoData canotoData_LargestFieldNumber
}
