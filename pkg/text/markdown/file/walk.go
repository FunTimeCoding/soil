package file

import (
	"github.com/funtimecoding/soil/pkg/text/markdown/file/flat"
	"github.com/yuin/goldmark/ast"
)

func Walk(
	s *[]byte,
	n ast.Node,
	l *flat.Flat,
) {
	if n == nil {
		return
	}

	l.Add(NodeValue(s, n))
	PrintNode(s, n)

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		Walk(s, c, l)
	}
}
