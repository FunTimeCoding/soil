package service

import (
	"go/ast"
	"go/token"
)

type movePlan struct {
	set               *token.FileSet
	entries           []*moveEntry
	qualifications    map[string]*fileQualification
	renames           map[*ast.Ident]string
	packagePath       string
	targetPackagePath string
	targetPackageName string
	moveDirectory     string
	createTarget      bool
	sourceLocalName   string
}
