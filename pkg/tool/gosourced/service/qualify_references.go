package service

import (
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/ast"
	"golang.org/x/tools/go/ast/astutil"
)

func qualifyReferences(
	q *fileQualification,
	sourcePackagePath string,
	targetPackagePath string,
) {
	if q.samePackage {
		astutil.Apply(
			q.file,
			func(c *astutil.Cursor) bool {
				ident, okay := c.Node().(*ast.Ident)

				if !okay {
					return true
				}

				newName, moved := q.idents[ident]

				if !moved {
					return true
				}

				c.Replace(
					&ast.SelectorExpr{
						X: &ast.Ident{
							Name:    q.name.local,
							NamePos: ident.NamePos,
						},
						Sel: &ast.Ident{
							Name:    newName,
							NamePos: ident.NamePos,
						},
					},
				)

				return true
			},
			nil,
		)
	} else {
		ast.Inspect(
			q.file,
			func(n ast.Node) bool {
				s, okay := n.(*ast.SelectorExpr)

				if !okay {
					return true
				}

				newName, moved := q.idents[s.Sel]

				if !moved {
					return true
				}

				x, okay := s.X.(*ast.Ident)

				if !okay {
					return true
				}

				x.Name = q.name.local
				s.Sel.Name = newName

				return true
			},
		)
		removeImportIfUnused(q.file, sourcePackagePath)
	}

	if !q.name.imported {
		imports.Add(q.file, targetPackagePath, q.name.alias)
	}
}
