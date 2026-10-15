package suppress

import (
	"go/ast"
	"go/token"
)

func IsSuppressed(
	fileSet *token.FileSet,
	files []*ast.File,
	position token.Pos,
	key string,
) bool {
	line := fileSet.Position(position).Line
	filename := fileSet.Position(position).Filename

	for _, file := range files {
		if fileSet.File(file.Pos()).Name() != filename {
			continue
		}

		for _, group := range file.Comments {
			for _, comment := range group.List {
				if fileSet.Position(comment.Pos()).Line != line {
					continue
				}

				if CommentSuppresses(comment.Text, key) {
					return true
				}
			}
		}
	}

	return false
}
