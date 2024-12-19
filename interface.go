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
	// MarshalCanotoInto writes the field into a canoto.Writer.
	//
	// It is assumed that CalculateCanotoCache has been called since the last
	// modification to this field.
	//
	// It is assumed that this field is ValidCanoto.
	MarshalCanotoInto(w *Writer)
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
	UnmarshalCanotoFrom(r *Reader) error
	// ValidCanoto validates that the field can be correctly marshaled into the
	// Canoto format.
	ValidCanoto() bool
}
