package load

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func Lenient(
	directory string,
	patterns ...string,
) ([]*packages.Package, []*packages.Package, *token.FileSet, error) {
	all, set, e := resolve.LoadPackagesUnchecked(directory, nil, patterns...)

	if e != nil {
		return nil, nil, nil, e
	}

	var loaded []*packages.Package
	var broken []*packages.Package

	for _, p := range all {
		if len(p.Errors) == 0 {
			loaded = append(loaded, p)

			continue
		}

		broken = append(broken, p)
	}

	return resolve.PreferTestVariants(loaded), broken, set, nil
}
