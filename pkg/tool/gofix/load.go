package gofix

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func Load(
	directory string,
	patterns []string,
) ([]*packages.Package, *token.FileSet) {
	result, set, e := resolve.LoadPackages(directory, patterns...)
	resolve.ExitOnLoadFailure(e)

	return result, set
}
