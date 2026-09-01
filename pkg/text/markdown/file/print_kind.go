package file

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/yuin/goldmark/ast"
)

func PrintKind(
	s *[]byte,
	n ast.Node,
) {
	if n.Kind() == ast.KindText {
		t := n.(*ast.Text)
		e := t.Segment
		console.Format("Text kind: %s\n", string(e.Value(*s)))
	}
}
