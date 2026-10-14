package unclosed_resource

import "go/types"

func elementType(t types.Type) types.Type {
	switch v := t.Underlying().(type) {
	case *types.Slice:
		return v.Elem()
	case *types.Array:
		return v.Elem()
	case *types.Map:
		return v.Elem()
	case *types.Chan:
		return v.Elem()
	}

	return nil
}
