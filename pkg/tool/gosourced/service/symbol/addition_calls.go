package symbol

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

func additionCalls(
	all []*packages.Package,
	set *token.FileSet,
	targets []*addition.Target,
	names map[string][]string,
	explicit map[string]string,
	index map[string]*types.Package,
	r *output.Results,
) []*addition.Call {
	var result []*addition.Call

	for _, t := range targets {
		count := t.Function.Type().(*types.Signature).Params().Len()

		for _, reference := range resolve.FindAllReferences(all, t.Function) {
			if reference.Ident.Pos() == t.Function.Pos() {
				continue
			}

			owner, file := lookup.FindOwningFile(all, reference.Ident.Pos())

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

			if len(call.Args) == 1 && count > 1 {
				if _, tuple := owner.TypesInfo.TypeOf(call.Args[0]).(*types.Tuple); tuple {
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
			}

			argument, qualifiers, e := evaluate(
				set,
				owner,
				call.Pos(),
				t.Addition.Argument,
				false,
				names,
				explicit,
				index,
			)

			if e != nil {
				addPositionConcern(
					r,
					set,
					call.Pos(),
					constant.ConcernValidation,
					false,
					"%s: argument %s does not resolve here: %s",
					t.Function.Name(),
					t.Addition.Argument,
					e,
				)

				continue
			}

			if !types.AssignableTo(argument, t.Type) {
				addPositionConcern(
					r,
					set,
					call.Pos(),
					constant.ConcernValidation,
					false,
					"%s: argument %s (%s) is not assignable to %s",
					t.Function.Name(),
					t.Addition.Argument,
					argument,
					t.Type,
				)

				continue
			}

			result = append(
				result,
				addition.NewCall(t, owner, file, call, qualifiers),
			)
		}

		checkInterfaceSatisfaction(all, set, t.Function, r)
	}

	return result
}
