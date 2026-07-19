package service

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

// sourceReferencesMoved reports whether any code staying in the source package
// references a moving symbol. When it does and the moved code also references
// back into source, the move would create a source<->target import cycle, so
// back-reference qualification must refuse.
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
