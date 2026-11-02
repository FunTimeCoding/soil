package utilization

import (
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"runtime"
)

func Supported() bool {
	return runtime.GOOS == systemConstant.Darwin || runtime.GOOS == systemConstant.Linux
}
