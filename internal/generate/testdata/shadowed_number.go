package testdata

// shadowedNumber is generated with the number template "{field}", so the field
// number constant of field is shadowed by the field variable returned from
// ReadTag in UnmarshalCanotoFrom.
type shadowedNumber struct {
	field int64 `canoto:"int,1"`
	Other int64 `canoto:"int,2"`

	canotoData canotoData_shadowedNumber
}
