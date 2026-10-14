package unclosed_resource

import "go/types"

func isInterfaceMethod(f *types.Func) bool {
	_, result := f.Signature().Recv().Type().Underlying().(*types.Interface)

	return result
}
