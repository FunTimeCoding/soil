package workspace

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strings"
)

func (s *Service) LenientCensusPackages(
	directory string,
	packagePath string,
) ([]*packages.Package, []*packages.Package, *token.FileSet, error) {
	g := s.Workspace(directory).Graph()
	paths := []string{packagePath}
	patterns, wide := censusPatterns(g, paths, s.full)

	if len(patterns) == 0 {
		return nil, nil, token.NewFileSet(), nil
	}

	loaded, broken, set, e := load.Lenient(directory, patterns...)

	if e != nil || !wide {
		return loaded, broken, set, e
	}

	reached := load.ReachedPaths(g, paths)
	var inReach []*packages.Package

	for _, p := range broken {
		if reached[p.PkgPath] ||
			reached[strings.TrimSuffix(p.PkgPath, constant.TestClauseSuffix)] {
			inReach = append(inReach, p)
		}
	}

	return loaded, inReach, set, nil
}
