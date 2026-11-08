package census

import "go/types"

func named(t types.Type) *types.Named {
	for {
		switch x := t.(type) {
		case *types.Pointer:
			t = x.Elem()
		case *types.Slice:
			t = x.Elem()
		case *types.Array:
			t = x.Elem()
		case *types.Named:
			if x.Obj().Pkg() == nil {
				return nil
			}

			return x
		default:
			return nil
		}
	}
}
