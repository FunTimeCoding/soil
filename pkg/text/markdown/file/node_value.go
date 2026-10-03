package file

import "github.com/yuin/goldmark/v2/ast"

func NodeValue(
	s *[]byte,
	n ast.Node,
) string {
	switch o := n.(type) {
	case *ast.Heading:
		return Value(s, o)
	case *ast.Paragraph:
		return Value(s, o)
	case *ast.CodeBlock:
		return Value(s, o)
	}

	return ""
}
