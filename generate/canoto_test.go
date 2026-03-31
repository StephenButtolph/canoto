package generate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeOneOfCaseTypesUsesFormattedOneOfNameForConsts(t *testing.T) {
	m := message{
		name:              "CustomFormat",
		canonicalizedName: "CustomFormat",
		numberTemplate:    "${struct}${field}Number",
		oneOfTemplate:     "${struct}${oneOf}",
		fields: []field{
			{
				name:              "A1",
				canonicalizedName: "A1",
				oneOfName:         "A",
			},
			{
				name:              "A2",
				canonicalizedName: "A2",
				oneOfName:         "A",
			},
		},
	}

	const want = `// CustomFormatA identifies which field is populated in A.
type CustomFormatA uint32

const (
	CustomFormatA__Unset CustomFormatA = 0
	CustomFormatA__A1    CustomFormatA = CustomFormatA1Number
	CustomFormatA__A2    CustomFormatA = CustomFormatA2Number
)

`

	require.Equal(t, want, makeOneOfCaseTypes(m))
}
