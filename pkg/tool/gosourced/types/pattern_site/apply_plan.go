package pattern_site

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type ApplyPlan struct {
	Package   *packages.Package
	Statement ast.Node
	Parent    ast.Node
	Bindings  map[string]ast.Expr
	Anchor    ast.Node
}
