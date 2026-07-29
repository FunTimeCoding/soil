package variable_naming

import "github.com/funtimecoding/soil/pkg/lint/constant"

func applyParameterExemptions(variables []typedVariable) {
	typeCount := map[string]int{}

	for _, v := range variables {
		if v.kind != constant.VariableKindParameter {
			continue
		}

		typeCount[v.typ.String()]++
	}

	for i := range variables {
		v := &variables[i]

		if v.kind != constant.VariableKindParameter {
			continue
		}

		if isPrimitiveType(v.typ) {
			continue
		}

		if len(v.ident.Name) == 1 {
			continue
		}

		if typeCount[v.typ.String()] >= 2 {
			v.exempt = true

			continue
		}

		if isPrimitiveSlice(v.typ) {
			v.exempt = true
		}
	}
}
