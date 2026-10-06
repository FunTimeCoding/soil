package service

import (
	"github.com/funtimecoding/soil/pkg/source/module_graph"
	"go/token"
	"golang.org/x/tools/go/packages"
	"slices"
)

func (s *Service) censusPackages(
	directory string,
	packagePaths ...string,
) ([]*packages.Package, *token.FileSet, error) {
	g := s.workspace(directory).Graph()
	whole := []string{"./..."}
	var patterns []string

	for _, packagePath := range packagePaths {
		_, own := g.Units[packagePath]
		_, reached := g.Nodes[packagePath]

		if own || reached {
			patterns = append(patterns, packagePath)
		}

		if reached && !own {
			whole = append(whole, packagePath)
		}

		for _, u := range module_graph.Dependents(g, packagePath) {
			patterns = append(patterns, u.Directory)
		}
	}

	if s.full {
		return loadPackages(directory, whole...)
	}

	patterns = slices.Compact(slices.Sorted(slices.Values(patterns)))

	if 2*len(patterns) > len(g.Units) {
		return loadPackages(directory, whole...)
	}

	if len(patterns) == 0 {
		return nil, token.NewFileSet(), nil
	}

	return loadPackages(directory, patterns...)
}
