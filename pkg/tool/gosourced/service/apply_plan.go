package service

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type applyPlan struct {
	p         *packages.Package
	statement ast.Node
	parent    ast.Node
	bindings  map[string]ast.Expr
	anchor    ast.Node
}
