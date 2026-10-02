package goflow

import "github.com/yuin/goldmark/v2/ast"

func quoted(n ast.Node) bool {
	for p := n.Parent(); p != nil; p = p.Parent() {
		if p.Kind() == ast.KindBlockquote {
			return true
		}
	}

	return false
}
