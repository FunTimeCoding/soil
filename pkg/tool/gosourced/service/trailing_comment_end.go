package service

import (
	"go/ast"
	"go/token"
)

func trailingCommentEnd(
	set *token.FileSet,
	file *ast.File,
	end token.Pos,
) token.Pos {
	line := set.Position(end).Line

	for _, group := range file.Comments {
		if group.Pos() > end && set.Position(group.Pos()).Line == line {
			return group.End()
		}
	}

	return end
}
