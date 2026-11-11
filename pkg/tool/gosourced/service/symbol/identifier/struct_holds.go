package identifier

import "go/types"

func structHolds(
	s *types.Struct,
	field *types.Var,
) bool {
	for f := range s.Fields() {
		if f.Pos() == field.Pos() {
			return true
		}
	}

	return false
}
