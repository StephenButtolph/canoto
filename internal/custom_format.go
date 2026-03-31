<<<<<<< HEAD
//go:generate go run github.com/StephenButtolph/canoto/canoto --format-cache={struct}Cache --format-number={struct}{field}Number --format-tag={struct}{field}Tag --format-oneof-type={struct}{oneOf} --format-oneof-unset={struct}{oneOf}Unset --format-oneof-field={struct}{oneOf}{field} $GOFILE
=======
//go:generate go run github.com/StephenButtolph/canoto/canoto --format-cache={struct}Cache --format-number={struct}{field}Number --format-tag={struct}{field}Tag --format-oneof={struct}{oneOf} $GOFILE
>>>>>>> 33dfac36ce3c28de64ec69ad70cbe97d38466697

package examples

import "github.com/StephenButtolph/canoto"

<<<<<<< HEAD
var _ canoto.Message = (*CustomFormat)(nil)

type CustomFormat struct {
	A uint64 `canoto:"uint,1,Fields"`
	B uint64 `canoto:"uint,2,Fields"`
=======
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
>>>>>>> 33dfac36ce3c28de64ec69ad70cbe97d38466697

	canotoData CustomFormatCache
}
