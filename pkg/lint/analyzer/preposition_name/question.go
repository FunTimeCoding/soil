package preposition_name

import "go/types"

func question(o types.Object) bool {
	switch v := o.(type) {
	case *types.Func:
		r := v.Type().(*types.Signature).Results()

		return r.Len() > 0 &&
			isBool(r.At(0).Type()) &&
			(r.Len() == 1 || r.Len() == 2 && isError(r.At(1).Type()))
	case *types.Var:
		return isBool(v.Type())
	}

	return false
}
