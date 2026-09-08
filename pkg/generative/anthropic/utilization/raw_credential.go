package utilization

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"runtime"
)

func rawCredential() string {
	switch runtime.GOOS {
	case constant.Darwin:
		return keychainRaw()
	case constant.Linux:
		return fileRaw()
	default:
		return ""
	}
}
