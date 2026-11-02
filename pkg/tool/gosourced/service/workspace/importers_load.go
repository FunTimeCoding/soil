package workspace

import (
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"go/token"
	"golang.org/x/tools/go/packages"
	"slices"
	"strings"
)

func (s *Service) ImportersLoad(
	directory string,
	packagePath string,
	subtree bool,
) ([]*packages.Package, *token.FileSet, error) {
	g := s.Workspace(directory).Graph()
	reach := []string{packagePath}

	if subtree {
		reach = []string{
			join.Empty(packagePath, go_modConstant.RecursivePattern),
		}
	}

	if s.full {
		return load.Around(directory, load.ReachedPaths(g, reach), "./...")
	}

	prefix := join.Empty(packagePath, "/")
	inside := func(p string) bool {
		return p == packagePath || subtree && strings.HasPrefix(p, prefix)
	}
	units := g.Units
	var patterns []string

	for _, u := range units {
		if slices.ContainsFunc(u.Imports, inside) {
			patterns = append(patterns, u.Directory)
		}
	}

	if len(patterns) == 0 {
		return nil, token.NewFileSet(), nil
	}

	patterns = slices.Compact(slices.Sorted(slices.Values(patterns)))

	if 2*len(patterns) > len(units) {
		return load.Around(directory, load.ReachedPaths(g, reach), "./...")
	}

	return load.Packages(directory, patterns...)
}
