package addition

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type Call struct {
	Target     *Target
	Owner      *packages.Package
	File       *ast.File
	Call       *ast.CallExpr
	Qualifiers map[string]string
}
