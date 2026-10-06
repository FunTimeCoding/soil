package module_graph

type Graph struct {
	Nodes        map[string]*Node
	Units        map[string]*Node
	Order        []string
	Tags         []string
	FileTags     map[string][]string
	Requirements map[string]string
}
