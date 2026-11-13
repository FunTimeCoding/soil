package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
)

func load(
	directory string,
	reportedPatterns []string,
	loadedPatterns []string,
) ([]*packages.Package, map[string]bool) {
	reported, e := resolve.ListPackages(directory, reportedPatterns...)
	errors.PanicOnError(e)
	result, _, e := resolve.LoadPackages(directory, loadedPatterns...)
	resolve.ExitOnLoadFailure(e)

	return result, reported
}
