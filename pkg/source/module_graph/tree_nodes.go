package module_graph

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/source/warm_tree"
	"go/build"
	"path"
	"path/filepath"
)

func treeNodes(
	c *build.Context,
	t *warm_tree.Tree,
	module string,
) (map[string]*Node, map[string]*Node) {
	result := make(map[string]*Node)
	units := make(map[string]*Node)
	var visit func(string)
	visit = func(relative string) {
		p := module

		if relative != soil.CurrentDirectory {
			p = path.Join(module, relative)
		}

		n, directoryUnits := importDirectory(
			c,
			filepath.Join(t.Root(), filepath.FromSlash(relative)),
			p,
		)

		if n != nil {
			result[p] = n
		}

		for _, u := range directoryUnits {
			units[u.Path] = u
		}

		for _, name := range t.Children(relative) {
			child := path.Join(relative, name)
			e, _ := t.Stat(child)

			if e.Directory && !excludedName(name) &&
				!t.IsFile(path.Join(child, goModConstant.ModFile)) {
				visit(child)
			}
		}
	}
	visit(soil.CurrentDirectory)

	return result, units
}
