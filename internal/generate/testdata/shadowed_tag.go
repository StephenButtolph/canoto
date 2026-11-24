package testdata

// shadowedTag is generated with the tag template "{field}", so the tag constant
// of v is shadowed by the v variable used to range over the field when
// calculating the size and when appending.
type shadowedTag struct {
	v []string `canoto:"repeated string,1"`

	canotoData canotoData_shadowedTag
}
