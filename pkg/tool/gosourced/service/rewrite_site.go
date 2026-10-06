package service

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/dstutil"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/literal"
	"go/token"
	"go/types"
	"slices"
)

func rewriteSite(
	decorations *decoration.Set,
	set *token.FileSet,
	structure *types.Struct,
	site *literal.Site,
	parameters []*types.Var,
	packagePath string,
	constructor string,
	assignRest bool,
) (string, error) {
	if !pointerForm(site) {
		return "value literal", nil
	}

	var names, missing []string

	for _, v := range parameters {
		names = append(names, v.Name())

		if !slices.Contains(site.Fields, v.Name()) {
			missing = append(missing, v.Name())
		}
	}

	if len(missing) > 0 {
		return join.Empty("misses ", join.CommaSpace(missing)), nil
	}

	values, order := siteValues(structure, site.Target)
	var extra []string

	for _, name := range order {
		if !slices.Contains(names, name) {
			extra = append(extra, name)
		}
	}

	beyond := fmt.Sprintf(
		"sets %s beyond the constructor",
		join.CommaSpace(extra),
	)
	statement, target := assignmentTarget(site)

	if len(extra) > 0 && (!assignRest || statement == nil) {
		return beyond, nil
	}

	if reordersCalls(values, order, append(slices.Clone(names), extra...)) {
		return "evaluation order", nil
	}

	file, e := decorations.DecorateFile(set, site.Package, site.File)

	if e != nil {
		return "", e
	}

	old := decorations.DecoratedNode(site.Package, site.Outer)

	if old == nil || hasDecorations(old) {
		return "comments inside", nil
	}

	clone := func(name string) dst.Expr {
		return dst.Clone(
			decorations.DecoratedNode(site.Package, values[name]),
		).(dst.Expr)
	}
	call := &dst.CallExpr{Fun: &dst.Ident{Name: constructor, Path: packagePath}}

	for _, name := range names {
		call.Args = append(call.Args, clone(name))
	}

	var assignments []dst.Stmt

	for _, name := range extra {
		assignments = append(
			assignments,
			&dst.AssignStmt{
				Lhs: []dst.Expr{
					&dst.SelectorExpr{
						X:   dst.Clone(decorations.DecoratedNode(site.Package, target)).(dst.Expr),
						Sel: dst.NewIdent(name),
					},
				},
				Tok: token.ASSIGN,
				Rhs: []dst.Expr{clone(name)},
			},
		)
	}

	var anchor dst.Node

	if statement != nil {
		anchor = decorations.DecoratedNode(site.Package, statement)
	}

	if anchor != nil && len(assignments) > 0 {
		spacing := anchor.Decorations()

		for _, a := range assignments {
			a.Decorations().Before = dst.NewLine
			a.Decorations().After = dst.NewLine
		}

		assignments[len(assignments)-1].Decorations().After = spacing.After
		spacing.After = dst.NewLine
	}

	dstutil.Apply(
		file,
		func(c *dstutil.Cursor) bool {
			switch c.Node() {
			case old:
				c.Replace(call)

				return false
			case anchor:
				if anchor != nil {
					for i := len(assignments) - 1; i >= 0; i-- {
						c.InsertAfter(assignments[i])
					}
				}
			}

			return true
		},
		nil,
	)

	return "", nil
}
