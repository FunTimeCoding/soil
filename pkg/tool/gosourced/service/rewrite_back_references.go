package service

import (
	"go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

// rewriteBackReferences turns each bare back-reference identifier in the moved
// declaration into a source.Name selector, so it resolves once the declaration
// lands in the target package.
func rewriteBackReferences(
	entry *moveEntry,
	sourceName string,
) {
	set := make(map[*ast.Ident]bool)

	for _, ident := range entry.backIdentifiers {
		set[ident] = true
	}

	astutil.Apply(
		entry.node,
		func(c *astutil.Cursor) bool {
			ident, okay := c.Node().(*ast.Ident)

			if !okay || !set[ident] {
				return true
			}

			c.Replace(
				&ast.SelectorExpr{
					X: &ast.Ident{
						Name:    sourceName,
						NamePos: ident.NamePos,
					},
					Sel: &ast.Ident{
						Name:    ident.Name,
						NamePos: ident.NamePos,
					},
				},
			)

			return true
		},
		nil,
	)
}
