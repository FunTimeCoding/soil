package service

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func sourceReferencesMoved(
	p *packages.Package,
	entries []*moveEntry,
	moved map[types.Object]bool,
) bool {
	for _, file := range p.Syntax {
		reference := false
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				ident, okay := n.(*ast.Ident)

				if !okay {
					return true
				}

				o := p.TypesInfo.Uses[ident]

				if o == nil || !moved[o] {
					return true
				}

				if insideMoved(entries, ident.Pos()) {
					return true
				}

				reference = true

				return false
			},
		)

		if reference {
			return true
		}
	}

	return false
}
