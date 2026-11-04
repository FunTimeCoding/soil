package symbol

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/naming"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func namingSite(
	p *packages.Package,
	file *ast.File,
	set *token.FileSet,
	signature *ast.FuncType,
	body *ast.BlockStmt,
	t types.Type,
	name string,
	label string,
	r *output.Results,
) *naming.Site {
	var result *naming.Site
	candidates := 0

	for i, field := range signature.Params.List {
		if !types.Identical(p.TypesInfo.TypeOf(field.Type), t) {
			continue
		}

		if len(field.Names) == 0 {
			candidates++
			result = naming.NewSite(p, file, signature, i, -1, label)

			continue
		}

		for j, n := range field.Names {
			if n.Name == "_" {
				candidates++
				result = naming.NewSite(p, file, signature, i, j, label)
			}
		}
	}

	if candidates == 0 {
		return nil
	}

	if candidates > 1 {
		addPositionConcern(
			r,
			set,
			signature.Pos(),
			constant.ConcernValidation,
			false,
			"%s has %d unnamed parameters of type %s; naming one is ambiguous",
			label,
			candidates,
			t,
		)

		return nil
	}

	if signatureNames(signature)[name] || bodyUsesName(body, name) {
		addPositionConcern(
			r,
			set,
			signature.Pos(),
			constant.ConcernValidation,
			false,
			"%s already uses %s; naming the parameter would collide or shadow",
			label,
			name,
		)

		return nil
	}

	return result
}
