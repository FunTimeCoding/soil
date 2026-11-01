package field

func New(
	kind string,
	name string,
	label string,
	value string,
) *Field {
	return &Field{kind: kind, name: name, label: label, value: value}
}
