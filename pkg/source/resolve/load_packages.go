package resolve

import (
	"go/token"
	"golang.org/x/tools/go/packages"
)

func LoadPackages(
	directory string,
	patterns ...string,
) ([]*packages.Package, *token.FileSet, error) {
	return LoadPackagesThrough(directory, nil, patterns...)
}
