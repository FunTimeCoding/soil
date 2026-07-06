package example

import (
	"github.com/funtimecoding/soil/pkg/system/debian"
	"runtime"
)

func Download() {
	debian.CheckLatestImage()
	debian.New().DownloadImage(debian.Bookworm, runtime.GOARCH)
}
