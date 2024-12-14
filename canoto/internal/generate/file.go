package generate

import (
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

const (
	goExtension     = ".go"
	canotoExtension = ".canoto.go"
)

var errNonGoExtension = errors.New("file must be a go file")

func File(inputFilePath string) error {
	extension := filepath.Ext(inputFilePath)
	if extension != goExtension {
		return fmt.Errorf("%w not %q", errNonGoExtension, extension)
	}

	// Create a new parser
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, inputFilePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	packageName, messages, err := parse(fs, f)
	if err != nil {
		return err
	}
	if len(messages) == 0 {
		return nil
	}

	outputFilePath := inputFilePath[:len(inputFilePath)-len(goExtension)] + canotoExtension
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return write(outputFile, packageName, messages)
}
