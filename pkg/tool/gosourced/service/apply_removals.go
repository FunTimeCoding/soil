package service

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strings"
)

func applyRemovals(
	set *token.FileSet,
	decorations *decoration.Set,
	targets []*removal.Target,
	calls []*removal.Call,
	locals []*removal.Local,
	r *output.Results,
) error {
	decorated := func(
		owner *packages.Package,
		file *ast.File,
		node ast.Node,
	) (dst.Node, error) {
		if _, e := decorations.DecorateFile(set, owner, file); e != nil {
			return nil, e
		}

		d := decorations.DecoratedNode(owner, node)

		if d == nil {
			return nil, validation.New(
				"no decorated node at %s",
				set.Position(node.Pos()),
			)
		}

		return d, nil
	}

	for _, t := range targets {
		d, e := decorated(t.Owner, t.File, t.Declaration)

		if e != nil {
			return e
		}

		removeDeclaredParameters(d.(*dst.FuncDecl), t.Indices)
		var names []string

		for _, v := range t.Variables {
			names = append(names, v.Name())
		}

		addPositionConcern(
			r,
			set,
			t.Declaration.Pos(),
			constant.ConcernRemoved,
			true,
			"%s: removed %s",
			t.Function.Name(),
			strings.Join(names, ", "),
		)
	}

	for _, c := range calls {
		d, e := decorated(c.Owner, c.File, c.Call)

		if e != nil {
			return e
		}

		removeCallArguments(d.(*dst.CallExpr), c.Target)
		addPositionConcern(
			r,
			set,
			c.Call.Pos(),
			constant.ConcernRemoved,
			true,
			"%s: dropped %d argument(s)",
			c.Target.Function.Name(),
			len(c.Arguments),
		)
	}

	for _, l := range locals {
		_, file := findOwningFile(
			[]*packages.Package{l.Owner},
			l.Statement.Pos(),
		)
		statement, e := decorated(l.Owner, file, l.Statement)

		if e != nil {
			return e
		}

		block, e := decorated(l.Owner, file, l.Block)

		if e != nil {
			return e
		}

		if !removeBlockStatement(block.(*dst.BlockStmt), statement.(dst.Stmt)) {
			return validation.New(
				"local %s not found in its block",
				l.Variable.Name(),
			)
		}

		addPositionConcern(
			r,
			set,
			l.Statement.Pos(),
			constant.ConcernRemoved,
			true,
			"local %s removed",
			l.Variable.Name(),
		)
	}

	return nil
}
