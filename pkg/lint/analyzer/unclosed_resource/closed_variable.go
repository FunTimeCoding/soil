package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func closedVariable(
	p *packages.Package,
	body *ast.BlockStmt,
	o *types.Var,
) bool {
	result := false
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			if s, okay := n.(*ast.SelectorExpr); okay {
				f, _ := p.TypesInfo.Uses[s.Sel].(*types.Func)
				i, isIdent := ast.Unparen(s.X).(*ast.Ident)
				result = result || (isCloseMethod(f) && isIdent &&
					p.TypesInfo.Uses[i] == o)
			}

			return !result
		},
	)

	return result
}
