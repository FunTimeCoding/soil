package suppress

import (
	"go/ast"
	"go/token"
)

func IsSuppressed(
	fileSet *token.FileSet,
	files []*ast.File,
	node ast.Node,
	key string,
) bool {
	filename := fileSet.Position(node.Pos()).Filename

	for _, file := range files {
		if fileSet.File(file.Pos()).Name() != filename {
			continue
		}

		for _, group := range file.Comments {
			for _, comment := range group.List {
				if !onBoundary(fileSet, node, comment) {
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
