package service

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

type movePlan struct {
	set               *token.FileSet
	all               []*packages.Package
	source            *packages.Package
	target            *packages.Package
	resolver          *resolve.Names
	entries           []*moveEntry
	qualifications    map[string]*fileQualification
	renames           map[*ast.Ident]string
	packagePath       string
	targetPackagePath string
	targetPackageName string
	moveDirectory     string
	createTarget      bool
}
