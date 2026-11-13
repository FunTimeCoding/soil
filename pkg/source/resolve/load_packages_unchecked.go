package resolve

import (
	"github.com/funtimecoding/soil/pkg/source/build_tag"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func LoadPackagesUnchecked(
	directory string,
	overlay map[string][]byte,
	patterns ...string,
) ([]*packages.Package, *token.FileSet, error) {
	return LoadPackagesTagged(
		directory,
		build_tag.Discover(directory),
		overlay,
		patterns...,
	)
}
