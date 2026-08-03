package stray_variable

import (
	"go/ast"
	"strings"
)

func isEmbedded(
	declaration *ast.GenDecl,
	value *ast.ValueSpec,
) bool {
	for _, group := range []*ast.CommentGroup{
		declaration.Doc,
		value.Doc,
	} {
		if group == nil {
			continue
		}

		for _, c := range group.List {
			if strings.HasPrefix(c.Text, "//go:embed") {
				return true
			}
		}
	}

	return false
}
