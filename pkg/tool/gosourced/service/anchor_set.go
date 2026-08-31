package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"go/types"
)

func anchorSet(
	declaration types.Object,
	receiver string,
) ([]types.Object, error) {
	name, okay := declaration.(*types.TypeName)

	if !okay || receiver != "" {
		return []types.Object{declaration}, nil
	}

	named, valid := name.Type().(*types.Named)

	if !valid {
		return []types.Object{declaration}, nil
	}

	set := types.NewMethodSet(types.NewPointer(named))
	var result []types.Object

	for method := range set.Methods() {
		result = append(result, method.Obj())
	}

	if len(result) == 0 {
		return nil, validation.New(
			"type %s has no methods to anchor on",
			name.Name(),
		)
	}

	return result, nil
}
