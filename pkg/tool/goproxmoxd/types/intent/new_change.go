package intent

func NewChange(
	name string,
	node string,
	from string,
	to string,
) *Change {
	return &Change{Name: name, Node: node, From: from, To: to}
}
