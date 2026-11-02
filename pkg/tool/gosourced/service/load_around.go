package service

import (
	"github.com/funtimecoding/soil/pkg/errors/load_failure"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strings"
)

func loadAround(
	directory string,
	reached map[string]bool,
	patterns ...string,
) ([]*packages.Package, *token.FileSet, error) {
	all, set, e := resolve.LoadPackagesUnchecked(directory, nil, patterns...)

	if e != nil {
		return nil, nil, e
	}

	var result []*packages.Package

	for _, p := range all {
		if len(p.Errors) == 0 {
			result = append(result, p)

			continue
		}

		if reached[p.PkgPath] ||
			reached[strings.TrimSuffix(p.PkgPath, constant.TestClauseSuffix)] {
			return nil, nil, load_failure.New(p.PkgPath, p.Errors[0])
		}
	}

	return resolve.PreferTestVariants(result), set, nil
}
