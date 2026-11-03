package canoto_test

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/StephenButtolph/canoto"
)

func ExampleNewCodec() {
	// The sender encodes a message along with the Spec describing it. Any
	// generated type works; Spec is used here because it is itself a canoto
	// message.
	msg := &canoto.Spec{
		Name: "Point",
		Fields: []canoto.FieldType{
			{FieldNumber: 1, Name: "x", TypeInt: canoto.SizeEnum64},
			{FieldNumber: 2, Name: "y", TypeInt: canoto.SizeEnum64},
		},
	}
	msgBytes := msg.MarshalCanoto()
	specBytes := msg.DescribeCanoto().MarshalCanoto()

	// The receiver decodes the message without knowing its Go type. Both
	// inputs may be untrusted, so memory usage and nesting depth are limited.
	const (
		memoryLimit = 1 << 20
		depthLimit  = 16
	)

	var spec canoto.Spec
	_, err := spec.UnmarshalCanotoFrom(canoto.Reader{
		B:               specBytes,
		MemoryRemaining: memoryLimit,
		DepthLimit:      depthLimit,
	})
	if err != nil {
		log.Fatal(err)
	}

	codec, err := canoto.NewCodec(&spec)
	if err != nil {
		log.Fatal(err)
	}

	decoded, _, err := codec.Unmarshal(canoto.Reader{
		B:               msgBytes,
		MemoryRemaining: memoryLimit,
		DepthLimit:      depthLimit,
	})
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := json.MarshalIndent(decoded, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonBytes))
	// Output:
	// {
	//   "Name": "Point",
	//   "Fields": [
	//     {
	//       "FieldNumber": 1,
	//       "Name": "x",
	//       "TypeInt": 4
	//     },
	//     {
	//       "FieldNumber": 2,
	//       "Name": "y",
	//       "TypeInt": 4
	//     }
	//   ]
	// }
}
