package gofix

import "go/ast"

func anyExported(violations []Violation) bool {
	for _, v := range violations {
		if v.fix != "" && ast.IsExported(v.ident.Name) {
			return true
		}
	}

	return false
}
