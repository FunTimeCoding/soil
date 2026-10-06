package sink

type operation struct {
	kind    string
	path    string
	target  string
	content []byte
}
