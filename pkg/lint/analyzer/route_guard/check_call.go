package route_guard

import (
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

	if a == nil || a.Path() != "net/http" {
		return
	}

	if f.Name() != "Handle" && f.Name() != "HandleFunc" {
		return
	}

	results.AddConcern(
		concern.NewPosition(
			"route_guard",
			"register routes through guard verbs (Open, Token, Session), not directly on the mux",
			p.Fset.Position(call.Pos()),
		),
	)
}
