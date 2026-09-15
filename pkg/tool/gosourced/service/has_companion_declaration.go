package service

import (
	"go/ast"
	"go/token"
)

func hasCompanionDeclaration(
	file *ast.File,
	index int,
) bool {
	for i, d := range file.Decls {
		if i == index {
			continue
		}

		if g, okay := d.(*ast.GenDecl); okay && g.Tok == token.IMPORT {
			continue
		}

		return true
	}

	return false
}
