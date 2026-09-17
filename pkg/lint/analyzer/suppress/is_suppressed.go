package suppress

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/token"
	"strings"
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

				text := strings.TrimSpace(
					strings.TrimPrefix(comment.Text, constant.CommentPrefix),
				)

				if text == "goanalyze:ignore" || text == fmt.Sprintf(
					"goanalyze:ignore %s",
					key,
				) {
					return true
				}
			}
		}
	}

	return false
}
