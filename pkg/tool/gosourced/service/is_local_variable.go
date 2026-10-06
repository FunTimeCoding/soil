package service

import "go/types"

func isLocalVariable(o types.Object) (*types.Var, bool) {
	v, okay := o.(*types.Var)

	if !okay || v.IsField() || v.Pkg() == nil || v.Parent() == nil {
		return nil, false
	}

	return v, v.Parent() != v.Pkg().Scope()
}
