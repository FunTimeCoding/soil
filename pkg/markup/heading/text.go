package heading

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/yuin/goldmark/v2/ast"
	"strings"
)

func text(
	h *ast.Heading,
	source []byte,
) string {
	var result strings.Builder
	errors.PanicOnError(
		ast.Walk(
			h,
			func(
				n ast.Node,
				entering bool,
			) (ast.WalkStatus, error) {
				if !entering {
					return ast.WalkContinue, nil
				}

				switch v := n.(type) {
				case *ast.Text:
					result.WriteString(v.Value.Value(source))

					if v.SoftLineBreak() {
						result.WriteRune(' ')
					}
				case *ast.CodeSpan:
					result.WriteString(v.Value.Value(source))

					return ast.WalkSkipChildren, nil
				}

				return ast.WalkContinue, nil
			},
		),
	)

	return result.String()
}
