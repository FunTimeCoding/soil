package goflow

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goflow/constant"
	"github.com/yuin/goldmark/v2/ast"
)

func literal(
	source []byte,
	document ast.Node,
) []bool {
	result := make([]bool, len(source))
	errors.PanicOnError(
		ast.Walk(
			document,
			func(
				n ast.Node,
				entering bool,
			) (ast.WalkStatus, error) {
				c, okay := n.(*ast.CodeSpan)

				if !entering || !okay || c.Value.IsOwned() {
					return ast.WalkContinue, nil
				}

				indices := c.Value.Indices()

				if len(indices) == 0 {
					return ast.WalkContinue, nil
				}

				start := indices[0].Start
				stop := indices[len(indices)-1].Stop

				for start > 0 && source[start-1] == constant.Backtick {
					start--
				}

				for stop < len(source) && source[stop] == constant.Backtick {
					stop++
				}

				for i := start; i < stop; i++ {
					result[i] = true
				}

				return ast.WalkContinue, nil
			},
		),
	)

	return result
}
