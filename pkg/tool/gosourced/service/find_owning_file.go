package service

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func findOwningFile(
	all []*packages.Package,
	set *token.FileSet,
	position token.Pos,
) (*packages.Package, *ast.File) {
	filename := set.Position(position).Filename

	for _, p := range all {
		if file := findSyntaxFile(set, p, filename); file != nil {
			return p, file
		}
	}

	return nil, nil
}
