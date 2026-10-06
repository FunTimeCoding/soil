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

func reportLeftovers(
	all []*packages.Package,
	set *token.FileSet,
	targets []*removal.Target,
	calls []*removal.Call,
	locals []*removal.Local,
	removed []ast.Node,
	r *output.Results,
) {
	seen := map[token.Pos]bool{}

	for _, t := range targets {
		for _, v := range t.Variables {
			seen[v.Pos()] = true
		}
	}

	inspect := func(
		owner *packages.Package,
		file *ast.File,
		n ast.Node,
	) {
		ast.Inspect(
			n,
			func(x ast.Node) bool {
				i, okay := x.(*ast.Ident)

				if !okay {
					return true
				}

				v, isVariable := owner.TypesInfo.Uses[i].(*types.Var)

				if !isVariable || seen[v.Pos()] {
					return true
				}

				seen[v.Pos()] = true

				if v.IsField() {
					reads, writes := variableUses(all, v, removed)

					if reads == 0 && writes > 0 {
						addPositionConcern(
							r,
							set,
							v.Pos(),
							constant.ConcernWrittenOnly,
							true,
							"field %s is now only written",
							v.Name(),
						)
					}

					return true
				}

				if !isParameterPath(definitionPath(file, v)) {
					return true
				}

				if reads, _ := variableUses(all, v, removed); reads == 0 {
					addPositionConcern(
						r,
						set,
						v.Pos(),
						constant.ConcernUnusedParameter,
						true,
						"parameter %s is no longer used",
						v.Name(),
					)
				}

				return true
			},
		)
	}

	for _, c := range calls {
		for _, a := range c.Arguments {
			inspect(c.Owner, c.File, a)
		}
	}

	for _, l := range locals {
		_, file := findOwningFile(
			[]*packages.Package{l.Owner},
			l.Statement.Pos(),
		)
		inspect(l.Owner, file, l.Statement)
	}
}
