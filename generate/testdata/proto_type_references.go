package testdata

import (
	"example.com/other"

	"github.com/StephenButtolph/canoto"
)

type Foo struct {
	A int64 `canoto:"int,1"`

	canotoData canotoData_Foo
}

// double is also the name of a proto scalar type.
type double struct {
	A int64 `canoto:"int,1"`

	canotoData canotoData_double
}

// message is also a proto keyword.
type message struct {
	A int64 `canoto:"int,1"`

	canotoData canotoData_message
}

type Outer struct {
	Local    Foo       `canoto:"value,1"`
	Imported other.Foo `canoto:"value,2"`
	Double   double    `canoto:"value,3"`
	Message  message   `canoto:"value,4"`
	Pointers []*Foo    `canoto:"repeated pointer,5"`

	canotoData canotoData_Outer
}

// Generic has a type parameter with the same name as a local message.
type Generic[Foo any, _ canoto.FieldPointer[Foo]] struct {
	Param    Foo       `canoto:"value,1"`
	Imported other.Foo `canoto:"value,2"`

	canotoData canotoData_Generic
}
