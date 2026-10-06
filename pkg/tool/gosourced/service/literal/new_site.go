package literal

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func NewSite(
	p *packages.Package,
	file *ast.File,
	path []ast.Node,
	target ast.Expr,
	outer ast.Expr,
	parent ast.Node,
) *Site {
	return &Site{
		Package: p,
		File:    file,
		Path:    path,
		Target:  target,
		Outer:   outer,
		Parent:  parent,
	}
}
