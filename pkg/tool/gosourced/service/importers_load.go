package service

import (
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
	if s.full {
		return loadPackages(directory, "./...")
	}

	prefix := join.Empty(packagePath, "/")
	inside := func(p string) bool {
		return p == packagePath || subtree && strings.HasPrefix(p, prefix)
	}
	units := s.workspace(directory).Graph().Units
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
		return loadPackages(directory, "./...")
	}

	return loadPackages(directory, patterns...)
}
