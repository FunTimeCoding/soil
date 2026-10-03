package reflow

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/yuin/goldmark/v2/ast"
	"strings"
)

func Interruptions(content string) []int {
	source := []byte(content)
	body := front(source)
	var result []int
	e := ast.Walk(
		parse(source),
		func(
			n ast.Node,
			entering bool,
		) (ast.WalkStatus, error) {
			if !entering || n.Kind() != ast.KindList {
				return ast.WalkContinue, nil
			}

			if start, okay := interruption(source, n); okay && start >= body {
				result = append(
					result,
					strings.Count(content[:start], constant.Unix)+1,
				)
			}

			return ast.WalkContinue, nil
		},
	)

	if e != nil {
		return nil
	}

	return result
}
