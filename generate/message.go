package generate

import (
	"maps"
	"slices"
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

func (m *message) OneOfs() []string {
	oneOfs := make(map[string]struct{})
	for _, f := range m.fields {
		if f.oneOfName != "" {
			oneOfs[f.oneOfName] = struct{}{}
		}
	}

	return slices.Sorted(maps.Keys(oneOfs))
}
