package symbol

import "go/types"

func implementsAny(
	t *types.Named,
	interfaces []*types.Interface,
) bool {
	pointer := types.NewPointer(t)

	for _, i := range interfaces {
		if types.Implements(t, i) || types.Implements(pointer, i) {
			return true
		}
	}

	return false
}
