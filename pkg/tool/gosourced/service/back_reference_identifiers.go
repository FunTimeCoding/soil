package service

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

// backReferenceIdentifiers returns every identifier inside node that resolves to a
// top-level symbol of package p and is not itself moving (excluded). These are
// the references that, once the declaration lands in the target package, must
// be qualified as source.Name and the source package imported.
func backReferenceIdentifiers(
	p *packages.Package,
	excluded map[token.Pos]bool,
	node ast.Node,
) []*ast.Ident {
	var result []*ast.Ident
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

			result = append(result, ident)

			return true
		},
	)

	return result
}
