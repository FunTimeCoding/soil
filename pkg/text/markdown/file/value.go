package file

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/yuin/goldmark/v2/ast"
	"strings"
)

func Value(
	s *[]byte,
	n ast.Node,
) string {
	if c, okay := n.(*ast.CodeBlock); okay {
		return c.Value.Str(*s)
	}

	var b strings.Builder
	errors.PanicOnError(
		ast.Walk(
			n,
			func(
				n ast.Node,
				entering bool,
			) (ast.WalkStatus, error) {
				if t, okay := n.(*ast.Text); okay && entering {
					b.WriteString(t.Value.Value(*s))
				}

				return ast.WalkContinue, nil
			},
		),
	)

	return b.String()
}
