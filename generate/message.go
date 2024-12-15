package generate

type message struct {
	name              string
	canonicalizedName string
	numTypes          int
	fields            []field
}
