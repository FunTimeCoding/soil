package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"go/token"
	"go/types"
)

func checkSideEffects(
	set *token.FileSet,
	calls []*removal.Call,
	r *output.Results,
) {
	for _, c := range calls {
		for _, a := range c.Arguments {
			if !hasSideEffect(c.Owner.TypesInfo, a) {
				continue
			}

			addPositionConcern(
				r,
				set,
				a.Pos(),
				constant.ConcernSideEffect,
				false,
				"dropping the argument %s would remove a call",
				types.ExprString(a),
			)
		}
	}
}
