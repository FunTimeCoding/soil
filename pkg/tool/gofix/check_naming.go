package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"go/ast"
	"go/types"
)

func checkNaming(
	ident *ast.Ident,
	o types.Object,
) *Violation {
	v, isVariable := o.(*types.Var)
	isField := isVariable && v.IsField()
	r := segment.Check(ident.Name, isVariable, isField)

	if r == nil || r.Banned {
		return nil
	}

	fix := segment.ResolveFix(ident.Name, r.Segment, r.Applicable, o)

	return &Violation{
		ident:   ident,
		object:  o,
		segment: r.Segment,
		fix:     fix,
	}
}
