package generate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/StephenButtolph/canoto"
)

const (
	defaultCanotoSelector             = "canoto"
	fileName                          = "canoto.go"
	readWrite                         = 0o640
	readWriteExecute      os.FileMode = 0o750
)

// Library generates the canoto serialization primitives package in the provided
// folder.
//
// Specifically, if "./internal" is provided, this will generate the package
// "./internal/canoto" with the serialization primitives included in
// "./internal/canoto/canoto.go".
func Library(parentDir string) error {
	library := filepath.Join(parentDir, defaultCanotoSelector)
	if err := os.MkdirAll(library, readWriteExecute); err != nil {
		return fmt.Errorf("failed to create directory %q: %w", library, err)
	}

	path := filepath.Join(library, fileName)
	if err := os.WriteFile(path, []byte(canoto.Code), readWrite); err != nil {
		return fmt.Errorf("failed to write file %q: %w", path, err)
	}
	return nil
}
