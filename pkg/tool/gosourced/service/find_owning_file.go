package service

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func findOwningFile(
	all []*packages.Package,
	position token.Pos,
) (*packages.Package, *ast.File) {
	for _, p := range all {
		for _, file := range p.Syntax {
			if file.Pos() <= position && position <= file.End() {
				return p, file
			}
		}
	}

	return nil, nil
}
