package unclosed_resource

import "go/types"

func returnsError(t types.Type) bool {
	if isResponse(t) {
		return true
	}

	method, _, _ := types.LookupFieldOrMethod(t, true, nil, "Close")

	if method == nil {
		return false
	}

	signature, okay := method.Type().(*types.Signature)

	return okay && signature.Results().Len() == 1
}
