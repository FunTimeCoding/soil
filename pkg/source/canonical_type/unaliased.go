package canonical_type

import "go/types"

func unaliased(t types.Type) types.Type {
	switch v := t.(type) {
	case *types.Alias:
		return unaliased(types.Unalias(v))
	case *types.Pointer:
		return types.NewPointer(unaliased(v.Elem()))
	case *types.Slice:
		return types.NewSlice(unaliased(v.Elem()))
	case *types.Array:
		return types.NewArray(unaliased(v.Elem()), v.Len())
	case *types.Map:
		return types.NewMap(unaliased(v.Key()), unaliased(v.Elem()))
	case *types.Chan:
		return types.NewChan(v.Dir(), unaliased(v.Elem()))
	case *types.Signature:
		if v.TypeParams().Len() > 0 {
			return v
		}

		return types.NewSignatureType(
			nil,
			nil,
			nil,
			unaliasedTuple(v.Params()),
			unaliasedTuple(v.Results()),
			v.Variadic(),
		)
	case *types.Named:
		return unaliasedNamed(v)
	case *types.Struct:
		return unaliasedStruct(v)
	case *types.Interface:
		return unaliasedInterface(v)
	default:
		return t
	}
}
