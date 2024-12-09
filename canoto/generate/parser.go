package generate

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

		name := ts.Name.Name
		message := message{
			name:              name,
			canonicalizedName: canonicalizeName(name),
		}
		for _, sf := range st.Fields.List {
			var (
				field  field
				hasTag bool
			)
			field, hasTag, err = parseField(fs, message.canonicalizedName, sf)
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

func parseField(fs *token.FileSet, canonicalizedStructName string, af *ast.Field) (field, bool, error) {
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

	name := af.Names[0].Name
	f := field{
		name:              name,
		canonicalizedName: canonicalizeName(name),
		canotoType:        canotoType,
		fieldNumber:       fieldNumber,
	}
	switch t := af.Type.(type) {
	case *ast.Ident:
		f.goType = goType(t.Name)
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
				errUnexpectedGoType,
				t.Elt,
				fs.Position(t.Elt.Pos()),
			)
		}

		if ident.Name == "byte" {
			f.goType = goBytes
		} else {
			return field{}, false, fmt.Errorf("%w at %s",
				errRepeatedFieldsUnsupported,
				fs.Position(t.Elt.Pos()),
			)
		}
	default:
		return field{}, false, fmt.Errorf("%w %T at %s",
			errUnexpectedGoType,
			t,
			fs.Position(af.Pos()),
		)
	}

	f.templateArgs, err = makeTemplateArgs(canonicalizedStructName, f)
	return f, true, err
}

// canonicalizeName replaces "_" with "_1" to avoid collisions with "__" which
// is used as a reserved separator.
func canonicalizeName(name string) string {
	return strings.ReplaceAll(name, "_", "_1")
}

// parseFieldTag parses the tag of the provided field and returns the canoto
// description, if one exists.
func parseFieldTag(fs *token.FileSet, field *ast.Field) (canotoType, uint32, bool, error) {
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
		return "", 0, false, nil //nolint: nilerr // errors imply the tag was not found
	}

	fieldType := canotoType(tag.Name)
	if !fieldType.IsValid() {
		return "", 0, false, fmt.Errorf("%w %s at %s",
			errUnexpectedCanotoType,
			tag.Name,
			fs.Position(field.Pos()),
		)
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
	return fieldType, uint32(fieldNumber), true, nil
}

// isUniquelySorted returns true if the provided slice is sorted in ascending
// order and contains no duplicates.
func isUniquelySorted[S ~[]E, E any](x S, cmp func(a E, b E) int) bool {
	for i := 1; i < len(x); i++ {
		if cmp(x[i-1], x[i]) >= 0 {
			return false
		}
	}
	return true
}

func makeTemplateArgs(structName string, field field) (map[string]string, error) {
	wireType, err := field.WireType()
	if err != nil {
		return nil, err
	}

	args := map[string]string{
		"escapedStructName": structName,
		"fieldNumber":       strconv.FormatUint(uint64(field.fieldNumber), 10),
		"wireType":          wireType.String(),
		"fieldName":         field.name,
		"escapedFieldName":  field.canonicalizedName,
	}
	switch field.canotoType {
	case canotoInt:
		args["readFunction"] = fmt.Sprintf("Int[%s]", field.goType)
	case canotoSint:
		args["readFunction"] = fmt.Sprintf("Sint[%s]", field.goType)
	case canotoFint:
		switch field.goType {
		case goInt32, goUint32:
			args["readFunction"] = fmt.Sprintf("Fint32[%s]", field.goType)
			args["bitSize"] = "32"
		case goInt64, goUint64:
			args["readFunction"] = fmt.Sprintf("Fint64[%s]", field.goType)
			args["bitSize"] = "64"
		default:
			return nil, fmt.Errorf("%w: %q should have fixed size", errUnexpectedGoType, field.goType)
		}
	case canotoBool:
	case canotoBytes:
		switch field.goType {
		case goString:
			args["readFunction"] = "String"
		case goBytes:
			args["readFunction"] = "Bytes"
		}
	default:
		return nil, fmt.Errorf("%w: %q", errUnexpectedCanotoType, field.canotoType)
	}
	return args, nil
}
