package resolve

func (n *Names) Override(
	importPath string,
	name string,
) {
	n.names[importPath] = name
}
