package service

import "go/types"

func sameObject(a types.Object, b types.Object) bool {
	if a == b {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	first, okay := a.(*types.Func)

	if okay {
		second, also := b.(*types.Func)

		return also && first.FullName() == second.FullName()
	}

	if a.Pkg() == nil || b.Pkg() == nil {
		return false
	}

	if a.Pkg().Path() != b.Pkg().Path() || a.Name() != b.Name() {
		return false
	}

	return a.Parent() == a.Pkg().Scope() && b.Parent() == b.Pkg().Scope()
}
