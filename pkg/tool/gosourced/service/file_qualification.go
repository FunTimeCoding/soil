package service

import "go/ast"

type fileQualification struct {
	file        *ast.File
	samePackage bool
	name        *importName
	idents      map[*ast.Ident]string
	positions   []qualifiedPosition
}
