package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func planLocals(
	all []*packages.Package,
	set *token.FileSet,
	calls []*removal.Call,
	removed []ast.Node,
	r *output.Results,
) ([]*removal.Local, []ast.Node) {
	var result []*removal.Local
	done := map[token.Pos]bool{}

	for _, c := range calls {
		var queue []ast.Node

		for _, a := range c.Arguments {
			queue = append(queue, a)
		}

		for len(queue) > 0 {
			n := queue[0]
			queue = queue[1:]
			ast.Inspect(
				n,
				func(x ast.Node) bool {
					i, okay := x.(*ast.Ident)

					if !okay {
						return true
					}

					v, local := isLocalVariable(c.Owner.TypesInfo.Uses[i])

					if !local || done[v.Pos()] {
						return true
					}

					path := definitionPath(c.File, v)

					if path == nil || isParameterPath(path) {
						return true
					}

					reads, writes := variableUses(all, v, removed)

					if reads > 0 {
						return true
					}

					done[v.Pos()] = true

					if writes > 0 {
						addPositionConcern(
							r,
							set,
							v.Pos(),
							constant.ConcernLocal,
							false,
							"local %s would be left only assigned; remove it by hand",
							v.Name(),
						)

						return true
					}

					statement, block, initializer := declaringStatement(path)

					if statement == nil {
						addPositionConcern(
							r,
							set,
							v.Pos(),
							constant.ConcernLocal,
							false,
							"local %s would be left unused and its declaration cannot be removed alone",
							v.Name(),
						)

						return true
					}

					if initializer != nil &&
						hasSideEffect(c.Owner.TypesInfo, initializer) {
						addPositionConcern(
							r,
							set,
							v.Pos(),
							constant.ConcernSideEffect,
							false,
							"removing the unused local %s would remove a call: %s",
							v.Name(),
							types.ExprString(initializer),
						)

						return true
					}

					result = append(
						result,
						removal.NewLocal(v, c.Owner, statement, block),
					)
					removed = append(removed, statement)

					if initializer != nil {
						queue = append(queue, initializer)
					}

					return true
				},
			)
		}
	}

	return result, removed
}
