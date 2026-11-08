package census

import "go/types"

func owner(f *types.Func) *types.Named {
	s, okay := f.Type().(*types.Signature)

	if !okay || s.Recv() == nil {
		return nil
	}

	return named(s.Recv().Type())
}
