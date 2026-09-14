package file

import "github.com/yuin/goldmark/v2/ast"

func NodeValue(
	s *[]byte,
	n ast.Node,
) string {
	switch o := n.(type) {
	case *ast.Document:
		// Has no value
		if false {
			return Value(s, o)
		}
	case *ast.Heading:
		return Value(s, o)
	case *ast.Paragraph:
		return Value(s, o)
	case *ast.Text:
		// Both heading and paragraph can be text
		if false {
			return o.Value.Value(*s)
		}
	case *ast.CodeBlock:
		return Value(s, o)
	}

	return ""
}
