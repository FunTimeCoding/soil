package service

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func loadPackages(
	directory string,
	patterns ...string,
) ([]*packages.Package, *token.FileSet, error) {
	all, set, e := resolve.LoadPackages(directory, patterns...)

	if e != nil {
		return nil, nil, e
	}

	return resolve.PreferTestVariants(all), set, nil
}
