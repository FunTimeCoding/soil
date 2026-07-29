package variable_naming

import "github.com/funtimecoding/soil/pkg/lint/constant"

func isEligible(v typedVariable) bool {
	if v.exempt {
		return false
	}

	if v.kind == constant.VariableKindReceiver {
		return true
	}

	if v.kind == constant.VariableKindParameter {
		if isPrimitiveType(v.typ) {
			return len(v.ident.Name) == 1
		}

		return true
	}

	if len(v.ident.Name) == 1 {
		return true
	}

	return isErrorType(v.typ)
}
