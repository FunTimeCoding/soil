package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

func removalCalls(
	all []*packages.Package,
	set *token.FileSet,
	targets []*removal.Target,
	r *output.Results,
) []*removal.Call {
	var result []*removal.Call

	for _, t := range targets {
		for _, reference := range resolve.FindAllReferences(all, t.Function) {
			if reference.Ident.Pos() == t.Function.Pos() {
				continue
			}

			owner, file := findOwningFile(all, reference.Ident.Pos())

			if file == nil {
				continue
			}

			path, _ := astutil.PathEnclosingInterval(
				file,
				reference.Ident.Pos(),
				reference.Ident.End(),
			)
			call := enclosingCall(path)

			if call == nil {
				addPositionConcern(
					r,
					set,
					reference.Ident.Pos(),
					constant.ConcernValueUse,
					false,
					"%s is used as a value; its signature cannot change",
					t.Function.Name(),
				)

				continue
			}

			arguments, okay := droppedArguments(owner, call, t)

			if !okay {
				addPositionConcern(
					r,
					set,
					call.Pos(),
					constant.ConcernValueUse,
					false,
					"%s is called with a multi-value result",
					t.Function.Name(),
				)

				continue
			}

			result = append(
				result,
				removal.NewCall(t, owner, file, call, arguments),
			)
		}

		checkInterfaceSatisfaction(all, set, t, r)
	}

	return result
}
