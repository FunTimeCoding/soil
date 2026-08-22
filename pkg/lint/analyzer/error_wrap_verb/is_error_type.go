package error_wrap_verb

import "go/types"

func isErrorType(t types.Type) bool {
	if t == nil {
		return false
	}

	universe := types.Universe.Lookup("error").Type()
	face, okay := universe.Underlying().(*types.Interface)

	if !okay {
		return false
	}

	return types.Implements(t, face)
}
