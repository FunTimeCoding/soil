package service

import (
	"go/ast"
	"go/token"
)

func remainingIdentityName(file *ast.File) string {
	for _, d := range file.Decls {
		switch t := d.(type) {
		case *ast.FuncDecl:
			return t.Name.Name
		case *ast.GenDecl:
			if t.Tok == token.TYPE {
				return t.Specs[0].(*ast.TypeSpec).Name.Name
			}
		}
	}

	return ""
}
