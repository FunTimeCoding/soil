package example

import (
	"github.com/funtimecoding/soil/pkg/system/debian"
	"github.com/funtimecoding/soil/pkg/system/debian/constant"
	"runtime"
)

func Download() {
	debian.CheckLatestImage()
	debian.New().DownloadImage(constant.Bookworm, runtime.GOARCH)
}
