package canonical_type

import "go/types"

func Unnamed(s *types.Signature) *types.Signature {
	return types.NewSignatureType(
		nil,
		nil,
		nil,
		unnamedTuple(s.Params()),
		unnamedTuple(s.Results()),
		s.Variadic(),
	)
}
