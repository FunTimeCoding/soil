package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func lentFrom(
	p *packages.Package,
	body *ast.BlockStmt,
	e ast.Expr,
	indices map[*types.Var]int,
	visited map[*types.Var]bool,
	via *[]string,
) bool {
	switch v := ast.Unparen(e).(type) {
	case *ast.Ident:
		o, okay := p.TypesInfo.Uses[v].(*types.Var)

		if !okay {
			return false
		}

		if _, parameter := indices[o]; parameter {
			return true
		}

		if visited[o] {
			return false
		}

		visited[o] = true

		if keptByReceiver(p, body, o, indices) {
			return true
		}

		values := assignedValues(p, body, o)

		if len(values) == 0 {
			return false
		}

		for _, a := range values {
			if a == nil || !lentFrom(p, body, a, indices, visited, via) {
				return false
			}
		}

		return true
	case *ast.TypeAssertExpr:
		return lentFrom(p, body, v.X, indices, visited, via)
	case *ast.SelectorExpr, *ast.StarExpr, *ast.IndexExpr:
		r := rootIdentifier(p, v)
		_, parameter := indices[r]

		return parameter
	case *ast.CallExpr:
		s, okay := v.Fun.(*ast.SelectorExpr)

		if !okay {
			return false
		}

		selection := p.TypesInfo.Selections[s]
		f := calleeOf(p, v)

		if selection == nil || selection.Kind() != types.MethodVal || f == nil {
			return false
		}

		if !lentFrom(p, body, s.X, indices, visited, via) {
			return false
		}

		*via = append(*via, functionKey(f))

		return true
	}

	return false
}
