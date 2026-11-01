package module_range

func New(
	name string,
	path string,
	from string,
	to string,
) *Range {
	return &Range{Name: name, Path: path, From: from, To: to}
}
