package symbol

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"go/ast"
	"go/parser"
	"golang.org/x/tools/go/packages"
	"sort"
)

func extendGraph(
	directory string,
	all []*packages.Package,
	sources []string,
	imports []string,
) []*packages.Package {
	names := lookup.PackageNames(all)
	index := packageIndex(all)
	wanted := map[string]bool{}

	for _, path := range imports {
		if _, loaded := index[path]; !loaded {
			wanted[path] = true
		}
	}

	for _, source := range sources {
		expression, e := parser.ParseExpr(source)

		if e != nil {
			continue
		}

		ast.Inspect(
			expression,
			func(n ast.Node) bool {
				s, okay := n.(*ast.SelectorExpr)

				if !okay {
					return true
				}

				if x, plain := s.X.(*ast.Ident); plain && len(names[x.Name]) == 0 {
					wanted[x.Name] = true
				}

				return true
			},
		)
	}

	var paths []string

	for path := range wanted {
		paths = append(paths, path)
	}

	sort.Strings(paths)
	result := all

	for _, path := range paths {
		loaded, _, e := resolve.LoadPackagesUnchecked(directory, nil, path)

		if e != nil {
			continue
		}

		for _, p := range loaded {
			if p.Types != nil && len(p.Errors) == 0 && p.PkgPath == path {
				result = append(result, p)
			}
		}
	}

	return result
}
