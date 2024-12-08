package canoto

import (
	"cmp"
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

var (
	errNonGoExtension = errors.New("file must be a go file")
	errUnexpectedType = errors.New("unexpected type")
)

type message struct {
	name              string
	canonicalizedName string
	fields            []field
}

type field struct {
	name              string
	canonicalizedName string
	goType            string
	canotoType        string
	fieldNumber       uint32
}

func (f field) Compare(other field) int {
	return cmp.Compare(f.fieldNumber, other.fieldNumber)
}

func Generate(inputFilePath string) error {
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

	outputFilePath := inputFilePath[:len(inputFilePath)-len(goExtension)] + canotoExtension
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return write(outputFile, packageName, messages)
}
