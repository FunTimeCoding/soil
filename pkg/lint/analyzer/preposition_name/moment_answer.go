package preposition_name

import "go/types"

func momentAnswer(s *types.Signature) bool {
	r := s.Results()

	if r.Len() == 0 || !isTime(r.At(0).Type()) {
		return false
	}

	next := 1

	if next < r.Len() && isBool(r.At(next).Type()) {
		next++
	}

	if next < r.Len() && isError(r.At(next).Type()) {
		next++
	}

	return next == r.Len()
}
