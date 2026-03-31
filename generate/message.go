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
	template          templates
}

type templates struct {
	cache      string
	number     string
	tag        string
	oneOfType  string
	oneOfUnset string
	oneOfField string
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
