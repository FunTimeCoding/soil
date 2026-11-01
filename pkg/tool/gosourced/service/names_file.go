package service

import "go/types"

func namesFile(declaration types.Object) bool {
	switch declaration.(type) {
	case *types.Func, *types.TypeName:
		return true
	default:
		return false
	}
}
