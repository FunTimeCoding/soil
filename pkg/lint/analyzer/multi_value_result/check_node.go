package multi_value_result

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkNode(
	p *packages.Package,
	results *output.Results,
	n ast.Node,
) {
	switch node := n.(type) {
	case *ast.ExprStmt:
		call, okay := node.X.(*ast.CallExpr)

		if okay && multiValue(p, call) {
			report(
				p,
				results,
				call,
				fmt.Sprintf(
					"multi-value result of %s discarded as a bare statement",
					calleeName(call),
				),
			)
		}
	case *ast.GoStmt:
		if multiValue(p, node.Call) {
			report(
				p,
				results,
				node.Call,
				fmt.Sprintf(
					"multi-value result of %s discarded by go",
					calleeName(node.Call),
				),
			)
		}
	case *ast.DeferStmt:
		if multiValue(p, node.Call) {
			report(
				p,
				results,
				node.Call,
				fmt.Sprintf(
					"multi-value result of %s discarded by defer",
					calleeName(node.Call),
				),
			)
		}
	case *ast.CallExpr:
		if len(node.Args) != 1 {
			return
		}

		if f := calleeFunction(p, node); f != nil &&
			constant.Absorbers[f.FullName()] {
			return
		}

		inner, okay := node.Args[0].(*ast.CallExpr)

		if okay && multiValue(p, inner) {
			report(
				p,
				results,
				inner,
				fmt.Sprintf(
					"multi-value result absorbed by %s - every value becomes an argument silently",
					calleeName(node),
				),
			)
		}
	}
}
