package service

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func findSyntaxFile(
	fileSet *token.FileSet,
	p *packages.Package,
	filename string,
) *ast.File {
	for _, file := range p.Syntax {
		if fileSet.Position(file.Pos()).Filename == filename {
			return file
		}
	}

	return nil
}
