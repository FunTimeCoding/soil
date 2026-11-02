package service

import "go/types"

func isMethod(o types.Object) bool {
	f, okay := o.(*types.Func)

	return okay && f.Signature().Recv() != nil
}
