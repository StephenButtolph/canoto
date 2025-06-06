package generate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/StephenButtolph/canoto"
)

const (
	defaultCanotoSelector    = "canoto"
	libraryFileName          = "canoto.go"
	generatedLibraryFileName = "canoto.canoto.go"
	readWrite                = 0o640
	readWriteExecute         = 0o750

	codeGoGeneratePrefix = `//go:generate canoto --internal $GOFILE

`
	doNotGenerate = `// Code generated by canoto. DO NOT EDIT.
// versions:
// 	canoto ` + canoto.Version + `

`
)

// Library generates the canoto serialization primitives package in the provided
// folder.
//
// Specifically, if "./internal" is provided, this will generate the package
// "./internal/canoto" with the serialization primitives included in
// "./internal/canoto/canoto.go" and "./internal/canoto/canoto.canoto.go".
func Library(parentDir string) error {
	library := filepath.Join(parentDir, defaultCanotoSelector)
	if err := os.MkdirAll(library, readWriteExecute); err != nil {
		return fmt.Errorf("failed to create directory %q: %w", library, err)
	}

	libraryPath := filepath.Join(library, libraryFileName)
	codeWithoutPrefix, _ := strings.CutPrefix(canoto.Code, codeGoGeneratePrefix)
	if err := os.WriteFile(libraryPath, []byte(doNotGenerate+codeWithoutPrefix), readWrite); err != nil {
		return fmt.Errorf("failed to write file %q: %w", libraryPath, err)
	}

	generatedLibraryPath := filepath.Join(library, generatedLibraryFileName)
	if err := os.WriteFile(generatedLibraryPath, []byte(canoto.GeneratedCode), readWrite); err != nil {
		return fmt.Errorf("failed to write file %q: %w", libraryPath, err)
	}
	return nil
}
