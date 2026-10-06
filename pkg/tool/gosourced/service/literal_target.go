package service

import (
	"go/ast"
	"go/types"
)

func literalTarget(
	information *types.Info,
	named *types.Named,
	n ast.Node,
) ast.Expr {
	switch x := n.(type) {
	case *ast.CompositeLit:
		t := information.TypeOf(x)

		if p, okay := t.(*types.Pointer); okay && x.Type == nil {
			t = p.Elem()
		}

		if sameNamed(t, named) {
			return x
		}
	case *ast.CallExpr:
		i, okay := x.Fun.(*ast.Ident)

		if !okay || len(x.Args) != 1 {
			return nil
		}

		if _, builtin := information.Uses[i].(*types.Builtin); !builtin {
			return nil
		}

		if sameNamed(information.TypeOf(x.Args[0]), named) {
			return x
		}
	}

	return nil
}
