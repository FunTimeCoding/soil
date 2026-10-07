package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_candidate"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func candidates(
	p *packages.Package,
	body *ast.BlockStmt,
) []resource_candidate.Candidate {
	var result []resource_candidate.Candidate
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			a, okay := n.(*ast.AssignStmt)

			if !okay || a.Tok != token.DEFINE || len(a.Rhs) != 1 {
				return true
			}

			call, okay := a.Rhs[0].(*ast.CallExpr)

			if !okay {
				return true
			}

			for _, l := range a.Lhs {
				i, okay := l.(*ast.Ident)

				if !okay || i.Name == constant.Underscore {
					continue
				}

				o, okay := p.TypesInfo.Defs[i].(*types.Var)

				if !okay || !carriesObligation(o.Type()) {
					continue
				}

				result = append(
					result,
					resource_candidate.Candidate{
						Object:   o,
						Position: i.Pos(),
						Call:     call,
					},
				)
			}

			return true
		},
	)

	return result
}
