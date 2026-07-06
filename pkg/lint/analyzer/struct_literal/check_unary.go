package struct_literal

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkUnary(
	p *packages.Package,
	results *output.Results,
	expr *ast.UnaryExpr,
	module string,
) {
	lit, okay := expr.X.(*ast.CompositeLit)

	if !okay {
		return
	}

	sel, okay := lit.Type.(*ast.SelectorExpr)

	if !okay {
		return
	}

	t := ownedType(p, sel, module)

	if t == nil {
		return
	}

	if suppress.IsSuppressed(p.Fset, p.Syntax, expr.Pos(), "struct_literal") {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"struct_literal",
			fmt.Sprintf(
				"use a constructor function instead of &%s.%s{}",
				t.Pkg().Name(),
				t.Name(),
			),
			p.Fset.Position(expr.Pos()).Filename,
			false,
		),
	)
}
