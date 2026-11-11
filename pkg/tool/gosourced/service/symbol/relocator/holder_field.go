package relocator

import "go/types"

func holderField(
	source *types.Named,
	destination *types.Named,
) *types.Var {
	s, okay := source.Underlying().(*types.Struct)

	if !okay {
		return nil
	}

	var result *types.Var

	for f := range s.Fields() {
		if !types.Identical(f.Type(), destination) &&
			!types.Identical(f.Type(), types.NewPointer(destination)) {
			continue
		}

		if result != nil {
			return nil
		}

		result = f
	}

	return result
}
