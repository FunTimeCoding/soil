package goflow

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goflow/constant"
	"github.com/yuin/goldmark/v2/ast"
)

func shape(document ast.Node) []string {
	var result []string
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

				t, okay := n.(*ast.Text)

				if !okay {
					result = append(result, n.Kind().String())
				} else if t.HardLineBreak() {
					result = append(result, constant.HardLineBreak)
				}

				return ast.WalkContinue, nil
			},
		),
	)

	return result
}
