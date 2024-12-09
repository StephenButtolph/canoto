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

var errNonGoExtension = errors.New("file must be a go file")

type message struct {
	name              string
	canonicalizedName string
	fields            []field
}

type field struct {
	name              string
	canonicalizedName string
	goType            goType
	canotoType        canotoType
	fieldNumber       uint32
	templateArgs      map[string]string
}

func (f field) Compare(other field) int {
	return cmp.Compare(f.fieldNumber, other.fieldNumber)
}

func (f field) WireType() (WireType, error) {
	switch f.canotoType {
	case canotoInt, canotoSint, canotoBool:
		return Varint, nil
	case canotoFint:
		switch f.goType {
		case goInt32, goUint32:
			return I32, nil
		case goInt64, goUint64:
			return I64, nil
		default:
			return 0, fmt.Errorf("%w: %q with canotoType %q", errUnexpectedGoType, f.goType, f.canotoType)
		}
	case canotoBytes:
		return Len, nil
	default:
		return 0, fmt.Errorf("%w: %q", errUnexpectedCanotoType, f.canotoType)
	}
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
