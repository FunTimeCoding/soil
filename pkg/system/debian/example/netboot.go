package example

import (
	"github.com/funtimecoding/soil/pkg/system/debian"
	"runtime"
)

func Netboot() {
	debian.New().DownloadNetboot(debian.Bookworm, runtime.GOARCH)
}
