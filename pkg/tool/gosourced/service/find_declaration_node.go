package service

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func findDeclarationNode(
	p *packages.Package,
	o types.Object,
) (*ast.File, ast.Decl, ast.Spec) {
	for _, file := range p.Syntax {
		if o.Pos() < file.Pos() || o.Pos() > file.End() {
			continue
		}

		for _, d := range file.Decls {
			switch declaration := d.(type) {
			case *ast.FuncDecl:
				if declaration.Name.Pos() == o.Pos() {
					return file, declaration, nil
				}
			case *ast.GenDecl:
				for _, s := range declaration.Specs {
					if specDeclares(s, o) {
						return file, declaration, s
					}
				}
			}
		}
	}

	return nil, nil, nil
}
