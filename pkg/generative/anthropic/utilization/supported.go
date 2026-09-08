package utilization

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"runtime"
)

func Supported() bool {
	return runtime.GOOS == constant.Darwin || runtime.GOOS == constant.Linux
}
