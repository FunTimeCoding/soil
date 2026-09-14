package file

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/yuin/goldmark/v2/ast"
)

func PrintKind(
	s *[]byte,
	n ast.Node,
) {
	if n.Kind() == ast.KindText {
		t := n.(*ast.Text)
		console.Format("Text kind: %s\n", t.Value.Value(*s))
	}
}
