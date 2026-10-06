package service

import "go/types"

func namedStruct(o types.Object) (*types.Named, *types.Struct) {
	t, okay := o.(*types.TypeName)

	if !okay {
		return nil, nil
	}

	named, okay := t.Type().(*types.Named)

	if !okay {
		return nil, nil
	}

	structure, okay := named.Underlying().(*types.Struct)

	if !okay {
		return nil, nil
	}

	return named, structure
}
