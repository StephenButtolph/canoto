package generate

import (
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteCanoto(t *testing.T) {
	tests := []struct {
		name string

		filePath string
		opts     Options

		wantErr error
	}{
		{
			name:     "default_templates",
			filePath: "testdata/shadowed_number.go",
		},
		{
			name:     "shadowed_number",
			filePath: "testdata/shadowed_number.go",
			opts: Options{
				Templates: Templates{
					Number: "{field}",
				},
			},
			wantErr: errShadowedIdentifier,
		},
		{
			name:     "shadowed_tag",
			filePath: "testdata/shadowed_tag.go",
			opts: Options{
				Templates: Templates{
					Tag: "{field}",
				},
			},
			wantErr: errShadowedIdentifier,
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
			require.NoError(err)

			var output strings.Builder
			err = writeCanoto(
				&output,
				test.filePath,
				packageName,
				messages,
				test.opts.CanotoImport,
				test.opts.Internal,
			)
			t.Logf("generated output:\n%s", output.String())
			require.ErrorIs(err, test.wantErr)
		})
	}
}
