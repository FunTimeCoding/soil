package module_graph

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/build"
	"slices"
)

func importDirectory(
	c *build.Context,
	directory string,
	path string,
) (*Node, []*Node) {
	p, e := c.ImportDir(directory, 0)

	if p == nil || e != nil && len(p.TestGoFiles)+len(p.XTestGoFiles) == 0 {
		return nil, nil
	}

	base := inDirectory(directory, slices.Concat(p.GoFiles, p.CgoFiles))
	var node *Node

	if len(base) > 0 {
		node = NewNode(path, directory, base, p.Imports)
	}

	var units []*Node
	internal := slices.Concat(base, inDirectory(directory, p.TestGoFiles))

	if len(internal) > 0 {
		units = append(
			units,
			NewNode(
				path,
				directory,
				internal,
				slices.Concat(p.Imports, p.TestImports),
			),
		)
	}

	if len(p.XTestGoFiles) > 0 {
		units = append(
			units,
			NewNode(
				join.Empty(path, "_test"),
				directory,
				inDirectory(directory, p.XTestGoFiles),
				p.XTestImports,
			),
		)
	}

	return node, units
}
