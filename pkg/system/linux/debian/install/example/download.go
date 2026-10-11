package example

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/install"
	"runtime"
)

func Download() {
	install.CheckLatestImage()
	install.New().DownloadImage(constant.Bookworm, runtime.GOARCH)
}
