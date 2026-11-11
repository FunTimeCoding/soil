package canonical_type

import "go/types"

func String(
	t types.Type,
	q types.Qualifier,
) string {
	return types.TypeString(unaliased(t), q)
}
