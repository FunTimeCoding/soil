package revision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"go/token"
)

func anonymous(
	s *token.FileSet,
	n ast.Node,
	names ...*ast.Ident,
) string {
	original := make([]string, len(names))

	for i, name := range names {
		original[i] = name.Name
		name.Name = constant.BlankName
	}

	result := render(s, n)

	for i, name := range names {
		name.Name = original[i]
	}

	return result
}
