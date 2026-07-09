package restricted_call

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
	rule Rule,
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

	if a == nil || a.Path() != rule.Package || f.Name() != rule.Function {
		return
	}

	if suppress.IsSuppressed(
		p.Fset,
		p.Syntax,
		call.Pos(),
		"restricted_call",
	) {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"restricted_call",
			rule.Message,
			p.Fset.Position(call.Pos()).Filename,
			false,
		),
	)
}
