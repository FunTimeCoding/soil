package variable_naming

import "github.com/funtimecoding/soil/pkg/lint/constant"

func scopePriority(k constant.VariableKind) int {
	switch k {
	case constant.VariableKindReceiver:
		return 0
	case constant.VariableKindParameter:
		return 1
	default:
		return 2
	}
}
