package module_graph

import "github.com/funtimecoding/soil/pkg/source/build_tag"

func NewGraph(
	nodes map[string]*Node,
	units map[string]*Node,
	fileTags map[string][]string,
	requirements map[string]string,
) *Graph {
	return &Graph{
		Nodes:        nodes,
		Units:        units,
		Order:        order(nodes),
		Tags:         build_tag.Union(fileTags),
		FileTags:     fileTags,
		Requirements: requirements,
	}
}
