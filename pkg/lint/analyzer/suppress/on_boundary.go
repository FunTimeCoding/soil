package suppress

import (
	"go/ast"
	"go/token"
)

func onBoundary(
	fileSet *token.FileSet,
	node ast.Node,
	comment *ast.Comment,
) bool {
	line := fileSet.Position(comment.Pos()).Line
	first := fileSet.Position(node.Pos()).Line
	last := fileSet.Position(node.End()).Line

	if line == first || line == last {
		return true
	}

	return line == first+1 && commentOnly(fileSet, node, comment)
}
