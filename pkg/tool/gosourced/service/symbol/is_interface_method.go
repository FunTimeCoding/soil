package symbol

import "go/types"

func isInterfaceMethod(o types.Object) bool {
	f, isFunction := o.(*types.Func)

	if !isFunction {
		return false
	}

	receiver := f.Type().(*types.Signature).Recv()

	if receiver == nil {
		return false
	}

	_, isInterface := receiver.Type().Underlying().(*types.Interface)

	return isInterface
}
