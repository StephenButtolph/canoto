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
	Fixed32     uint32         `canoto:"fint,7"`
	Fixed64     uint64         `canoto:"fint,8"`
	Sfixed32    int32          `canoto:"fint,9"`
	Sfixed64    int64          `canoto:"fint,10"`
	Bool        bool           `canoto:"bool,11"`
	String      string         `canoto:"bytes,12"`
	Bytes       []byte         `canoto:"bytes,13"`
	OtherStruct ExampleStruct1 `canoto:"bytes,14"`

	canotoSize int
}

type ExampleStruct1 struct {
	Int32 int32 `canoto:"int,536870911"`

	canotoSize int
}
```

All structs must include a field called `canotoSize` that will cache the results of calculating the size of the struct.

Additionally, Canoto implements all of the following functions on the struct:

```golang
type Canoto interface {
  MarshalCanoto() []byte
  MarshalCanotoInto(w *canoto.Writer)
  CalculateCanotoSize() int
  CachedCanotoSize() int
  UnmarshalCanoto(bytes []byte) error
  UnmarshalCanotoFrom(r *canoto.Reader) error
  ValidCanoto() bool
}
```

## Generate

In order to generate canoto information for all of the structs in a file, simply run the `canoto` command with one or more files.

```sh
canoto example0.go example1.go
```

The above example will generate `example0.canoto.go` and `example1.canoto.go`.

## Why not Proto?

Proto is a fast, compact encoding format with extensive language support. However, [Proto is not canonical](https://protobuf.dev/programming-guides/serialization-not-canonical/).

Proto is designed to be forwards-compatible. Almost by definition, a forwards-compatible serialization format can not be canonical. The format of a field can not validated to be canonical if the expected type of the field is not known during decoding.

## Why is being canonical important?

In some cases, non-canonical serialization formats are subtle to work with.

For example, if the hash of the serialized data is important or if the serialized data is cryptographically signed.

In order to ensure that the hash of the serialized data does not change, it is important to carefully avoid re-serializing a message that was previously serialized.

For canonical serialization formats, the hash of the serialized data is guaranteed never to change, and every correct implementation of the format will produce the same hash.

## Why be read compatible with Proto?

By being read compatible with Proto, users of the Canoto format to inherit some of the cross language support that Proto has already made.

If an application only needs to read Canoto messages, but not write them, it can simply treat the Canoto message as a Proto message.

## Is Canoto Fast?

Canoto is more performant for both serialization and deserialization than Proto.

## Is Canoto Forwards Compatible?

No. Canoto chooses to be a canonical serialization format rather than being forwards compatible.
