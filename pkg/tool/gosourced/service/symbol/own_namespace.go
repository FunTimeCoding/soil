package symbol

import "go/types"

func ownNamespace(o types.Object) bool {
	switch v := o.(type) {
	case *types.Var:
		return v.IsField()
	case *types.Func:
		return v.Type().(*types.Signature).Recv() != nil
	case *types.Label:
		return true
	}

	return false
}
