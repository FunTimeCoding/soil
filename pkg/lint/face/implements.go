package face

import "go/types"

func (s *Set) Implements(o types.Object) bool {
	f, isFunction := o.(*types.Func)

	if !isFunction {
		return false
	}

	signature, isSignature := f.Type().(*types.Signature)

	if !isSignature || signature.Recv() == nil {
		return false
	}

	r := signature.Recv().Type()

	if i, isPointer := r.(*types.Pointer); isPointer {
		r = i.Elem()
	}

	named, isNamed := r.(*types.Named)

	if !isNamed {
		return false
	}

	for _, candidate := range s.byMethod[f.Name()] {
		if types.Implements(named, candidate) ||
			types.Implements(types.NewPointer(named), candidate) {
			return true
		}
	}

	return false
}
