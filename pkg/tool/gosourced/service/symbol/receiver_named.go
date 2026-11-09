package symbol

import "go/types"

func receiverNamed(f *types.Func) *types.Named {
	t := f.Type().(*types.Signature).Recv().Type()

	if p, isPointer := t.(*types.Pointer); isPointer {
		t = p.Elem()
	}

	n, _ := t.(*types.Named)

	return n
}
