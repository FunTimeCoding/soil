package restricted_literal

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/types/restriction"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkLiteral(
	p *packages.Package,
	results *output.Results,
	l *ast.CompositeLit,
	rule restriction.Restriction,
) {
	named, isNamed := p.TypesInfo.TypeOf(l).(*types.Named)

	if !isNamed {
		return
	}

	o := named.Obj()

	if o.Pkg() == nil ||
		o.Pkg().Path() != rule.Package ||
		o.Name() != rule.Symbol {
		return
	}

	if suppress.IsSuppressed(p.Fset, p.Syntax, l.Pos(), "restricted_literal") {
		return
	}

	results.AddConcern(
		concern.NewPosition(
			"restricted_literal",
			rule.Message,
			p.Fset.Position(l.Pos()),
		),
	)
}
