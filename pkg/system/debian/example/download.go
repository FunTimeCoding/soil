package example

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/debian"
	"runtime"
)

func Download() {
	debian.CheckLatestImage()
	debian.New().DownloadImage(constant.Bookworm, runtime.GOARCH)
}
