package generate

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/structtag"
)

const (
	canotoImport          = `"github.com/StephenButtolph/canoto"`
	defaultCanotoSelector = "canoto"
	canotoTag             = "canoto"
	goBytes               = "[]byte"
)

var (
	oneOfRegex = regexp.MustCompile(`\A[a-zA-Z0-9_]+\z`)

	errUnexpectedNumberOfIdentifiers       = errors.New("unexpected number of identifiers")
	errInvalidGoType                       = errors.New("invalid Go type")
	errMalformedTag                        = errors.New("expected type,fieldNumber[,oneof] got")
	errInvalidOneOfName                    = errors.New("invalid oneof name")
	errStructContainsDuplicateFieldNumbers = errors.New("struct contains duplicate field numbers")
)

func parse(fs *token.FileSet, f ast.Node) (string, []message, error) {
	var (
		canotoImportName string
		packageName      string
		messages         []message
		err              error
	)
	ast.Inspect(f, func(n ast.Node) bool {
		if err != nil {
			return false
		}

		if f, ok := n.(*ast.File); ok {
			packageName = f.Name.Name
			return true
		}

		if f, ok := n.(*ast.ImportSpec); ok {
			if f.Path.Value != canotoImport {
				return false
			}
			if f.Name == nil {
				canotoImportName = defaultCanotoSelector
				return false
			}
			canotoImportName = f.Name.Name
			return false
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

		genericPointers := make(map[string]int)
		if ts.TypeParams != nil {
			typesToIndex := make(map[string]int)
			for _, field := range ts.TypeParams.List {
				for _, name := range field.Names {
					typesToIndex[name.Name] = message.numTypes
					message.numTypes++
				}
			}

			var currentTypeNumber int
			for _, field := range ts.TypeParams.List {
				currentTypeNumber += len(field.Names)

				t, ok := field.Type.(*ast.IndexExpr)
				if !ok {
					continue
				}

				var typeName string
				if canotoImportName == "." {
					x, ok := t.X.(*ast.Ident)
					if !ok {
						continue
					}
					typeName = x.Name
				} else {
					x, ok := t.X.(*ast.SelectorExpr)
					if !ok {
						continue
					}
					if ident, ok := x.X.(*ast.Ident); !ok || ident.Name != canotoImportName {
						continue
					}
					typeName = x.Sel.Name
				}
				if typeName != "FieldPointer" {
					continue
				}

				ident, ok := t.Index.(*ast.Ident)
				if !ok {
					continue
				}
				// Make sure the type is generic
				if _, ok := typesToIndex[ident.Name]; !ok {
					continue
				}

				genericPointers[ident.Name] = currentTypeNumber
			}
		}
		for _, sf := range st.Fields.List {
			var (
				field  field
				hasTag bool
			)
			field, hasTag, err = parseField(
				fs,
				message.canonicalizedName,
				genericPointers,
				sf,
			)
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

func parseField(
	fs *token.FileSet,
	canonicalizedStructName string,
	genericTypes map[string]int,
	af *ast.Field,
) (field, bool, error) {
	canotoType, fieldNumber, oneOfName, hasTag, err := parseFieldTag(fs, af)
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

	var (
		unmarshalOneOf  string
		sizeOneOf       string
		sizeOneOfIndent string
	)
	if oneOfName != "" {
		assignOneOf := fmt.Sprintf("c.canotoData.%sOneOf = %d", oneOfName, fieldNumber)
		unmarshalOneOf = fmt.Sprintf(`
			if c.canotoData.%sOneOf != 0 {
				return canoto.ErrDuplicateOneOf
			}
			%s`,
			oneOfName,
			assignOneOf,
		)
		sizeOneOf = "\n\t\t" + assignOneOf
		sizeOneOfIndent = "\n\t\t\t" + assignOneOf
	}

	var (
		t      = af.Type
		goType string
	)
	for {
		switch tt := t.(type) {
		case *ast.Ident:
			goType = tt.Name
		case *ast.SelectorExpr:
		case *ast.StarExpr:
			t = tt.X
			continue
		case *ast.ArrayType:
			t = tt.Elt
			continue
		case *ast.IndexExpr:
			t = tt.X
			continue
		case *ast.IndexListExpr:
			t = tt.X
			continue
		default:
			return field{}, false, fmt.Errorf("%w %T at %s",
				errInvalidGoType,
				t,
				fs.Position(t.Pos()),
			)
		}
		break
	}

	var genericTypeCast string
	if genericType, ok := genericTypes[goType]; ok {
		genericTypeCast = fmt.Sprintf("T%d", genericType)
	}

	name := af.Names[0].Name
	canonicalizedName := canonicalizeName(name)
	protoType := canotoType.ProtoType(goType)
	return field{
		name:              name,
		canonicalizedName: canonicalizedName,
		goType:            goType,
		protoType:         protoType,
		canotoType:        canotoType,
		fieldNumber:       fieldNumber,
		oneOfName:         oneOfName,
		templateArgs: map[string]string{
			"escapedStructName": canonicalizedStructName,
			"fieldNumber":       strconv.FormatUint(uint64(fieldNumber), 10),
			"wireType":          canotoType.WireType().String(),
			"goType":            goType,
			"genericTypeCast":   genericTypeCast,
			"protoType":         protoType,
			"protoTypePrefix":   canotoType.ProtoTypePrefix(),
			"protoTypeSuffix":   canotoType.ProtoTypeSuffix(),
			"fieldName":         name,
			"escapedFieldName":  canonicalizedName,
			"suffix":            canotoType.Suffix(),
			"oneOf":             oneOfName,
			"unmarshalOneOf":    unmarshalOneOf,
			"sizeOneOf":         sizeOneOf,
			"sizeOneOfIndent":   sizeOneOfIndent,
		},
	}, true, nil
}

// canonicalizeName replaces "_" with "_1" to avoid collisions with "__" which
// is used as a reserved separator.
func canonicalizeName(name string) string {
	return strings.ReplaceAll(name, "_", "_1")
}

// parseFieldTag parses the tag of the provided field and returns the canoto
// description, if one exists.
func parseFieldTag(fs *token.FileSet, field *ast.Field) (
	canotoType,
	uint32,
	string,
	bool,
	error,
) {
	if field.Tag == nil {
		return "", 0, "", false, nil
	}

	rawTag := strings.Trim(field.Tag.Value, "`")
	tags, err := structtag.Parse(rawTag)
	if err != nil {
		return "", 0, "", false, err
	}

	tag, err := tags.Get(canotoTag)
	if err != nil {
		return "", 0, "", false, nil //nolint: nilerr // errors imply the tag was not found
	}

	fieldType := canotoType(tag.Name)
	if !fieldType.IsValid() {
		return "", 0, "", false, fmt.Errorf("%w %q at %s",
			errUnexpectedCanotoType,
			tag.Name,
			fs.Position(field.Pos()),
		)
	}

	if len(tag.Options) > 2 {
		return "", 0, "", false, fmt.Errorf("%w %q at %s",
			errMalformedTag,
			tag.Value(),
			fs.Position(field.Pos()),
		)
	}

	fieldNumber, err := strconv.ParseUint(tag.Options[0], 10, 32)
	if err != nil {
		return "", 0, "", false, fmt.Errorf("%w at %s",
			err,
			fs.Position(field.Pos()),
		)
	}

	var oneof string
	if len(tag.Options) == 2 {
		oneof = tag.Options[1]
		if !oneOfRegex.MatchString(oneof) {
			return "", 0, "", false, fmt.Errorf("%w %q at %s",
				errInvalidOneOfName,
				oneof,
				fs.Position(field.Pos()),
			)
		}
	}
	return fieldType, uint32(fieldNumber), oneof, true, nil
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
