//go:generate go run github.com/StephenButtolph/canoto/canoto --format-cache={struct}Cache --format-number={struct}{field}Number --format-tag={struct}{field}Tag --format-oneof={struct}{oneOf} $GOFILE

package examples

import "github.com/StephenButtolph/canoto"

const (
	_ = CustomFormatUintNumber
	_ = CustomFormatUintTag
)

var _ CustomFormatA

var _ canoto.Message = (*CustomFormat)(nil)

type CustomFormat struct {
	Uint uint64 `canoto:"uint,1"`

	A1 int32 `canoto:"int,2,A"`
	A2 int64 `canoto:"int,3,A"`

	canotoData CustomFormatCache
}
