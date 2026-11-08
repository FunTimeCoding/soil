package preposition_name

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"go/ast"
)

func declaredNames(
	d ast.Decl,
	test bool,
) []*ast.Ident {
	if f, okay := d.(*ast.FuncDecl); okay {
		if test && lint.IsTestFunction(f) {
			return nil
		}

		return []*ast.Ident{f.Name}
	}

	g, okay := d.(*ast.GenDecl)

	if !okay {
		return nil
	}

	var result []*ast.Ident

	for _, spec := range g.Specs {
		s, isType := spec.(*ast.TypeSpec)

		if !isType {
			continue
		}

		result = append(result, s.Name)

		switch t := s.Type.(type) {
		case *ast.StructType:
			result = append(result, fieldNames(t.Fields)...)
		case *ast.InterfaceType:
			result = append(result, fieldNames(t.Methods)...)
		}
	}

	return result
}
