package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
)

func loadReached(
	directory string,
	loaded []*packages.Package,
) []*packages.Package {
	paths := resolve.ReachedPackages(
		loaded,
		resolve.LocalReplacements(directory),
	)

	if len(paths) == 0 {
		return nil
	}

	result, _, e := resolve.LoadPackages(directory, paths...)
	errors.PanicOnError(e)

	return resolve.PreferTestVariants(result)
}
