package face

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/types"
)

func (s *Set) Satisfied(o types.Object) []*fact.Interface {
	f, isFunction := o.(*types.Func)

	if !isFunction {
		return nil
	}

	signature, isSignature := f.Type().(*types.Signature)

	if !isSignature || signature.Recv() == nil {
		return nil
	}

	r := signature.Recv().Type()

	if i, isPointer := r.(*types.Pointer); isPointer {
		r = i.Elem()
	}

	named, isNamed := r.(*types.Named)

	if !isNamed {
		return nil
	}

	candidates := s.byMethod[f.Name()]

	if len(candidates) == 0 {
		return nil
	}

	methods := methodSignatures(named)
	var result []*fact.Interface

	for _, c := range candidates {
		if satisfies(methods, c.Methods) {
			result = append(result, c)
		}
	}

	return result
}
