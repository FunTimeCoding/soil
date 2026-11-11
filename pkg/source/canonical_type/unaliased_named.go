package canonical_type

import "go/types"

func unaliasedNamed(n *types.Named) types.Type {
	arguments := n.TypeArgs()

	if arguments.Len() == 0 {
		return n
	}

	resolved := make([]types.Type, arguments.Len())

	for i := range arguments.Len() {
		resolved[i] = unaliased(arguments.At(i))
	}

	result, e := types.Instantiate(nil, n.Origin(), resolved, false)

	if e != nil {
		return n
	}

	return result
}
