package service

import "go/types"

func structField(
	named *types.Named,
	name string,
) *types.Var {
	structure, okay := named.Underlying().(*types.Struct)

	if !okay {
		return nil
	}

	for i := range structure.NumFields() {
		if f := structure.Field(i); f.Name() == name {
			return f
		}
	}

	return nil
}
