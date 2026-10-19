package error_answer

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func conditionClasses(
	p *packages.Package,
	e ast.Expr,
) []string {
	if x, okay := e.(*ast.ParenExpr); okay {
		return conditionClasses(p, x.X)
	}

	if x, okay := e.(*ast.BinaryExpr); okay && x.Op == token.LOR {
		left := conditionClasses(p, x.X)
		right := conditionClasses(p, x.Y)

		if left == nil || right == nil {
			return nil
		}

		return append(left, right...)
	}

	if x, okay := e.(*ast.CallExpr); okay {
		if c := classOf(p, x); c != "" {
			return []string{c}
		}
	}

	return nil
}
