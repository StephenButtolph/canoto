package generate

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func defaultTemplates() Templates {
	var t Templates
	t.init()
	return t
}

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
			name:            "duplicate_field_name",
			filePath:        "testdata/duplicate_field_name.go",
			wantPackageName: "testdata",
			wantErr:         errInvalidOneOfName,
		},
		{
			name:            "duplicate_field_number",
			filePath:        "testdata/duplicate_field_number.go",
			wantPackageName: "testdata",
			wantErr:         errStructContainsDuplicateFieldNumbers,
		},
		{
			name:            "field_number_0",
			filePath:        "testdata/field_number_0.go",
			wantPackageName: "testdata",
			wantErr:         errInvalidFieldNumber,
		},
		{
			name:            "field_number_too_large",
			filePath:        "testdata/field_number_too_large.go",
			wantPackageName: "testdata",
			wantErr:         errInvalidFieldNumber,
		},
		{
			name:            "field_number_before_reserved",
			filePath:        "testdata/field_number_before_reserved.go",
			wantPackageName: "testdata",
			wantMessages: []message{
				{
					name:              "fieldNumberBeforeReserved",
					canonicalizedName: "fieldNumberBeforeReserved",
					template:          defaultTemplates(),
					fields: []field{
						{
							name:              "Int",
							canonicalizedName: "Int",
							goType:            "int64",
							isLocal:           true,
							protoType:         "sint64",
							canotoType:        canotoInt,
							fieldNumber:       18999,
							templateArgs: map[string]string{
								"selector":                  "canoto.",
								"fieldNumber":               "18999",
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
			name:            "field_number_reserved_first",
			filePath:        "testdata/field_number_reserved_first.go",
			wantPackageName: "testdata",
			wantErr:         errInvalidFieldNumber,
		},
		{
			name:            "field_number_reserved_last",
			filePath:        "testdata/field_number_reserved_last.go",
			wantPackageName: "testdata",
			wantErr:         errInvalidFieldNumber,
		},
		{
			name:            "field_number_after_reserved",
			filePath:        "testdata/field_number_after_reserved.go",
			wantPackageName: "testdata",
			wantMessages: []message{
				{
					name:              "fieldNumberAfterReserved",
					canonicalizedName: "fieldNumberAfterReserved",
					template:          defaultTemplates(),
					fields: []field{
						{
							name:              "Int",
							canonicalizedName: "Int",
							goType:            "int64",
							isLocal:           true,
							protoType:         "sint64",
							canotoType:        canotoInt,
							fieldNumber:       20000,
							templateArgs: map[string]string{
								"selector":                  "canoto.",
								"fieldNumber":               "20000",
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
			name:            "function_scoped_struct",
			filePath:        "testdata/function_scoped_struct.go",
			wantPackageName: "testdata",
		},
		{
			name:            "malformed_after_canoto",
			filePath:        "testdata/malformed_after_canoto.go",
			wantPackageName: "testdata",
			wantMessages: []message{
				{
					name:              "malformedAfterCanoto",
					canonicalizedName: "malformedAfterCanoto",
					template:          defaultTemplates(),
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
			name:            "malformed_struct_tag",
			filePath:        "testdata/malformed_struct_tag.go",
			wantPackageName: "testdata",
		},
		{
			name:            "missing_field_number",
			filePath:        "testdata/missing_field_number.go",
			wantPackageName: "testdata",
			wantErr:         errMalformedTag,
		},
		{
			name:            "multiple_fields_per_type",
			filePath:        "testdata/multiple_fields_per_type.go",
			wantPackageName: "testdata",
			wantErr:         errUnexpectedNumberOfIdentifiers,
		},
		{
			name:            "pointer_canoto_data",
			filePath:        "testdata/pointer_canoto_data.go",
			wantPackageName: "testdata",
			wantErr:         errInvalidCanotoData,
		},
		{
			name:            "promoted_pointer_canoto_data",
			filePath:        "testdata/promoted_pointer_canoto_data.go",
			wantPackageName: "testdata",
			wantErr:         errMissingCanotoData,
		},
		{
			name:            "promoted_value_canoto_data",
			filePath:        "testdata/promoted_value_canoto_data.go",
			wantPackageName: "testdata",
			wantErr:         errMissingCanotoData,
		},
		{
			name:            "quoted_struct_tag",
			filePath:        "testdata/quoted_struct_tag.go",
			wantPackageName: "testdata",
			wantErr:         errInvalidFieldNumber,
		},
		{
			name:            "repeated_oneof",
			filePath:        "testdata/repeated_oneof.go",
			wantPackageName: "testdata",
			wantErr:         errRepeatedOneOf,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require := require.New(t)

			fs := token.NewFileSet()
			f, err := parser.ParseFile(fs, test.filePath, nil, parser.ParseComments)
			require.NoError(err)

			test.opts.init()
			packageName, messages, err := parse(fs, f, test.opts)
			require.ErrorIs(err, test.wantErr)
			require.Equal(test.wantPackageName, packageName)
			require.Equal(test.wantMessages, messages)
		})
	}
}
