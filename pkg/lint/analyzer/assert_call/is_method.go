package assert_call

import "go/types"

func isMethod(f *types.Func) bool {
	s, okay := f.Type().(*types.Signature)

	return okay && s.Recv() != nil
}
