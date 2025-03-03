package generate

import (
	"slices"

	"golang.org/x/exp/maps"
)

type message struct {
	name              string
	canonicalizedName string
	numTypes          int
	fields            []field
	useAtomic         bool
}

func (m *message) OneOfs() []string {
	oneOfs := make(map[string]struct{})
	for _, f := range m.fields {
		if f.oneOfName != "" {
			oneOfs[f.oneOfName] = struct{}{}
		}
	}

	oneOfsSlice := maps.Keys(oneOfs)
	slices.Sort(oneOfsSlice)
	return oneOfsSlice
}
