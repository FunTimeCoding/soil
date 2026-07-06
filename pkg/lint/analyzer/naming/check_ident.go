package naming

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/segment"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkIdent(
	p *packages.Package,
	results *output.Results,
	ident *ast.Ident,
	o types.Object,
) {
	if isInterfaceMethod(o) {
		return
	}

	v, isVariable := o.(*types.Var)
	isField := isVariable && v.IsField()
	r := segment.Check(ident.Name, isVariable, isField)

	if r == nil {
		return
	}

	if r.Banned {
		results.AddConcern(
			concern.NewFile(
				"naming",
				fmt.Sprintf(
					"avoid %q in name, use a more specific term",
					r.Segment,
				),
				p.Fset.Position(ident.Pos()).Filename,
				false,
			),
		)

		return
	}

	results.AddConcern(
		concern.NewFile(
			"naming",
			segment.FormatMessage(
				r.Applicable,
				r.Segment,
				ident.Name,
			),
			p.Fset.Position(ident.Pos()).Filename,
			false,
		),
	)
}
