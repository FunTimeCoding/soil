package preposition_name

import "go/ast"

func fieldNames(l *ast.FieldList) []*ast.Ident {
	var result []*ast.Ident

	for _, f := range l.List {
		result = append(result, f.Names...)
	}

	return result
}
