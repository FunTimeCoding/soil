package literal

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type Site struct {
	Package *packages.Package
	File    *ast.File
	Path    []ast.Node
	Target  ast.Expr
	Outer   ast.Expr
	Parent  ast.Node
	Fields  []string
	Nested  bool
	Shape   string
}
