package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
)

func load(
	directory string,
	patterns []string,
) ([]*packages.Package, map[string]bool, []*packages.Package) {
	reported, e := resolve.ListPackages(directory, patterns...)
	errors.PanicOnError(e)
	result, _, e := resolve.LoadPackages(
		directory,
		resolve.WithMainModule(patterns)...,
	)
	errors.PanicOnError(e)
	loaded := resolve.PreferTestVariants(result)

	return loaded, reported, loadReached(directory, loaded)
}
