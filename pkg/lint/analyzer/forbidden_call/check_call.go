package forbidden_call

import (
	"fmt"
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
	sel, isSel := call.Fun.(*ast.SelectorExpr)

	if !isSel {
		return
	}

	o, isUse := p.TypesInfo.Uses[sel.Sel]

	if !isUse {
		return
	}

	f, isFunction := o.(*types.Func)

	if !isFunction {
		return
	}

	a := f.Pkg()

	if a == nil || a.Path() != "os/exec" {
		return
	}

	if !banned[f.Name()] {
		return
	}

	if suppress.IsSuppressed(p.Fset, p.Syntax, call.Pos(), "forbidden_call") {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"forbidden_call",
			fmt.Sprintf(
				"use pkg/system/run instead of exec.%s",
				f.Name(),
			),
			p.Fset.Position(call.Pos()).Filename,
			false,
		),
	)
}
