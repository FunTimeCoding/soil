package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func holdsClosure(
	p *packages.Package,
	e ast.Expr,
	o *types.Var,
	closures map[*types.Var]bool,
) bool {
	switch v := ast.Unparen(e).(type) {
	case *ast.FuncLit:
		return references(p, v.Body, o)
	case *ast.Ident:
		c, okay := p.TypesInfo.Uses[v].(*types.Var)

		return okay && closures[c]
	}

	return false
}
