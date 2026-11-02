package unclosed_resource

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func escapes(
	p *packages.Package,
	body *ast.BlockStmt,
	o *types.Var,
	match matcher,
) bool {
	closures := closureVariables(p, body, o)
	leaves := func(
		p *packages.Package,
		e ast.Expr,
		o *types.Var,
	) bool {
		return match(p, e, o) || holdsClosure(p, e, o, closures)
	}
	var result bool
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			if result {
				return false
			}

			switch v := n.(type) {
			case *ast.ReturnStmt:
				result = anyMatches(p, v.Results, o, leaves)
			case *ast.SendStmt:
				result = leaves(p, v.Value, o)
			case *ast.AssignStmt:
				result = anyMatches(p, appended(p, v.Rhs), o, match) ||
					storesClosure(p, v, o, closures)
			case *ast.CompositeLit:
				result = anyMatches(p, elementValues(v.Elts), o, leaves)
			case *ast.UnaryExpr:
				result = v.Op == token.AND && match(p, v.X, o)
			case *ast.CallExpr:
				result = anyMatches(
					p,
					v.Args,
					o,
					func(
						p *packages.Package,
						e ast.Expr,
						o *types.Var,
					) bool {
						return holdsClosure(p, e, o, closures)
					},
				)
			}

			return !result
		},
	)

	return result
}
