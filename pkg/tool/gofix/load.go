package gofix

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
)

func load(
	directory string,
	patterns []string,
) ([]*packages.Package, *token.FileSet) {
	result, set, e := resolve.LoadPackages(directory, patterns...)

	if e != nil {
		errors.Printf("load: %s\n", e)
		os.Exit(1)
	}

	return result, set
}
