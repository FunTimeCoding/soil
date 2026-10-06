package service

import (
	"go/ast"
	"go/token"
)

func functionDeclarationAt(
	file *ast.File,
	position token.Pos,
) *ast.FuncDecl {
	if file == nil {
		return nil
	}

	for _, d := range file.Decls {
		if f, okay := d.(*ast.FuncDecl); okay && f.Name.Pos() == position {
			return f
		}
	}

	return nil
}
