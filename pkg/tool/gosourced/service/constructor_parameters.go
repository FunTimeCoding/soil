package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"go/types"
	"slices"
)

func constructorParameters(
	structure *types.Struct,
	names []string,
) ([]*types.Var, error) {
	var result []*types.Var

	for i := range structure.NumFields() {
		if f := structure.Field(i); slices.Contains(names, f.Name()) {
			result = append(result, f)
		}
	}

	if len(result) < len(names) {
		for _, name := range names {
			if !slices.ContainsFunc(
				result,
				func(f *types.Var) bool { return f.Name() == name },
			) {
				return nil, validation.New("no field %s on the struct", name)
			}
		}
	}

	return result, nil
}
