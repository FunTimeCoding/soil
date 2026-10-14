package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func parameterIndices(f *types.Func) map[*types.Var]int {
	signature := f.Signature()
	result := make(map[*types.Var]int)

	if r := signature.Recv(); r != nil && r.Name() != "" {
		result[r] = constant.ReceiverArgument
	}

	for i := range signature.Params().Len() {
		result[signature.Params().At(i)] = i
	}

	return result
}
