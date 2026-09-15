package service

import (
	"go/ast"
	"go/token"
)

func countIdentities(file *ast.File) int {
	var count int

	for _, d := range file.Decls {
		switch t := d.(type) {
		case *ast.FuncDecl:
			count++
		case *ast.GenDecl:
			if t.Tok == token.TYPE {
				count += len(t.Specs)
			}
		}
	}

	return count
}
