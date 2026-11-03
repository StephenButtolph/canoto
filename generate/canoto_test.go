package generate

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanotoFixedLength(t *testing.T) {
	tests := []struct {
		canotoType canotoType
		// goType is formatted with the length of the fixed array.
		goType string
	}{
		{canotoType: canotoFixedBytes, goType: "[%d]byte"},
		{canotoType: canotoRepeatedFixedBytes, goType: "[][%d]byte"},
		{canotoType: canotoFixedRepeatedInt, goType: "[%d]int64"},
		{canotoType: canotoFixedRepeatedUint, goType: "[%d]uint64"},
		{canotoType: canotoFixedRepeatedFint32, goType: "[%d]int32"},
		{canotoType: canotoFixedRepeatedFint64, goType: "[%d]int64"},
		{canotoType: canotoFixedRepeatedBool, goType: "[%d]bool"},
		{canotoType: canotoFixedRepeatedString, goType: "[%d]string"},
		{canotoType: canotoFixedRepeatedBytes, goType: "[%d][]byte"},
		{canotoType: canotoFixedRepeatedFixedBytes, goType: "[%d][1]byte"},
		{canotoType: canotoFixedRepeatedFixedBytes, goType: "[1][%d]byte"},
		{canotoType: canotoFixedRepeatedValue, goType: "[%d]inner"},
		{canotoType: canotoFixedRepeatedPointer, goType: "[%d]*inner"},
	}
	// Sharing the importer across tests avoids re-importing canoto.
	var (
		fs  = token.NewFileSet()
		imp = sourceImporter{
			importer: importer.ForCompiler(fs, "source", nil).(types.ImporterFrom),
		}
	)
	for _, test := range tests {
		for _, length := range []int{0, 1} {
			goType := fmt.Sprintf(test.goType, length)
			t.Run(fmt.Sprintf("%s %s", test.canotoType, goType), func(t *testing.T) {
				require := require.New(t)

				src := fmt.Sprintf(`package fixed

type outer struct {
	Field %s `+"`canoto:\"%s,1\"`"+`

	canotoData canotoData_outer
}

type inner struct {
	Field int64 `+"`canoto:\"int,1\"`"+`

	canotoData canotoData_inner
}
`, goType, test.canotoType)

				var (
					dir           = t.TempDir()
					path          = filepath.Join(dir, "fixed.go")
					generatedPath = filepath.Join(dir, "fixed.canoto.go")
				)
				require.NoError(os.WriteFile(path, []byte(src), 0o600))
				require.NoError(Canoto(path, Options{}))

				err := typeCheck(fs, imp, path, generatedPath)
				if length != 0 {
					require.NoError(err)
					return
				}

				// The generated code must not compile.
				var typeErr types.Error
				require.ErrorAs(err, &typeErr)
				pos := typeErr.Fset.Position(typeErr.Pos)
				require.Equal("fixed.canoto.go", filepath.Base(pos.Filename))
			})
		}
	}
}

// typeCheck type checks the package made up of paths.
func typeCheck(fs *token.FileSet, imp types.Importer, paths ...string) error {
	files := make([]*ast.File, len(paths))
	for i, path := range paths {
		f, err := parser.ParseFile(fs, path, nil, 0)
		if err != nil {
			return err
		}
		files[i] = f
	}

	conf := types.Config{
		Importer: imp,
	}
	_, err := conf.Check("fixed", fs, files, nil)
	return err
}

// sourceImporter imports packages from source, resolving them relative to the
// current directory rather than the directory of the importing package. This
// allows packages outside of the module to import packages from the module.
type sourceImporter struct {
	importer types.ImporterFrom
}

func (s sourceImporter) Import(path string) (*types.Package, error) {
	return s.importer.ImportFrom(path, ".", 0)
}
