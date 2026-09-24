package index

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/ast"
	"slices"
	"sort"
)

func ParallelTests(
	root string,
	patterns ...string,
) []string {
	loaded, _, e := resolve.LoadPackages(root, patterns...)
	errors.PanicOnError(e)
	var result []string

	for _, p := range loaded {
		for _, file := range p.Syntax {
			ast.Inspect(
				file,
				func(n ast.Node) bool {
					call, okay := n.(*ast.CallExpr)

					if !okay {
						return true
					}

					selector, okay := call.Fun.(*ast.SelectorExpr)

					if okay &&
						selector.Sel.Name == "Parallel" &&
						!slices.Contains(result, p.PkgPath) {
						result = append(result, p.PkgPath)
					}

					return true
				},
			)
		}
	}

	sort.Strings(result)

	return result
}
