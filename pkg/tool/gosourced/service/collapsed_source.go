package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"go/token"
	"strings"
)

func collapsedSource(
	content []byte,
	set *token.FileSet,
	node ast.Node,
) string {
	start := set.Position(node.Pos()).Offset
	end := set.Position(node.End()).Offset

	return strings.NewReplacer(
		"{ ",
		"{",
		"( ",
		"(",
		", }",
		"}",
		", )",
		")",
		" }",
		"}",
		" )",
		")",
	).Replace(
		strings.Join(
			strings.Fields(string(content[start:end])),
			constant.FieldSeparator,
		),
	)
}
