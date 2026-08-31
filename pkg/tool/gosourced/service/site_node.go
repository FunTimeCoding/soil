package service

import "go/ast"

func siteNode(path []ast.Node) ast.Node {
	for _, node := range path {
		if statement, okay := node.(ast.Stmt); okay {
			return statement
		}
	}

	for _, node := range path {
		if declaration, okay := node.(ast.Decl); okay {
			return declaration
		}
	}

	return path[0]
}
