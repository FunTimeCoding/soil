package unguarded_peer

import "go/types"

func takesGuard(s *types.Signature) bool {
	for i := range s.Params().Len() {
		if isGuard(s.Params().At(i).Type()) {
			return true
		}
	}

	return false
}
