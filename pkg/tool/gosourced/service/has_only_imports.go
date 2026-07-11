package service

import (
	"go/ast"
	"go/token"
)

func hasOnlyImports(file *ast.File) bool {
	for _, d := range file.Decls {
		g, okay := d.(*ast.GenDecl)

		if !okay || g.Tok != token.IMPORT {
			return false
		}
	}

	return true
}
