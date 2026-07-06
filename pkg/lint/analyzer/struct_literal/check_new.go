package struct_literal

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkNew(
	p *packages.Package,
	results *output.Results,
	call *ast.CallExpr,
	module string,
) {
	fun, okay := call.Fun.(*ast.Ident)

	if !okay || fun.Name != "new" {
		return
	}

	if len(call.Args) != 1 {
		return
	}

	sel, okay := call.Args[0].(*ast.SelectorExpr)

	if !okay {
		return
	}

	t := ownedType(p, sel, module)

	if t == nil {
		return
	}

	if suppress.IsSuppressed(p.Fset, p.Syntax, call.Pos(), "struct_literal") {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"struct_literal",
			fmt.Sprintf(
				"use a constructor function instead of new(%s.%s)",
				t.Pkg().Name(),
				t.Name(),
			),
			p.Fset.Position(call.Pos()).Filename,
			false,
		),
	)
}
