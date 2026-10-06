package service

import (
	"github.com/funtimecoding/soil/pkg/source/index"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func wholeLoad(
	directory string,
	w *index.Workspace,
	packagePath string,
) ([]*packages.Package, *token.FileSet, error) {
	patterns := []string{"./..."}

	if w.Reaches(packagePath) {
		patterns = append(patterns, packagePath)
	}

	return loadPackages(directory, patterns...)
}
