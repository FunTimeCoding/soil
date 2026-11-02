package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
	"io"
)

func loadReached(
	directory string,
	loaded []*packages.Package,
	notice io.Writer,
) []*packages.Package {
	paths := resolve.ReachedPackages(
		loaded,
		resolve.LocalReplacements(directory),
	)

	if len(paths) == 0 {
		return nil
	}

	result, _, e := resolve.LoadPackagesUnchecked(directory, nil, paths...)
	errors.PanicOnError(e)

	return resolve.PreferTestVariants(dropBroken(result, notice))
}
