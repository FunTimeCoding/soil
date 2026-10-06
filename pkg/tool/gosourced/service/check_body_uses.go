package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"go/ast"
	"go/token"
)

func checkBodyUses(
	set *token.FileSet,
	targets []*removal.Target,
	dropped []ast.Node,
	r *output.Results,
) {
	for _, t := range targets {
		for _, v := range t.Variables {
			ast.Inspect(
				t.Declaration.Body,
				func(n ast.Node) bool {
					i, okay := n.(*ast.Ident)

					if !okay || t.Owner.TypesInfo.Uses[i] != v ||
						spansContain(dropped, i) {
						return true
					}

					addPositionConcern(
						r,
						set,
						i.Pos(),
						constant.ConcernStillUsed,
						false,
						"parameter %s of %s is still used",
						v.Name(),
						t.Function.Name(),
					)

					return true
				},
			)
		}
	}
}
