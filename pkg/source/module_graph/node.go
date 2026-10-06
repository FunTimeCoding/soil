package module_graph

type Node struct {
	Path      string
	Directory string
	Files     []string
	Imports   []string
}
