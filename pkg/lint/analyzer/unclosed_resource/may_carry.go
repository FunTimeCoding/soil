package unclosed_resource

import "go/types"

func mayCarry(t types.Type) bool {
	if _, parameter := t.(*types.TypeParam); parameter {
		return true
	}

	return carriesObligation(t)
}
