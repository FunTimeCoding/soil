package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func returnsError(t types.Type) bool {
	if isResponse(t) {
		return true
	}

	method, _, _ := types.LookupFieldOrMethod(
		t,
		true,
		nil,
		constant.CloseMethod,
	)

	if method == nil {
		return false
	}

	signature, okay := method.Type().(*types.Signature)

	return okay && signature.Results().Len() == 1
}
