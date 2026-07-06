package unchecked_print_write

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkCall(
	p *packages.Package,
	results *output.Results,
	call *ast.CallExpr,
) {
	selector, okay := call.Fun.(*ast.SelectorExpr)

	if !okay {
		return
	}

	o, okay := p.TypesInfo.Uses[selector.Sel]

	if !okay {
		return
	}

	f, okay := o.(*types.Func)

	if !okay {
		return
	}

	a := f.Pkg()

	if a == nil || a.Path() != "fmt" || f.Name() != "Fprintf" {
		return
	}

	if suppress.IsSuppressed(
		p.Fset,
		p.Syntax,
		call.Pos(),
		"unchecked_print_write",
	) {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"unchecked_print_write",
			"use writer.Print, errors.Printf, or check the error from fmt.Fprintf",
			p.Fset.Position(call.Pos()).Filename,
			false,
		),
	)
}
