package preposition_name

import "go/types"

func state(s *types.Signature) bool {
	return s.Params().Len() == 0 &&
		s.Results().Len() == 1 &&
		isBool(s.Results().At(0).Type())
}
