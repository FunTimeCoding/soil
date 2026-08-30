package omit_empty_zero

import "go/types"

func isNumericOrBool(t types.Type) bool {
	if t == nil {
		return false
	}

	basic, okay := t.Underlying().(*types.Basic)

	if !okay {
		return false
	}

	return basic.Info()&(types.IsNumeric|types.IsBoolean) != 0
}
