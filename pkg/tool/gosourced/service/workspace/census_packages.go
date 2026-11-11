package workspace

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func (s *Service) CensusPackages(
	directory string,
	packagePaths ...string,
) ([]*packages.Package, *token.FileSet, error) {
	g := s.Workspace(directory).Graph()
	patterns, wide := censusPatterns(g, packagePaths, s.full)

	if wide {
		return load.Around(
			directory,
			load.ReachedPaths(g, packagePaths),
			patterns...)
	}

	if len(patterns) == 0 {
		return nil, token.NewFileSet(), nil
	}

	return load.Packages(directory, patterns...)
}
