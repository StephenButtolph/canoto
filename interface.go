package canoto

// Message defines a type that can be a stand-alone Canoto message.
type Message interface {
	Field
	// MarshalCanoto returns the Canoto representation of this message.
	//
	// It is assumed that this message is ValidCanoto.
	MarshalCanoto() []byte
	// UnmarshalCanoto unmarshals a Canoto-encoded byte slice into the message.
	UnmarshalCanoto(bytes []byte) error
}

// Field defines a type that can be included inside of a Canoto message.
type Field interface {
	// MarshalCanotoInto writes the field into a canoto.Writer and returns the
	// resulting canoto.Writer.
	//
	// It is assumed that CalculateCanotoCache has been called since the last
	// modification to this field.
	//
	// It is assumed that this field is ValidCanoto.
	MarshalCanotoInto(w Writer) Writer
	// CalculateCanotoCache populates internal caches based on the current
	// values in the struct.
	CalculateCanotoCache()
	// CachedCanotoSize returns the previously calculated size of the Canoto
	// representation from CalculateCanotoCache.
	//
	// If CalculateCanotoCache has not yet been called, or the field has been
	// modified since the last call to CalculateCanotoCache, the returned size
	// may be incorrect.
	CachedCanotoSize() int
	// UnmarshalCanotoFrom populates the field from a canoto.Reader.
	UnmarshalCanotoFrom(r Reader) error
	// ValidCanoto validates that the field can be correctly marshaled into the
	// Canoto format.
	ValidCanoto() bool
}

// FieldPointer is a pointer to a concrete Field value T.
//
// This type must be used when implementing a value for a generic Field.
type FieldPointer[T any] interface {
	Field
	*T
}

// FieldMaker is a Field that can create a new value of type T.
//
// The returned value must be able to be unmarshaled into.
//
// This type can be used when implementing a generic Field. However, if T is an
// interface, it is possible for generated code to compile and panic at runtime.
type FieldMaker[T any] interface {
	Field
	MakeCanoto() T
}
