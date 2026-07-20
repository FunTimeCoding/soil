package example

import (
	"github.com/funtimecoding/soil/pkg/system/debian"
	"github.com/funtimecoding/soil/pkg/system/debian/constant"
	"runtime"
)

func Netboot() {
	debian.New().DownloadNetboot(constant.Bookworm, runtime.GOARCH)
}
