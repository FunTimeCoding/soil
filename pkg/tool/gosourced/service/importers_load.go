package service

import (
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/token"
	"golang.org/x/tools/go/packages"
	"slices"
	"strings"
)

func (s *Service) importersLoad(
	directory string,
	packagePath string,
	subtree bool,
) ([]*packages.Package, *token.FileSet, error) {
	g := s.workspace(directory).Graph()
	reach := []string{packagePath}

	if subtree {
		reach = []string{join.Empty(packagePath, constant.RecursivePattern)}
	}

	if s.full {
		return loadAround(directory, reachedPaths(g, reach), "./...")
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
		return loadAround(directory, reachedPaths(g, reach), "./...")
	}

	return loadPackages(directory, patterns...)
}
