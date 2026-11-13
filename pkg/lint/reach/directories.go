package reach

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/go_mod"
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"github.com/funtimecoding/soil/pkg/source/warm_forest"
	"path"
	"path/filepath"
	"strings"
)

func Directories(
	trees *warm_forest.Forest,
	root string,
	files []string,
) *Reach {
	module := go_mod.ReadDirectory(root).Module.Mod.Path
	packages := map[string]bool{}
	documents := map[string]bool{}
	var packagePaths []string
	var g *module_graph.Graph

	for _, f := range files {
		directory := filepath.Dir(f)
		relative, okay := rootRelative(root, directory)

		if !okay {
			continue
		}

		if !strings.HasSuffix(f, soil.GoExtension) {
			documents[relative] = true

			continue
		}

		if hasGoFiles(directory) {
			packages[relative] = true
		}

		packagePaths = append(
			packagePaths,
			path.Join(module, filepath.ToSlash(relative)),
		)
	}

	if len(packagePaths) > 0 {
		g = module_graph.BuildForest(trees, root)

		for _, p := range packagePaths {
			for _, n := range module_graph.Dependents(g, p) {
				if relative, okay := rootRelative(root, n.Directory); okay {
					packages[relative] = true
				}
			}
		}
	}

	return NewReach(patterns(packages), patterns(documents), g)
}
