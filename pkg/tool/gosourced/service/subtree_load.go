package service

import (
	go_modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strings"
)

func (s *Service) subtreeLoad(
	directory string,
	packagePath string,
) ([]*packages.Package, *token.FileSet, error) {
	g := s.workspace(directory).Graph()
	reached := reachedPaths(
		g,
		[]string{join.Empty(packagePath, go_modConstant.RecursivePattern)},
	)

	if s.full {
		return loadAround(directory, reached, "./...")
	}

	prefix := join.Empty(packagePath, "/")
	inside := func(p string) bool {
		return p == packagePath || strings.HasPrefix(p, prefix) ||
			p == join.Empty(packagePath, constant.TestClauseSuffix)
	}
	var patterns []string
	own := false

	for path, u := range g.Units {
		if inside(path) {
			own = true
			patterns = append(patterns, u.Directory)

			continue
		}

		for _, i := range u.Imports {
			if inside(i) {
				patterns = append(patterns, u.Directory)

				break
			}
		}
	}

	if !own {
		return loadAround(directory, reached, "./...")
	}

	return loadPackages(directory, patterns...)
}
