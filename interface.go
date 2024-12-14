package canoto

// Message defines a type that can be a stand-alone Canoto message.
type Message interface {
	Field
	// MarshalCanoto returns the Canoto representation of this message.
	//
	// It is assumed that this message is ValidCanoto.
	MarshalCanoto() []byte
	// UnmarshalCanoto unmarshals a Canoto-encoded byte slice into the message.
	//
	// The message is not cleared before unmarshaling, any fields not present in
	// the bytes will retain their previous values.
	UnmarshalCanoto(bytes []byte) error
}

// Field defines a type that can be included inside of a Canoto message.
type Field interface {
	// MarshalCanotoInto writes the field into a canoto.Writer.
	//
	// It is assumed that CalculateCanotoSize has been called since the last
	// modification to this field.
	//
	// It is assumed that this field is ValidCanoto.
	MarshalCanotoInto(w *Writer)
	// CalculateCanotoSize calculates the size of this field's Canoto
	// representation and caches it.
	CalculateCanotoSize() int
	// CachedCanotoSize returns the previously calculated size of the Canoto
	// representation from CalculateCanotoSize.
	//
	// If CalculateCanotoSize has not yet been called, it will return 0.
	//
	// If the field has been modified since the last call to
	// CalculateCanotoSize, the returned size may be incorrect.
	CachedCanotoSize() int
	// UnmarshalCanotoFrom populates the field from a canoto.Reader.
	//
	// The field is not cleared before unmarshaling, any sub-fields not present
	// in the bytes will retain their previous values.
	UnmarshalCanotoFrom(r *Reader) error
	// ValidCanoto validates that the field can be correctly marshaled into the
	// Canoto format.
	ValidCanoto() bool
}
