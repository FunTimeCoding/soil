package gocheck

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gocheck/constant"
	"golang.org/x/tools/go/packages"
)

func loadFailures(
	root string,
	loaded []*packages.Package,
) *output.Results {
	result := output.NewResultsWithDirectory(root)
	seen := map[string]bool{}

	for _, p := range loaded {
		if len(p.Errors) == 0 || seen[p.PkgPath] {
			continue
		}

		seen[p.PkgPath] = true
		f := positioned(p.Errors)
		path := f.Pos

		if path == "" {
			path = p.PkgPath
		}

		result.AddConcern(
			concern.NewFile(
				constant.LoadFailure,
				fmt.Sprintf("does not load: %s - %s", p.PkgPath, f.Msg),
				path,
				false,
			),
		)
	}

	return result
}
