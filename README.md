# Canoto

Canoto is a serialization format designed to be:
1. Fast
2. Compact
3. Canonical
4. Backwards compatible
5. Read compatible with [Protocol Buffers](https://protobuf.dev/).

## Install

```sh
go install github.com/StephenButtolph/canoto/canoto@latest
```

## Define Messages

Canoto messages are defined as normal golang structs:
```golang
type ExampleStruct0 struct {
	Int32       int32          `canoto:"int,1"`
	Int64       int64          `canoto:"int,2"`
	Uint32      uint32         `canoto:"int,3"`
	Uint64      uint64         `canoto:"int,4"`
	Sint32      int32          `canoto:"sint,5"`
	Sint64      int64          `canoto:"sint,6"`
	Fixed32     uint32         `canoto:"fint32,7"`
	Fixed64     uint64         `canoto:"fint64,8"`
	Sfixed32    int32          `canoto:"fint32,9"`
	Sfixed64    int64          `canoto:"fint64,10"`
	Bool        bool           `canoto:"bool,11"`
	String      string         `canoto:"string,12"`
	Bytes       []byte         `canoto:"bytes,13"`
	OtherStruct ExampleStruct1 `canoto:"field,14"`

	canotoData canotoData_ExampleStruct0
}

type ExampleStruct1 struct {
	Int32 int32 `canoto:"int,536870911"`
}
```

All structs must include a field called `canotoData` that will cache the results of calculating the size of the struct.

The type `canotoData_${structName}` is automatically generated by Canoto.

Canoto implements the following `Message` interface for the struct:

```golang
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
```

## Generate

In order to generate canoto information for all of the structs in a file, simply run the `canoto` command with one or more files.

```sh
canoto example0.go example1.go
```

The above example will generate `example0.canoto.go` and `example1.canoto.go`.

The corresponding `proto` file for a `canoto` file can also be generated by adding the `--proto`.

```sh
canoto --proto example.go
```

The above example will generate `example.canoto.go` and `example.proto`.

### go:generate

To automatically generate the `.canoto.go` version of a file, it is recommended to use `go:generate`

Placing

```golang
//go:generate canoto $GOFILE
```

at the top of a file will update the `.canoto.go` version of the file every time `go generate ./...` is run.

### Best Practices

`canoto` only inspects a single golang file at a time, so it is recommended to define nested messages in the same file to be able to generate the most useful `proto` file.

Additionally, while fully supported in the `canoto` output, type aliases and generic types will result in `proto` files with default types. It is still guaranteed for the generated `proto` file to be able to parse `canoto` data, but the types may not be as specific as they could be.

If type aliases are needed, it may make sense to modify the generated proto file to specify the most specific proto type possible.

## Supported Types

| go type        | canoto type                  | proto type          | wire type |
|----------------|------------------------------|---------------------|-----------|
| `int8`         | `int`                        | `int32`             | `varint`  |
| `int16`        | `int`                        | `int32`             | `varint`  |
| `int32`        | `int`                        | `int32`             | `varint`  |
| `int64`        | `int`                        | `int64`             | `varint`  |
| `uint8`        | `int`                        | `uint32`            | `varint`  |
| `uint16`       | `int`                        | `uint32`            | `varint`  |
| `uint32`       | `int`                        | `uint32`            | `varint`  |
| `uint64`       | `int`                        | `uint64`            | `varint`  |
| `int8`         | `sint`                       | `sint32`            | `varint`  |
| `int16`        | `sint`                       | `sint32`            | `varint`  |
| `int32`        | `sint`                       | `sint32`            | `varint`  |
| `int64`        | `sint`                       | `sint64`            | `varint`  |
| `uint32`       | `fint32`                     | `fixed32`           | `i32`     |
| `uint64`       | `fint64`                     | `fixed64`           | `i64`     |
| `int32`        | `fint32`                     | `sfixed32`          | `i32`     |
| `int64`        | `fint64`                     | `sfixed64`          | `i64`     |
| `bool`         | `bool`                       | `bool`              | `varint`  |
| `string`       | `string`                     | `string`            | `len`     |
| `[]byte`       | `bytes`                      | `bytes`             | `len`     |
| `[x]byte`      | `fixed bytes`                | `bytes`             | `len`     |
| `T Field`      | `field`                      | `bytes`             | `len`     |
| `T Message`    | `field`                      | `message`           | `len`     |
| `[]int8`       | `repeated int`               | `repeated int32`    | `len`     |
| `[]int16`      | `repeated int`               | `repeated int32`    | `len`     |
| `[]int32`      | `repeated int`               | `repeated int32`    | `len`     |
| `[]int64`      | `repeated int`               | `repeated int64`    | `len`     |
| `[]uint8`      | `repeated int`               | `repeated uint32`   | `len`     |
| `[]uint16`     | `repeated int`               | `repeated uint32`   | `len`     |
| `[]uint32`     | `repeated int`               | `repeated uint32`   | `len`     |
| `[]uint64`     | `repeated int`               | `repeated uint64`   | `len`     |
| `[]int8`       | `repeated sint`              | `repeated sint32`   | `len`     |
| `[]int16`      | `repeated sint`              | `repeated sint32`   | `len`     |
| `[]int32`      | `repeated sint`              | `repeated sint32`   | `len`     |
| `[]int64`      | `repeated sint`              | `repeated sint64`   | `len`     |
| `[]uint32`     | `repeated fint32`            | `repeated fixed32`  | `len`     |
| `[]uint64`     | `repeated fint64`            | `repeated fixed64`  | `len`     |
| `[]int32`      | `repeated fint32`            | `repeated sfixed32` | `len`     |
| `[]int64`      | `repeated fint64`            | `repeated sfixed64` | `len`     |
| `[]bool`       | `repeated bool`              | `repeated bool`     | `len`     |
| `[]string`     | `repeated string`            | `repeated string`   | `len`     |
| `[][]byte`     | `repeated bytes`             | `repeated bytes`    | `len`     |
| `[][x]byte`    | `repeated fixed bytes`       | `repeated bytes`    | `len`     |
| `[]T Field`    | `repeated field`             | `repeated bytes`    | `len`     |
| `[]T Message`  | `repeated field`             | `repeated message`  | `len`     |
| `[x]int8`      | `fixed repeated int`         | `repeated int32`    | `len`     |
| `[x]int16`     | `fixed repeated int`         | `repeated int32`    | `len`     |
| `[x]int32`     | `fixed repeated int`         | `repeated int32`    | `len`     |
| `[x]int64`     | `fixed repeated int`         | `repeated int64`    | `len`     |
| `[x]uint8`     | `fixed repeated int`         | `repeated uint32`   | `len`     |
| `[x]uint16`    | `fixed repeated int`         | `repeated uint32`   | `len`     |
| `[x]uint32`    | `fixed repeated int`         | `repeated uint32`   | `len`     |
| `[x]uint64`    | `fixed repeated int`         | `repeated uint64`   | `len`     |
| `[x]int8`      | `fixed repeated sint`        | `repeated sint32`   | `len`     |
| `[x]int16`     | `fixed repeated sint`        | `repeated sint32`   | `len`     |
| `[x]int32`     | `fixed repeated sint`        | `repeated sint32`   | `len`     |
| `[x]int64`     | `fixed repeated sint`        | `repeated sint64`   | `len`     |
| `[x]uint32`    | `fixed repeated fint32`      | `repeated fixed32`  | `len`     |
| `[x]uint64`    | `fixed repeated fint64`      | `repeated fixed64`  | `len`     |
| `[x]int32`     | `fixed repeated fint32`      | `repeated sfixed32` | `len`     |
| `[x]int64`     | `fixed repeated fint64`      | `repeated sfixed64` | `len`     |
| `[x]bool`      | `fixed repeated bool`        | `repeated bool`     | `len`     |
| `[x]string`    | `fixed repeated string`      | `repeated string`   | `len`     |
| `[x][]byte`    | `fixed repeated bytes`       | `repeated bytes`    | `len`     |
| `[x][y]byte`   | `fixed repeated fixed bytes` | `repeated bytes`    | `len`     |
| `[x]T Field`   | `fixed repeated field`       | `repeated bytes`    | `len`     |
| `[x]T Message` | `fixed repeated field`       | `repeated message`  | `len`     |

### Non-standard encoding

It is valid to define a `Field` that implements a non-standard format. However, this format should still be canonical and the corresponding Proto file should report opaque bytes.

## Why not Proto?

Proto is a fast, compact, encoding format with extensive language support. However, [Proto is not canonical](https://protobuf.dev/programming-guides/serialization-not-canonical/).

Proto is designed to be forwards-compatible. Almost by definition, a forwards-compatible serialization format can not be canonical. The format of a field can not validated to be canonical if the expected type of the field is not known during decoding.

## Why is being canonical important?

In some cases, non-canonical serialization formats are subtle to work with.

For example, if the hash of the serialized data is important or if the serialized data is cryptographically signed.

In order to ensure that the hash of the serialized data does not change, it is important to carefully avoid re-serializing a message that was previously serialized.

For canonical serialization formats, the hash of the serialized data is guaranteed never to change. Every correct implementation of the format will produce the same hash.

## Why be read compatible with Proto?

By being read compatible with Proto, users of the Canoto format inherit some Proto's cross language support.

If an application only needs to read Canoto messages, but not write them, it can simply treat the Canoto message as a Proto message.

## Is Canoto Fast?

Canoto is typically more performant for both serialization and deserialization than Proto. However, Proto does not typically validate that fields are canonical. If a field is expensive to inspect, it's possible Canoto can be slightly slower.

Canoto is optimized to perform no unnecessary memory allocations, so careful management to ensure messages are stack allocated can significantly improve performance over Proto.

## Is Canoto Forwards Compatible?

No. Canoto chooses to be a canonical serialization format rather than being forwards compatible.
