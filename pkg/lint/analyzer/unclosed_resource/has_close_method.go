package unclosed_resource

import "go/types"

func hasCloseMethod(t types.Type) bool {
	method, _, _ := types.LookupFieldOrMethod(t, true, nil, "Close")

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
