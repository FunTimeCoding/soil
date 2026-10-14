package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_candidate"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func ownsValue(
	p *packages.Package,
	body *ast.BlockStmt,
	s *Summaries,
	owned map[*types.Var]bool,
	e ast.Expr,
) bool {
	switch v := ast.Unparen(e).(type) {
	case *ast.Ident:
		o, _ := p.TypesInfo.Uses[v].(*types.Var)

		return owned[o]
	case *ast.CallExpr:
		return carriesObligation(p.TypesInfo.TypeOf(v)) &&
			!s.isUnowned(p, body, resource_candidate.Candidate{Call: v})
	}

	return false
}
