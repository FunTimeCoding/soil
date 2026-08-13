package imports

import (
	"fmt"
	"go/ast"
	"go/token"
)

func Add(
	file *ast.File,
	importPath string,
	alias string,
) {
	spec := &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("%q", importPath),
		},
	}

	if alias != "" {
		spec.Name = ast.NewIdent(alias)
	}

	for _, d := range file.Decls {
		declaration, okay := d.(*ast.GenDecl)

		if !okay || declaration.Tok != token.IMPORT {
			continue
		}

		declaration.Lparen = 1
		declaration.Specs = append(declaration.Specs, spec)

		return
	}

	declaration := &ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{spec}}
	file.Decls = append([]ast.Decl{declaration}, file.Decls...)
}
