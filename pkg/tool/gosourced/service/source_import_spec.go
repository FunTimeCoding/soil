package service

import (
	"go/ast"
	"go/token"
	"strconv"
)

func sourceImportSpec(packagePath string) *ast.ImportSpec {
	return &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote(packagePath),
		},
	}
}
