package goflow

import "github.com/yuin/goldmark/v2/ast"

func wrappable(n ast.Node) (ast.BlockNode, bool) {
	if n.Kind() != ast.KindParagraph {
		return nil, false
	}

	b, okay := n.(ast.BlockNode)

	return b, okay
}
