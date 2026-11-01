package action

func New(
	path string,
	label string,
) *Action {
	return &Action{path: path, label: label}
}
