package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/types"
)

func isCloseMethod(f *types.Func) bool {
	return f != nil && f.Name() == constant.CloseMethod &&
		f.Signature().Recv() != nil && f.Signature().Params().Len() == 0
}
