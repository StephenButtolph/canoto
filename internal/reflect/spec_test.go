//go:generate canoto $GOFILE

package reflect

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type (
	testMessage struct {
		Int8       int8              `canoto:"int,1"`
		Int16      int16             `canoto:"int,2"`
		Int32      int32             `canoto:"int,3"`
		Int64      int64             `canoto:"int,4"`
		Uint8      uint8             `canoto:"int,5"`
		Uint16     uint16            `canoto:"int,6"`
		Uint32     uint32            `canoto:"int,7"`
		Uint64     uint64            `canoto:"int,8"`
		Sint8      int8              `canoto:"sint,9"`
		Sint16     int16             `canoto:"sint,10"`
		Sint32     int32             `canoto:"sint,11"`
		Sint64     int64             `canoto:"sint,12"`
		Fixed32    uint32            `canoto:"fint32,13"`
		Fixed64    uint64            `canoto:"fint64,14"`
		Sfixed32   int32             `canoto:"fint32,15"`
		Sfixed64   int64             `canoto:"fint64,16"`
		Bool       bool              `canoto:"bool,17"`
		String     string            `canoto:"string,18"`
		Bytes      []byte            `canoto:"bytes,19"`
		FixedBytes [32]byte          `canoto:"fixed bytes,20"`
		Recursive  *testMessage      `canoto:"pointer,21"`
		Message    testSimpleMessage `canoto:"value,22"`

		canotoData canotoData_testMessage
	}
	testSimpleMessage struct {
		Int8 int8 `canoto:"int,1"`

		canotoData canotoData_testMessage
	}
)

var testMessageSpec = Spec{
	Name: "testMessage",
	Fields: []*FieldType{
		{
			FieldNumber: 1,
			Name:        "Int8",
			TypeInt:     8,
		},
		{
			FieldNumber: 2,
			Name:        "Int16",
			TypeInt:     16,
		},
		{
			FieldNumber: 3,
			Name:        "Int32",
			TypeInt:     32,
		},
		{
			FieldNumber: 4,
			Name:        "Int64",
			TypeInt:     64,
		},
		{
			FieldNumber: 5,
			Name:        "Uint8",
			TypeUint:    8,
		},
		{
			FieldNumber: 6,
			Name:        "Uint16",
			TypeUint:    16,
		},
		{
			FieldNumber: 7,
			Name:        "Uint32",
			TypeUint:    32,
		},
		{
			FieldNumber: 8,
			Name:        "Uint64",
			TypeUint:    64,
		},
		{
			FieldNumber: 9,
			Name:        "Sint8",
			TypeSint:    8,
		},
		{
			FieldNumber: 10,
			Name:        "Sint16",
			TypeSint:    16,
		},
		{
			FieldNumber: 11,
			Name:        "Sint32",
			TypeSint:    32,
		},
		{
			FieldNumber: 12,
			Name:        "Sint64",
			TypeSint:    64,
		},
		{
			FieldNumber: 13,
			Name:        "Fixed32",
			TypeFint:    32,
		},
		{
			FieldNumber: 14,
			Name:        "Fixed64",
			TypeFint:    64,
		},
		{
			FieldNumber: 15,
			Name:        "Sfixed32",
			TypeSFint:   32,
		},
		{
			FieldNumber: 16,
			Name:        "Sfixed64",
			TypeSFint:   64,
		},
		{
			FieldNumber: 17,
			Name:        "Bool",
			TypeBool:    true,
		},
		{
			FieldNumber: 18,
			Name:        "String",
			TypeString:  true,
		},
		{
			FieldNumber: 19,
			Name:        "Bytes",
			TypeBytes:   true,
		},
		{
			FieldNumber:    20,
			Name:           "FixedBytes",
			TypeFixedBytes: 32,
		},
		{
			FieldNumber:   21,
			Name:          "Recursive",
			TypeRecursive: 1,
		},
		{
			FieldNumber: 22,
			Name:        "Message",
			TypeMessage: &Spec{
				Name: "testSimpleMessage",
				Fields: []*FieldType{
					{
						FieldNumber: 1,
						Name:        "Int8",
						TypeInt:     8,
					},
				},
			},
		},
	},
}

func TestSpec(t *testing.T) {
	require := require.New(t)

	m := testMessage{
		Int8:     1,
		Int16:    2,
		Int32:    3,
		Int64:    4,
		Uint8:    5,
		Uint16:   6,
		Uint32:   7,
		Uint64:   8,
		Sint8:    9,
		Sint16:   10,
		Sint32:   11,
		Sint64:   12,
		Fixed32:  13,
		Fixed64:  14,
		Sfixed32: 15,
		Sfixed64: 16,
		Bool:     true,
		String:   "string",
		Bytes:    []byte("bytes"),
		FixedBytes: [32]byte{
			0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
			0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
			0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
			0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		},
		Recursive: &testMessage{
			Int8: 21,
		},
		Message: testSimpleMessage{
			Int8: 22,
		},
	}
	mBytes := m.MarshalCanoto()

	a, err := testMessageSpec.Unmarshal(mBytes)
	require.NoError(err)

	specJSON, err := json.MarshalIndent(&testMessageSpec, "", "  ")
	require.NoError(err)
	t.Log(string(specJSON))

	aJSON, err := json.MarshalIndent(a, "", "  ")
	require.NoError(err)
	t.Log(string(aJSON))

	t.Fail()
}

func TestSpecSpec(t *testing.T) {
	require := require.New(t)

	specSpec := (*Spec)(nil).CanotoSpec()
	specBytes := specSpec.MarshalCanoto()
	msg, err := specSpec.Unmarshal(specBytes)
	require.NoError(err)

	specJSON, err := json.MarshalIndent(specSpec, "", "  ")
	require.NoError(err)
	t.Log(string(specJSON))

	msgJSON, err := json.MarshalIndent(msg, "", "  ")
	require.NoError(err)
	t.Log(string(msgJSON))

	t.Fail()
}
