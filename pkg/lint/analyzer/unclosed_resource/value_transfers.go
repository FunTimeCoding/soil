package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_transfer"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func valueTransfers(
	p *packages.Package,
	body *ast.BlockStmt,
	e ast.Expr,
) []resource_transfer.Transfer {
	switch v := ast.Unparen(e).(type) {
	case *ast.UnaryExpr:
		if v.Op == token.AND {
			return valueTransfers(p, body, v.X)
		}
	case *ast.CompositeLit:
		return literalTransfers(p, v)
	case *ast.Ident:
		o, okay := p.TypesInfo.Uses[v].(*types.Var)

		if !okay {
			return nil
		}

		var result []resource_transfer.Transfer

		for _, a := range assignedValues(p, body, o) {
			if l := literalOf(a); l != nil {
				result = append(result, literalTransfers(p, l)...)
			}
		}

		return append(result, localFieldTransfers(p, body, o)...)
	}

	return nil
}
