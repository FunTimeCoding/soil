package gofix

import "go/ast"

func anyExported(violations []violation) bool {
	for _, v := range violations {
		if v.fix != "" && ast.IsExported(v.ident.Name) {
			return true
		}
	}

	return false
}
