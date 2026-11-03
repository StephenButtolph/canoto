package generate

import (
	"maps"
	"slices"
	"strings"
)

type message struct {
	name              string
	canonicalizedName string
	numTypes          int
	fields            []field
	template          Templates
}

// Templates used to format identifiers in generated code. Empty templates use
// their defaults.
//
// Templates may reference {struct}, {oneOf}, and {field}, depending on what the
// identifier is generated for. The c-prefixed variants {cStruct}, {cOneOf}, and
// {cField} replace "_" with "_1", so "__" can be used as a separator.
//
// For example, the Number template "canotoNumber_{cStruct}__{cField}" names
// field My_Field of struct My_Struct "canotoNumber_My_1Struct__My_1Field".
type Templates struct {
	// Cache names the type of the canotoData field.
	Cache string
	// Number names the field number constants.
	Number string
	// Tag names the field tag constants.
	Tag string
	// OneOfType names the oneOf enum types.
	OneOfType string
	// OneOfUnset names the unset oneOf constants.
	OneOfUnset string
	// OneOfField names the oneOf field constants.
	OneOfField string
}

func (t *Templates) init() {
	if t.Cache == "" {
		t.Cache = "canotoData_{struct}"
	}
	if t.Number == "" {
		t.Number = "canotoNumber_{cStruct}__{cField}"
	}
	if t.Tag == "" {
		t.Tag = "canotoTag_{cStruct}__{cField}"
	}
	if t.OneOfType == "" {
		t.OneOfType = "canotoOneOfType_{cStruct}__{cOneOf}"
	}
	if t.OneOfUnset == "" {
		t.OneOfUnset = "canotoOneOfUnset_{cStruct}__{cOneOf}"
	}
	if t.OneOfField == "" {
		t.OneOfField = "canotoOneOf_{cStruct}__{cField}"
	}

	fields := []*string{&t.Cache, &t.Number, &t.Tag, &t.OneOfType, &t.OneOfUnset, &t.OneOfField}
	for _, s := range fields {
		*s = strings.ReplaceAll(*s, "{", "${")
	}
}

func (m *message) OneOfs() []string {
	oneOfs := make(map[string]struct{})
	for _, f := range m.fields {
		if f.oneOfName != "" {
			oneOfs[f.oneOfName] = struct{}{}
		}
	}

	return slices.Sorted(maps.Keys(oneOfs))
}
