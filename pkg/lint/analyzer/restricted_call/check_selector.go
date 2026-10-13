package restricted_call

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/types/restriction"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkSelector(
	p *packages.Package,
	results *output.Results,
	sel *ast.SelectorExpr,
	rule restriction.Restriction,
) {
	o, isUse := p.TypesInfo.Uses[sel.Sel]

	if !isUse || o.Pkg() == nil {
		return
	}

	switch v := o.(type) {
	case *types.Func:
		if v.Type().(*types.Signature).Recv() != nil {
			return
		}
	case *types.Var:
		if v.Parent() != v.Pkg().Scope() {
			return
		}
	default:
		return
	}

	if o.Pkg().Path() != rule.Package || o.Name() != rule.Symbol {
		return
	}

	if suppress.IsSuppressed(p.Fset, p.Syntax, sel.Pos(), "restricted_call") {
		return
	}

	results.AddConcern(
		concern.NewPosition(
			"restricted_call",
			rule.Message,
			p.Fset.Position(sel.Pos()),
		),
	)
}
