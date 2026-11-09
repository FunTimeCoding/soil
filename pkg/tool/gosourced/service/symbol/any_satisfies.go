package symbol

import "go/types"

func anySatisfies(
	concrete []*types.Named,
	i *types.Interface,
) bool {
	for _, t := range concrete {
		if implementsAny(t, []*types.Interface{i}) {
			return true
		}
	}

	return false
}
