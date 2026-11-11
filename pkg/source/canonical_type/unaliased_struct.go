package canonical_type

import "go/types"

func unaliasedStruct(s *types.Struct) *types.Struct {
	fields := make([]*types.Var, s.NumFields())
	tags := make([]string, s.NumFields())

	for i := range s.NumFields() {
		f := s.Field(i)
		fields[i] = types.NewField(
			f.Pos(),
			f.Pkg(),
			f.Name(),
			unaliased(f.Type()),
			f.Embedded(),
		)
		tags[i] = s.Tag(i)
	}

	return types.NewStruct(fields, tags)
}
