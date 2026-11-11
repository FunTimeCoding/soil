package workspace

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strings"
)

func (s *Service) SubtreeLoad(
	directory string,
	packagePath string,
) ([]*packages.Package, *token.FileSet, error) {
	g := s.Workspace(directory).Graph()
	reached := load.ReachedPaths(
		g,
		[]string{join.Empty(packagePath, goModConstant.RecursivePattern)},
	)

	if s.full {
		return load.Around(directory, reached, "./...")
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
		return load.Around(directory, reached, "./...")
	}

	return load.Packages(directory, patterns...)
}
