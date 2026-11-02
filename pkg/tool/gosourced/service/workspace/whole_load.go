package workspace

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func wholeLoad(
	directory string,
	w *index.Workspace,
	packagePaths ...string,
) ([]*packages.Package, *token.FileSet, error) {
	patterns := []string{"./..."}
	var reachable []string

	for _, p := range packagePaths {
		if p == "" {
			continue
		}

		reachable = append(reachable, p)

		if w.Reaches(p) {
			patterns = append(patterns, p)
		}
	}

	return load.Around(
		directory,
		load.ReachedPaths(w.Graph(), reachable),
		patterns...,
	)
}
