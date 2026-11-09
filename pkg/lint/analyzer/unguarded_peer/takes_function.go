package unguarded_peer

import "go/types"

func takesFunction(s *types.Signature) bool {
	for i := range s.Params().Len() {
		_, okay := s.Params().At(i).Type().Underlying().(*types.Signature)

		if okay {
			return true
		}
	}

	return false
}
