package example

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/install"
	"runtime"
)

func Packer() {
	install.New().Packer(constant.Bookworm, runtime.GOARCH)
}
