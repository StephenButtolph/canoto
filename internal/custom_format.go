//go:generate go run github.com/StephenButtolph/canoto/canoto --format-cache={struct}Cache --format-number={struct}{field}Number --format-tag={struct}{field}Tag $GOFILE

package examples

import "github.com/StephenButtolph/canoto"

const (
	_ = CustomFormatUintNumber
	_ = CustomFormatUintTag
)

var _ canoto.Message = (*CustomFormat)(nil)

type CustomFormat struct {
	Uint uint64 `canoto:"uint,1"`

	canotoData CustomFormatCache
}
