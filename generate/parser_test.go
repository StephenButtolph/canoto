package generate

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string

		filePath string
		opts     Options

		wantPackageName string
		wantMessages    []message
		wantErr         error
	}{
		{
			name:            "duplicate field name",
			filePath:        "testdata/duplicate_field_name.go",
			wantPackageName: "testdata",
			wantErr:         errInvalidOneOfName,
		},
		{
			name:            "duplicate field number",
			filePath:        "testdata/duplicate_field_number.go",
			wantPackageName: "testdata",
			wantErr:         errStructContainsDuplicateFieldNumbers,
		},
		{
			name:            "field number 0",
			filePath:        "testdata/field_number_0.go",
			wantPackageName: "testdata",
			wantErr:         errInvalidFieldNumber,
		},
		{
			name:            "field number too large",
			filePath:        "testdata/field_number_too_large.go",
			wantPackageName: "testdata",
			wantErr:         errInvalidFieldNumber,
		},
		{
			name:            "malformed after canoto",
			filePath:        "testdata/malformed_after_canoto.go",
			wantPackageName: "testdata",
			wantMessages: []message{
				{
					name:              "malformedAfterCanoto",
					canonicalizedName: "malformedAfterCanoto",
					fields: []field{
						{
							name:              "Int",
							canonicalizedName: "Int",
							goType:            "int64",
							isLocal:           true,
							protoType:         "sint64",
							canotoType:        canotoInt,
							fieldNumber:       1,
							templateArgs: map[string]string{
								"selector":                  "canoto.",
								"fieldNumber":               "1",
								"wireType":                  "Varint",
								"goType":                    "int64",
								"genericTypeCast":           "",
								"castOpen":                  "",
								"castClose":                 "",
								"protoType":                 "sint64",
								"protoTypePrefix":           "",
								"protoTypeSuffix":           "sint64",
								"fieldName":                 "Int",
								"suffix":                    "Int",
								"oneOf":                     "",
								"unmarshalOneOf":            "",
								"firstUnmarshalOneOf":       "",
								"sizeOneOf":                 "",
								"sizeOneOfIndent":           "",
								"checkDuplicateOneOf":       "",
								"checkDuplicateOneOfIndent": "",
								"signedSpace":               " ",
								"wrapperName":               "int64__Pointer",
							},
						},
					},
				},
			},
		},
		{
			name:            "malformed struct tag",
			filePath:        "testdata/malformed_struct_tag.go",
			wantPackageName: "testdata",
		},
		{
			name:            "missing field number",
			filePath:        "testdata/missing_field_number.go",
			wantPackageName: "testdata",
			wantErr:         errMalformedTag,
		},
		{
			name:            "multiple fields per type",
			filePath:        "testdata/multiple_fields_per_type.go",
			wantPackageName: "testdata",
			wantErr:         errUnexpectedNumberOfIdentifiers,
		},
		{
			name:            "quoted struct tag",
			filePath:        "testdata/quoted_struct_tag.go",
			wantPackageName: "testdata",
			wantErr:         errInvalidFieldNumber,
		},
		{
			name:            "repeated oneof",
			filePath:        "testdata/repeated_oneof.go",
			wantPackageName: "testdata",
			wantErr:         errRepeatedOneOf,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)

			// Create a new parser
			fs := token.NewFileSet()
			f, err := parser.ParseFile(fs, test.filePath, nil, parser.ParseComments)
			require.NoError(err)

			packageName, messages, err := parse(fs, f, test.opts)
			require.ErrorIs(err, test.wantErr)
			require.Equal(test.wantPackageName, packageName)
			require.Equal(test.wantMessages, messages)
		})
	}
}
