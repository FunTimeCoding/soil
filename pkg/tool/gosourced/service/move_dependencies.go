package service

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"sort"
)

func moveDependencies(
	p *packages.Package,
	excluded map[token.Pos]bool,
	node ast.Node,
) []string {
	seen := make(map[string]bool)
	var result []string
	ast.Inspect(
		node,
		func(n ast.Node) bool {
			ident, okay := n.(*ast.Ident)

			if !okay {
				return true
			}

			o := p.TypesInfo.Uses[ident]

			if o == nil || o.Pkg() != p.Types {
				return true
			}

			if o.Parent() != p.Types.Scope() || excluded[o.Pos()] {
				return true
			}

			if seen[o.Name()] {
				return true
			}

			seen[o.Name()] = true
			result = append(result, o.Name())

			return true
		},
	)
	sort.Strings(result)

	return result
}
