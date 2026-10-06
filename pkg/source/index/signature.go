package index

import "go/types"

func Signature(f *types.Func) string {
	s := f.Type().(*types.Signature)

	return types.TypeString(
		types.NewSignatureType(
			nil,
			nil,
			nil,
			unnamed(s.Params()),
			unnamed(s.Results()),
			s.Variadic(),
		),
		qualifier,
	)
}
