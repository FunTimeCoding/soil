package census

import "go/types"

func sameType(
	a *types.Named,
	b *types.Named,
) bool {
	return qualified(a) == qualified(b)
}
