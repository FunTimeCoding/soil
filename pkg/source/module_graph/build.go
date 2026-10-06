package module_graph

import (
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"go/build"
)

func Build(root string) *Graph {
	fileTags := build_tag.Files(root)
	c := build.Default
	c.BuildTags = build_tag.Union(fileTags)
	module, requirements, replaced := readModule(root)
	nodes, units := mainNodes(&c, root, module)
	reachReplaced(&c, nodes, units, replaced)

	return NewGraph(nodes, units, fileTags, requirements)
}
