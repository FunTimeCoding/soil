package relocation

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

type Plan struct {
	Set               *token.FileSet
	All               []*packages.Package
	Source            *packages.Package
	Target            *packages.Package
	Resolver          *resolve.Names
	Entries           []*Entry
	Flipped           []types.Object
	Constraints       map[string][]string
	Qualifications    map[string]*FileQualification
	Renames           map[*ast.Ident]string
	Holders           map[*ast.SelectorExpr]string
	PackagePath       string
	TargetPackagePath string
	TargetPackageName string
	MoveDirectory     string
	CreateTarget      bool
}
