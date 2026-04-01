//go:generate go run github.com/StephenButtolph/canoto/canoto --format-cache={struct}Cache --format-number={struct}{field}Number --format-tag={struct}{field}Tag --format-oneof-type={struct}{oneOf} --format-oneof-unset={struct}{oneOf}Unset --format-oneof-field={struct}{oneOf}{field} $GOFILE

package examples

import "github.com/StephenButtolph/canoto"

var _ canoto.Message = (*CustomFormat)(nil)

type CustomFormat struct {
	A uint64 `canoto:"uint,1,Fields"`
	B uint64 `canoto:"uint,2,Fields"`

	canotoData CustomFormatCache
}
