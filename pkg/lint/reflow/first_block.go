package reflow

import "github.com/yuin/goldmark/v2/ast"

func firstBlock(n ast.Node) ast.BlockNode {
	for c := n.FirstChild(); c != nil; c = c.FirstChild() {
		if b, okay := c.(ast.BlockNode); okay && len(b.Source()) > 0 {
			return b
		}
	}

	return nil
}
