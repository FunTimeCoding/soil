package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func additionTargets(
	all []*packages.Package,
	set *token.FileSet,
	additions []*addition.Parameter,
	names map[string][]string,
	explicit map[string]string,
	index map[string]*types.Package,
	r *output.Results,
) []*addition.Target {
	var result []*addition.Target
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

	for _, a := range additions {
		o, _, e := lookup.FindDeclaration(
			all,
			a.PackagePath,
			a.Name,
			a.Receiver,
		)

		if e != nil {
			refuse("%s", e)

			continue
		}

		f, okay := o.(*types.Func)

		if !okay {
			refuse("%s in %s is not a function", a.Name, a.PackagePath)

			continue
		}

		if seen[f.FullName()] {
			refuse("%s is listed twice", f.FullName())

			continue
		}

		seen[f.FullName()] = true
		owner, file := lookup.FindOwningFile(all, f.Pos())
		declaration := functionDeclaration(file, f.Pos())

		if declaration == nil || declaration.Body == nil {
			refuse("%s has no body to rewrite", f.FullName())

			continue
		}

		if !token.IsIdentifier(a.Parameter) || a.Parameter == "_" {
			refuse("%s is not a parameter name", a.Parameter)

			continue
		}

		unnamed := false

		for _, field := range declaration.Type.Params.List {
			if len(field.Names) == 0 {
				unnamed = true
			}
		}

		if unnamed {
			refuse("%s has unnamed parameters; name them first", f.FullName())

			continue
		}

		s := f.Type().(*types.Signature)
		count := s.Params().Len()

		if a.Position < 0 || a.Position > count {
			refuse(
				"position %d is out of range for %s, which has %d parameters",
				a.Position,
				f.FullName(),
				count,
			)

			continue
		}

		if s.Variadic() && a.Position == count {
			refuse(
				"position %d follows the variadic parameter of %s",
				a.Position,
				f.FullName(),
			)

			continue
		}

		if signatureNames(declaration.Type)[a.Parameter] {
			refuse(
				"%s already has a parameter or result named %s",
				f.FullName(),
				a.Parameter,
			)

			continue
		}

		if bodyUsesName(declaration.Body, a.Parameter) {
			refuse(
				"%s already uses %s in its body; the parameter would collide or shadow",
				f.FullName(),
				a.Parameter,
			)

			continue
		}

		t, qualifiers, g := evaluate(
			set,
			owner,
			declaration.Pos(),
			a.Type,
			true,
			names,
			explicit,
			index,
		)

		if g != nil {
			refuse("%s: %s", f.FullName(), g)

			continue
		}

		result = append(
			result,
			addition.NewTarget(a, f, owner, file, declaration, t, qualifiers),
		)
	}

	return result
}
