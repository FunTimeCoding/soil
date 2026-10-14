package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func closeIndex(holder *types.Named) []int {
	pointer := types.NewPointer(holder)

	if !hasCloseMethod(pointer) {
		return nil
	}

	_, result, _ := types.LookupFieldOrMethod(
		pointer,
		true,
		holder.Obj().Pkg(),
		constant.CloseMethod,
	)

	return result
}
