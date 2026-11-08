package preposition_name

import "go/types"

func isBool(t types.Type) bool {
	b, okay := t.Underlying().(*types.Basic)

	return okay && b.Kind() == types.Bool
}
