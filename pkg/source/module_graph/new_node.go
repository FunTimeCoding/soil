package module_graph

func NewNode(
	path string,
	directory string,
	files []string,
	imports []string,
) *Node {
	return &Node{
		Path:      path,
		Directory: directory,
		Files:     files,
		Imports:   imports,
	}
}
