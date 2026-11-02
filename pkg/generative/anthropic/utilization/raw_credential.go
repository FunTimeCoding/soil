package utilization

import (
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"runtime"
)

func rawCredential() string {
	switch runtime.GOOS {
	case systemConstant.Darwin:
		return keychainRaw()
	case systemConstant.Linux:
		return fileRaw()
	default:
		return ""
	}
}
