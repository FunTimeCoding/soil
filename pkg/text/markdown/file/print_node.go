package file

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/yuin/goldmark/ast"
)

func PrintNode(
	s *[]byte,
	n ast.Node,
) {
	switch o := n.(type) {
	case *ast.Document:
		v := Value(s, o)

		if v == "" {
			return
		}

		console.Format("Document: %s\n", v)
	case *ast.Heading:
		console.Format("Heading: %s\n", Value(s, o))
	case *ast.Paragraph:
		console.Format("Paragraph: %s\n", Value(s, o))
	case *ast.Text:
		console.Format("Text: %s\n", string(o.Value(*s)))
	case *ast.FencedCodeBlock:
		console.Format("FencedCodeBlock: %s\n", Value(s, o))
	default:
		console.Format("Unknown node type: %T\n", o)
	}
}
