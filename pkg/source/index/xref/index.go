package xref

type Index struct {
	units       map[string]*References
	directories map[string]string
}
