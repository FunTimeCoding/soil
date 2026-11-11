package symbol

import "go/types"

func isPointerReceiver(o types.Object) bool {
	f, okay := o.(*types.Func)

	if !okay || f.Signature().Recv() == nil {
		return false
	}

	_, pointer := f.Signature().Recv().Type().(*types.Pointer)

	return pointer
}
