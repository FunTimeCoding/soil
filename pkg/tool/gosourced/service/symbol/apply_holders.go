package symbol

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/concern_line"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/relocation"
	"go/ast"
	"maps"
	"slices"
)

func applyHolders(
	r *output.Results,
	decorations *decoration.Set,
	plan *relocation.Plan,
) error {
	selectors := slices.SortedFunc(
		maps.Keys(plan.Holders),
		func(a, b *ast.SelectorExpr) int { return int(a.Pos() - b.Pos()) },
	)

	for _, s := range selectors {
		field := plan.Holders[s]
		owner, file := lookup.FindOwningFile(plan.All, s.Pos())

		if file == nil {
			continue
		}

		if _, e := decorations.DecorateFile(plan.Set, owner, file); e != nil {
			return e
		}

		d, okay := decorations.DecoratedNode(owner, s).(*dst.SelectorExpr)

		if !okay {
			continue
		}

		d.X = &dst.SelectorExpr{X: d.X, Sel: dst.NewIdent(field)}
		concern_line.Add(
			r,
			plan.Set,
			s.Pos(),
			constant.ConcernRouted,
			true,
			"%s → %s.%s",
			s.Sel.Name,
			field,
			s.Sel.Name,
		)
	}

	return nil
}
