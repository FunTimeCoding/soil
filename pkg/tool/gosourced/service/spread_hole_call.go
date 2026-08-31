package service

import "github.com/dave/dst"

func spreadHoleCall(call *dst.CallExpr, holes map[string]string) bool {
	if len(call.Args) != 1 || !call.Ellipsis {
		return false
	}

	hole, plain := call.Args[0].(*dst.Ident)

	if !plain {
		return false
	}

	_, isHole := holes[hole.Name]

	return isHole
}
