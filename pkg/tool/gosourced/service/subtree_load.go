package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strings"
)

func (s *Service) subtreeLoad(
	directory string,
	packagePath string,
) ([]*packages.Package, *token.FileSet, error) {
	if s.full {
		return loadPackages(directory, "./...")
	}

	prefix := join.Empty(packagePath, "/")
	inside := func(p string) bool {
		return p == packagePath || strings.HasPrefix(p, prefix) ||
			p == join.Empty(packagePath, "_test")
	}
	var patterns []string
	own := false

	for path, u := range s.workspace(directory).Graph().Units {
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
		return loadPackages(directory, "./...")
	}

	return loadPackages(directory, patterns...)
}
