package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func droppedArguments(
	owner *packages.Package,
	call *ast.CallExpr,
	t *removal.Target,
) ([]ast.Expr, bool) {
	parameters := t.Function.Type().(*types.Signature).Params().Len()

	if len(call.Args) == 1 && parameters > 1 {
		if _, tuple := owner.TypesInfo.TypeOf(call.Args[0]).(*types.Tuple); tuple {
			return nil, false
		}
	}

	var result []ast.Expr

	for i, a := range call.Args {
		if t.Indices[argumentParameter(t, i)] {
			result = append(result, a)
		}
	}

	return result, true
}
