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

	"github.com/StephenButtolph/canoto"
)

const (
	canotoTag = "canoto"
	goBytes   = "[]byte"

	// By default, enable concurrent serialization
	defaultConcurrent = true
)

var (
	// oneOfRegex is used to match a string that consists only of letters (both
	// uppercase and lowercase), digits, and underscores from start to end.
	//
	// \A asserts the position at the start of the string.
	// [a-zA-Z0-9_] matches any letter (both uppercase and lowercase), digit, or
	// underscore.
	// + matches one or more of the preceding token.
	// \z asserts the position at the end of the string.
	oneOfRegex = regexp.MustCompile(`\A[a-zA-Z0-9_]+\z`)

	errUnexpectedNumberOfIdentifiers       = errors.New("unexpected number of identifiers")
	errInvalidGoType                       = errors.New("invalid Go type")
	errMalformedCanotoDataTag              = errors.New(`expected "noatomic" got`)
	errMalformedTag                        = errors.New(`expected "type,fieldNumber[,oneof]" got`)
	errInvalidFieldNumber                  = errors.New("invalid field number")
	errRepeatedOneOf                       = errors.New("oneof must not be repeated")
	errInvalidOneOfName                    = errors.New("invalid oneof name")
	errStructContainsDuplicateFieldNumbers = errors.New("struct contains duplicate field numbers")
)

func parse(
	fs *token.FileSet,
	f ast.Node,
	canotoImport string,
	internal bool,
) (string, []message, error) {
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
				if canotoImportName == "." || internal {
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
		var useAtomic bool
		useAtomic, err = parseCanotoData(fs, st.Fields)
		if err != nil {
			return false
		}
		message.useAtomic = useAtomic

		for _, sf := range st.Fields.List {
			if len(sf.Names) >= 1 && sf.Names[0].Name == "canotoData" {
				continue
			}

			var (
				field  field
				hasTag bool
			)
			field, hasTag, err = parseField(
				fs,
				message.canonicalizedName,
				useAtomic,
				internal,
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

func parseCanotoData(
	fs *token.FileSet,
	afl *ast.FieldList,
) (bool, error) {
	for _, sf := range afl.List {
		if len(sf.Names) != 1 {
			continue
		}
		if sf.Names[0].Name != "canotoData" {
			continue
		}

		if sf.Tag == nil {
			return defaultConcurrent, nil
		}

		rawTag := strings.Trim(sf.Tag.Value, "`")
		tags, err := structtag.Parse(rawTag)
		if err != nil {
			return false, err
		}

		tag, err := tags.Get(canotoTag)
		if err != nil {
			return defaultConcurrent, nil //nolint: nilerr // errors imply the tag was not found
		}
		if tag.Name != "noatomic" || len(tag.Options) != 0 {
			return false, fmt.Errorf("%w %q at %s",
				errMalformedCanotoDataTag,
				tag.Value(),
				fs.Position(sf.Pos()),
			)
		}
		return false, nil
	}
	return defaultConcurrent, nil
}

func parseField(
	fs *token.FileSet,
	canonicalizedStructName string,
	useAtomic bool,
	internal bool,
	genericTypes map[string]int,
	af *ast.Field,
) (field, bool, error) {
	canotoType, fieldNumber, oneOfName, hasTag, err := parseFieldTag(fs, af)
	if err != nil || !hasTag {
		return field{}, false, err
	}

	var (
		load        string
		storePrefix = ` = `
		storeSuffix string
	)
	if useAtomic {
		load = ".Load()"
		storePrefix = ".Store(int64("
		storeSuffix = "))"
	}

	var selector string
	if !internal {
		selector = defaultCanotoImporter
	}

	var (
		unmarshalOneOf  string
		sizeOneOf       string
		sizeOneOfIndent string
	)
	if oneOfName != "" {
		if useAtomic {
			unmarshalOneOf = fmt.Sprintf(`
			if c.canotoData.%sOneOf.Swap(%d) != 0 {
				return %sErrDuplicateOneOf
			}`,
				oneOfName,
				fieldNumber,
				selector,
			)
		} else {
			unmarshalOneOf = fmt.Sprintf(`
			if c.canotoData.%sOneOf != 0 {
				return %sErrDuplicateOneOf
			}
			c.canotoData.%sOneOf = %d`,
				oneOfName,
				selector,
				oneOfName,
				fieldNumber,
			)
		}

		assignOneOf := fmt.Sprintf("%sOneOf = %d", oneOfName, fieldNumber)
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

	var name string
	switch len(af.Names) {
	case 0:
		name = goType
	case 1:
		name = af.Names[0].Name
	default:
		return field{}, false, fmt.Errorf("%w wanted <= 1 but got %d at %s",
			errUnexpectedNumberOfIdentifiers,
			len(af.Names),
			fs.Position(af.Pos()),
		)
	}

	var genericTypeCast string
	if genericType, ok := genericTypes[goType]; ok {
		genericTypeCast = fmt.Sprintf("T%d", genericType)
	}

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
			"selector":          selector,
			"load":              load,
			"storePrefix":       storePrefix,
			"storeSuffix":       storeSuffix,
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

	if len(tag.Options) > 2 || len(tag.Options) < 1 {
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
	if fieldNumber == 0 {
		return "", 0, "", false, fmt.Errorf("%w 0 at %s",
			errInvalidFieldNumber,
			fs.Position(field.Pos()),
		)
	}
	if fieldNumber > canoto.MaxFieldNumber {
		return "", 0, "", false, fmt.Errorf("%w %d exceeds maximum value of %d at %s",
			errInvalidFieldNumber,
			fieldNumber,
			canoto.MaxFieldNumber,
			fs.Position(field.Pos()),
		)
	}

	var oneof string
	if len(tag.Options) == 2 {
		if fieldType.IsRepeated() {
			return "", 0, "", false, fmt.Errorf("%w at %s",
				errRepeatedOneOf,
				fs.Position(field.Pos()),
			)
		}

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
