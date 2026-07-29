package variable_naming

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
)

func collectFromParams(
	y *types.Info,
	f *ast.FuncDecl,
	result *[]typedVariable,
) {
	if f.Type == nil || f.Type.Params == nil {
		return
	}

	for _, field := range f.Type.Params.List {
		for _, name := range field.Names {
			if name.Name == "_" {
				continue
			}

			o := y.ObjectOf(name)

			if o == nil {
				continue
			}

			*result = append(
				*result,
				typedVariable{
					ident:       name,
					typ:         o.Type(),
					precedence:  typePrecedence(o.Type()),
					scopedNames: collectScopedNames(o),
					kind:        constant.VariableKindParameter,
				},
			)
		}
	}
}
