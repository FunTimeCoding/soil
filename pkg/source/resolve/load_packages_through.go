package resolve

import (
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func LoadPackagesThrough(
	directory string,
	overlay map[string][]byte,
	patterns ...string,
) ([]*packages.Package, *token.FileSet, error) {
	result, set, e := LoadPackagesUnchecked(directory, overlay, patterns...)

	if e != nil {
		return nil, nil, e
	}

	for _, p := range result {
		for _, f := range p.Errors {
			return nil, nil, load_failure.New(p.PkgPath, f)
		}
	}

	return result, set, nil
}
