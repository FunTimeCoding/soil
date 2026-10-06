package service

import "go/ast"

func isParameterPath(path []ast.Node) bool {
	if len(path) < 4 {
		return false
	}

	_, field := path[1].(*ast.Field)
	_, list := path[2].(*ast.FieldList)
	_, function := path[3].(*ast.FuncType)

	return field && list && function
}
