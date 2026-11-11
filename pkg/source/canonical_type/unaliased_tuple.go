package canonical_type

import "go/types"

func unaliasedTuple(t *types.Tuple) *types.Tuple {
	result := make([]*types.Var, t.Len())

	for i := range t.Len() {
		v := t.At(i)
		result[i] = types.NewParam(
			v.Pos(),
			v.Pkg(),
			v.Name(),
			unaliased(v.Type()),
		)
	}

	return types.NewTuple(result...)
}
