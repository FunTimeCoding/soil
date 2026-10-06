package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"go/token"
	"go/types"
)

func literalShape(
	structure *types.Struct,
	target ast.Expr,
	outer ast.Expr,
	nested bool,
) string {
	var markers []string
	shape := "new()"

	if lit, okay := target.(*ast.CompositeLit); okay {
		fields, unkeyed := literalFields(structure, lit)
		shape = join.Empty("{", join.CommaSpace(fields), "}")

		if u, isUnary := outer.(*ast.UnaryExpr); isUnary && u.Op == token.AND {
			shape = join.Empty("&", shape)
		}

		if unkeyed {
			markers = append(markers, "unkeyed")
		}

		if lit.Type == nil {
			markers = append(markers, "elided")
		} else if outer == target {
			markers = append(markers, "value")
		}
	}

	if nested {
		markers = append(markers, "nested")
	}

	return join.Space(append([]string{shape}, markers...)...)
}
