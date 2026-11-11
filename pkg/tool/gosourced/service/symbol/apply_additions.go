package symbol

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/concern_line"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"
	"go/token"
)

func applyAdditions(
	set *token.FileSet,
	decorations *decoration.Set,
	targets []*addition.Target,
	calls []*addition.Call,
	r *output.Results,
) error {
	for _, t := range targets {
		d, e := decorated(set, decorations, t.Owner, t.File, t.Declaration)

		if e != nil {
			return e
		}

		parameterType, f := decoratedExpression(
			t.Addition.Type,
			true,
			t.Qualifiers,
		)

		if f != nil {
			return f
		}

		insertDeclaredParameter(
			d.(*dst.FuncDecl),
			t.Addition.Position,
			t.Addition.Parameter,
			parameterType,
		)
		concern_line.Add(
			r,
			set,
			t.Declaration.Pos(),
			constant.ConcernAdded,
			true,
			"%s: added %s %s",
			t.Function.Name(),
			t.Addition.Parameter,
			t.Addition.Type,
		)
	}

	for _, c := range calls {
		d, e := decorated(set, decorations, c.Owner, c.File, c.Call)

		if e != nil {
			return e
		}

		argument, f := decoratedExpression(
			c.Target.Addition.Argument,
			false,
			c.Qualifiers,
		)

		if f != nil {
			return f
		}

		insertCallArgument(
			d.(*dst.CallExpr),
			c.Target.Addition.Position,
			argument,
		)
		concern_line.Add(
			r,
			set,
			c.Call.Pos(),
			constant.ConcernAdded,
			true,
			"%s: passed %s",
			c.Target.Function.Name(),
			c.Target.Addition.Argument,
		)
	}

	return nil
}
