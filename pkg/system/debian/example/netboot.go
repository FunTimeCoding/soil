package example

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/debian"
	"runtime"
)

func Netboot() {
	debian.New().DownloadNetboot(constant.Bookworm, runtime.GOARCH)
}
