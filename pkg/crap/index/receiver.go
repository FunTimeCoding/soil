package index

import (
	"go/ast"
	"go/types"
)

func receiver(declaration *ast.FuncDecl) string {
	if declaration.Recv == nil || len(declaration.Recv.List) == 0 {
		return ""
	}

	return types.ExprString(declaration.Recv.List[0].Type)
}
