package goanalyze

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
)

func load(patterns []string) []*packages.Package {
	result, _, e := resolve.LoadPackages("", patterns...)
	errors.PanicOnError(e)

	return result
}
