package canoto

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/structtag"
)

const canotoTag = "canoto"

var (
	errUnexpectedNumberOfIdentifiers       = errors.New("unexpected number of identifiers")
	errMalformedTag                        = errors.New("expected type,fieldNumber got")
	errFixedLengthArraysUnsupported        = errors.New("fixed length arrays are not supported")
	errRepeatedFieldsUnsupported           = errors.New("repeated fields are not supported")
	errStructContainsDuplicateFieldNumbers = errors.New("struct contains duplicate field numbers")
)

func parse(fs *token.FileSet, f ast.Node) (string, []message, error) {
	var (
		packageName string
		messages    []message
		err         error
	)
	ast.Inspect(f, func(n ast.Node) bool {
		if err != nil {
			return false
		}

		if f, ok := n.(*ast.File); ok {
			packageName = f.Name.Name
			return true
		}

		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			return false
		}

		message := message{
			name: ts.Name.Name,
		}
		for _, sf := range st.Fields.List {
			var (
				field  field
				hasTag bool
			)
			field, hasTag, err = parseField(fs, sf)
			if err != nil {
				return false
			}
			if !hasTag {
				continue
			}
			message.fields = append(message.fields, field)
		}
		if len(message.fields) == 0 {
			return false
		}

		slices.SortFunc(message.fields, field.Compare)
		if !isUniquelySorted(message.fields, field.Compare) {
			err = fmt.Errorf("%w at %s",
				errStructContainsDuplicateFieldNumbers,
				fs.Position(st.Pos()),
			)
			return false
		}

		messages = append(messages, message)
		return false
	})
	return packageName, messages, err
}

func parseField(fs *token.FileSet, af *ast.Field) (field, bool, error) {
	canotoType, fieldNumber, hasTag, err := parseFieldTag(fs, af)
	if err != nil || !hasTag {
		return field{}, false, err
	}

	if len(af.Names) != 1 {
		return field{}, false, fmt.Errorf("%w wanted %d got %d at %s",
			errUnexpectedNumberOfIdentifiers,
			1,
			len(af.Names),
			fs.Position(af.Pos()),
		)
	}

	f := field{
		name:        af.Names[0].Name,
		canotoType:  canotoType,
		fieldNumber: fieldNumber,
	}
	switch t := af.Type.(type) {
	case *ast.Ident:
		f.goType = t.Name
	case *ast.ArrayType:
		// TODO: Support fixed length arrays
		if t.Len != nil {
			return field{}, false, fmt.Errorf("%w at %s",
				errFixedLengthArraysUnsupported,
				fs.Position(t.Len.Pos()),
			)
		}

		ident, ok := t.Elt.(*ast.Ident)
		if !ok {
			return field{}, false, fmt.Errorf("%w %T at %s",
				errUnexpectedType,
				t.Elt,
				fs.Position(t.Elt.Pos()),
			)
		}

		if ident.Name == "byte" {
			f.goType = "[]byte"
		} else {
			return field{}, false, fmt.Errorf("%w at %s",
				errRepeatedFieldsUnsupported,
				fs.Position(t.Elt.Pos()),
			)
		}
	default:
		return field{}, false, fmt.Errorf("%w %T at %s",
			errUnexpectedType,
			t,
			fs.Position(af.Pos()),
		)
	}
	return f, true, nil
}

func parseFieldTag(fs *token.FileSet, field *ast.Field) (string, uint32, bool, error) {
	if field.Tag == nil {
		return "", 0, false, nil
	}

	rawTag := strings.Trim(field.Tag.Value, "`")
	tags, err := structtag.Parse(rawTag)
	if err != nil {
		return "", 0, false, err
	}

	tag, err := tags.Get(canotoTag)
	if err != nil {
		return "", 0, false, nil //nolint: nilerr // errors imply the tag was not round
	}

	if len(tag.Options) != 1 {
		return "", 0, false, fmt.Errorf("%w %s at %s",
			errMalformedTag,
			tag.Value(),
			fs.Position(field.Pos()),
		)
	}

	fieldNumber, err := strconv.ParseUint(tag.Options[0], 10, 32)
	if err != nil {
		return "", 0, false, err
	}

	return tag.Name, uint32(fieldNumber), true, nil
}

func isUniquelySorted[S ~[]E, E any](x S, cmp func(a E, b E) int) bool {
	for i := 1; i < len(x); i++ {
		if cmp(x[i-1], x[i]) >= 0 {
			return false
		}
	}
	return true
}
