package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func hasCloseMethod(t types.Type) bool {
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

	if !okay {
		return false
	}

	if signature.Params().Len() != 0 {
		return false
	}

	if signature.Results().Len() == 0 {
		return true
	}

	return signature.Results().Len() == 1 &&
		signature.Results().At(0).Type().String() == "error"
}
