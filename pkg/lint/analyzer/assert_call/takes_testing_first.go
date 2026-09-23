package assert_call

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func takesTestingFirst(f *types.Func) bool {
	s, okay := f.Type().(*types.Signature)

	if !okay || s.Params().Len() == 0 {
		return false
	}

	return s.Params().At(0).Type().String() == constant.AssertTestingType
}
