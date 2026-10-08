package variable_naming

import (
	"go/ast"
	"go/types"
)

func collectVariables(
	y *types.Info,
	body *ast.BlockStmt,
) []TypedVariable {
	var result []TypedVariable
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			switch node := n.(type) {
			case *ast.AssignStmt:
				collectFromAssign(y, node, &result)
			case *ast.RangeStmt:
				collectFromRange(y, node, &result)
			case *ast.DeclStmt:
				collectFromVariableDeclaration(y, node, &result)
			}

			return true
		},
	)

	return result
}
