package xref

func New(
	units map[string]*References,
	directories map[string]string,
) *Index {
	return &Index{units: units, directories: directories}
}
