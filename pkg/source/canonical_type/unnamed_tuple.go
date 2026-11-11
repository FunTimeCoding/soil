package canonical_type

import (
	"go/token"
	"go/types"
)

func unnamedTuple(t *types.Tuple) *types.Tuple {
	result := make([]*types.Var, t.Len())

	for i := range t.Len() {
		result[i] = types.NewParam(token.NoPos, nil, "", t.At(i).Type())
	}

	return types.NewTuple(result...)
}
