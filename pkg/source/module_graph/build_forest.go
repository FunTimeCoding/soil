package module_graph

import (
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"github.com/funtimecoding/soil/pkg/system"
)

func BuildForest(
	f *warm_forest.Forest,
	root string,
) *Graph {
	t := f.Tree(root)
	module, requirements, replaced := readModule(root)

	for _, directory := range replaced {
		if system.DirectoryExists(directory) {
			f.Tree(directory)
		}
	}

	fileTags := build_tag.TreeFiles(t)
	c := forestContext(f)
	c.BuildTags = build_tag.Union(fileTags)
	nodes, units := treeNodes(&c, t, module)
	reachReplaced(&c, nodes, units, replaced)

	return NewGraph(nodes, units, fileTags, requirements)
}
