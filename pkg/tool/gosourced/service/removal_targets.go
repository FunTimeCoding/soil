package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/removal"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func removalTargets(
	all []*packages.Package,
	removals []*removal.Parameter,
	r *output.Results,
) []*removal.Target {
	var result []*removal.Target
	seen := map[string]bool{}
	refuse := func(format string, arguments ...any) {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernValidation,
				fmt.Sprintf(format, arguments...),
				"",
				false,
			),
		)
	}

	for _, m := range removals {
		o, _, e := findDeclaration(all, m.PackagePath, m.Name, m.Receiver)

		if e != nil {
			refuse("%s", e)

			continue
		}

		f, okay := o.(*types.Func)

		if !okay {
			refuse("%s in %s is not a function", m.Name, m.PackagePath)

			continue
		}

		if seen[f.FullName()] {
			refuse("%s is listed twice", f.FullName())

			continue
		}

		seen[f.FullName()] = true
		owner, file := findOwningFile(all, f.Pos())
		declaration := functionDeclarationAt(file, f.Pos())

		if declaration == nil || declaration.Body == nil {
			refuse("%s has no body to rewrite", f.FullName())

			continue
		}

		t := removal.NewTarget(f, owner, file, declaration)
		names := map[string]*ast.Ident{}
		indices := map[string]int{}
		index := 0

		for _, field := range declaration.Type.Params.List {
			if len(field.Names) == 0 {
				index++

				continue
			}

			for _, n := range field.Names {
				names[n.Name] = n
				indices[n.Name] = index
				index++
			}
		}

		for _, p := range m.Parameters {
			n, found := names[p]

			if !found || p == "_" {
				refuse("%s has no parameter %s", f.FullName(), p)

				continue
			}

			v, isVariable := owner.TypesInfo.Defs[n].(*types.Var)

			if !isVariable {
				refuse("%s parameter %s has no object", f.FullName(), p)

				continue
			}

			t.Indices[indices[p]] = true
			t.Variables = append(t.Variables, v)
		}

		result = append(result, t)
	}

	return result
}
