package service

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func syntaxFileAt(p *packages.Package, position token.Pos) *ast.File {
	for _, f := range p.Syntax {
		if f.FileStart <= position && position < f.FileEnd {
			return f
		}
	}

	return nil
}
