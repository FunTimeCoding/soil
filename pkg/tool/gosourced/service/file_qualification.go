package service

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type fileQualification struct {
	file        *ast.File
	owner       *packages.Package
	samePackage bool
	name        *importName
	idents      map[*ast.Ident]string
	positions   []qualifiedPosition
}
