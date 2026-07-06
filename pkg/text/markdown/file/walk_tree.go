package file

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/text/markdown/file/flat"
	"github.com/yuin/goldmark/ast"
)

func WalkTree(
	s *[]byte,
	n ast.Node,
	l *flat.Flat,
) *flat.Flat {
	errors.PanicOnError(
		ast.Walk(
			n,
			func(
				n ast.Node,
				entering bool,
			) (ast.WalkStatus, error) {
				if !entering {
					return ast.WalkContinue, nil
				}

				l.Add(NodeValue(s, n))

				if false {
					PrintNode(s, n)
					PrintKind(s, n)
				}

				return ast.WalkContinue, nil
			},
		),
	)

	return l
}
