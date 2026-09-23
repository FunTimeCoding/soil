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
	var result bool
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			if result {
				return false
			}

			switch v := n.(type) {
			case *ast.ReturnStmt:
				result = anyMatches(p, v.Results, o, match)
			case *ast.SendStmt:
				result = match(p, v.Value, o)
			case *ast.AssignStmt:
				result = anyMatches(p, v.Rhs, o, match)
			case *ast.CompositeLit:
				result = anyMatches(p, elementValues(v.Elts), o, match)
			case *ast.UnaryExpr:
				result = v.Op == token.AND && match(p, v.X, o)
			case *ast.FuncLit:
				result = references(p, v.Body, o)
			case *ast.CallExpr:
				result = !isCloseHelperCall(p, v, o) &&
					anyMatches(p, v.Args, o, match)
			}

			return !result
		},
	)

	return result
}
