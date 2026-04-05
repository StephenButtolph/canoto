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
	noCopy            bool
	template          Templates
}

// Templates controls the naming patterns used in generated code.
type Templates struct {
	Cache      string
	Number     string
	Tag        string
	OneOfType  string
	OneOfUnset string
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
