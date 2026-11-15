package suppress

import (
	"go/ast"
	"go/token"
)

func commentOnly(
	fileSet *token.FileSet,
	node ast.Node,
	comment *ast.Comment,
) bool {
	line := fileSet.Position(comment.Pos()).Line
	result := true
	ast.Inspect(
		node,
		func(n ast.Node) bool {
			if n == nil || !result {
				return false
			}

			if fileSet.Position(n.Pos()).Line == line {
				result = false
			}

			return true
		},
	)

	return result
}
