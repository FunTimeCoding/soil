package relocation

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type FileQualification struct {
	File        *ast.File
	Owner       *packages.Package
	SamePackage bool
	Name        *ImportName
	Idents      map[*ast.Ident]string
	Positions   []QualifiedPosition
}
