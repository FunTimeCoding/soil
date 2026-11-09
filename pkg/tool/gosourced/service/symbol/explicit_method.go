package symbol

import "go/types"

func explicitMethod(
	i *types.Interface,
	name string,
) *types.Func {
	for m := range i.ExplicitMethods() {
		if m.Name() == name {
			return m
		}
	}

	return nil
}
