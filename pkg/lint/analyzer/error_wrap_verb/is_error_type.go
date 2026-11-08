package error_wrap_verb

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func isErrorType(t types.Type) bool {
	if t == nil {
		return false
	}

	universe := types.Universe.Lookup(constant.ErrorType).Type()
	face, okay := universe.Underlying().(*types.Interface)

	if !okay {
		return false
	}

	return types.Implements(t, face)
}
