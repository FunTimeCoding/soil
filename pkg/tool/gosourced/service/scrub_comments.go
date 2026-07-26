package service

import (
	"go/ast"
	"go/token"
)

func scrubComments(
	file *ast.File,
	start token.Pos,
	end token.Pos,
) {
	var kept []*ast.CommentGroup

	for _, group := range file.Comments {
		if group.Pos() >= start && group.End() <= end {
			continue
		}

		kept = append(kept, group)
	}

	file.Comments = kept
}
