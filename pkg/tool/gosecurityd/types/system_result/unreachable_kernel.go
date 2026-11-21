package system_result

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"

func (f *Finding) UnreachableKernel() bool {
	return f.Kernel && f.Kind() == constant.KindPending
}
