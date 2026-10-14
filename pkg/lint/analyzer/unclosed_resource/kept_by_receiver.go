package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func keptByReceiver(
	p *packages.Package,
	body *ast.BlockStmt,
	o *types.Var,
	indices map[*types.Var]int,
) bool {
	var receiver *types.Var

	for v, index := range indices {
		if index == constant.ReceiverArgument {
			receiver = v
		}
	}

	if receiver == nil {
		return false
	}

	var result bool
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			a, okay := n.(*ast.AssignStmt)

			if result || !okay || len(a.Lhs) != len(a.Rhs) {
				return !result
			}

			for i, l := range a.Lhs {
				target := l

				if index, isIndex := l.(*ast.IndexExpr); isIndex {
					target = index.X
				}

				if fieldOfReceiver(p, target, receiver) == "" {
					continue
				}

				for _, v := range appended(p, []ast.Expr{a.Rhs[i]}) {
					if rootIdentifier(p, v) == o {
						result = true
					}
				}
			}

			return !result
		},
	)

	return result
}
