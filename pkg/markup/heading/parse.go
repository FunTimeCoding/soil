package heading

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
	"strconv"
)

func Parse(content string) []*Heading {
	source := []byte(body(content))
	seen := map[string]int{}
	var result []*Heading
	errors.PanicOnError(
		ast.Walk(
			parser.New().Parse(source),
			func(
				n ast.Node,
				entering bool,
			) (ast.WalkStatus, error) {
				h, okay := n.(*ast.Heading)

				if !entering || !okay {
					return ast.WalkContinue, nil
				}

				t := text(h, source)
				slug := Slug(t)

				if count := seen[slug]; count > 0 {
					seen[slug]++
					slug = join.Empty(slug, constant.Dash, strconv.Itoa(count))
				} else {
					seen[slug] = 1
				}

				result = append(result, New(t, slug))

				return ast.WalkSkipChildren, nil
			},
		),
	)

	return result
}
