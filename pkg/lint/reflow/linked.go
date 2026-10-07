package reflow

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/yuin/goldmark/v2/ast"
)

func linked(
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
				if !entering {
					return ast.WalkContinue, nil
				}

				start, stop, okay := linkSpan(source, n)

				if !okay {
					return ast.WalkContinue, nil
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
