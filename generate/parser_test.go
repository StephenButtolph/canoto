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

		filePath     string
		canotoImport string
		internal     bool

		wantPackageName string
		wantMessages    []message
		wantErr         error
	}{
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
			name:            "missing field number",
			filePath:        "testdata/missing_field_number.go",
			wantPackageName: "testdata",
			wantErr:         errMalformedTag,
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

			packageName, messages, err := parse(fs, f, test.canotoImport, test.internal)
			require.ErrorIs(err, test.wantErr)
			require.Equal(test.wantPackageName, packageName)
			require.Equal(test.wantMessages, messages)
		})
	}
}
