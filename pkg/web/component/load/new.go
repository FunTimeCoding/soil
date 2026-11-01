package load

func New(
	path string,
	label string,
) *Load {
	return &Load{path: path, label: label}
}
