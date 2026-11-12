package gofix

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func syntaxFiles(
	fileSet *token.FileSet,
	all []*packages.Package,
) map[string]*ast.File {
	result := make(map[string]*ast.File)

	for _, p := range all {
		for _, f := range p.Syntax {
			result[fileSet.File(f.Pos()).Name()] = f
		}
	}

	return result
}
