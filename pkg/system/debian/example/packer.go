package example

import (
	"github.com/funtimecoding/soil/pkg/system/debian"
	"github.com/funtimecoding/soil/pkg/system/debian/constant"
	"runtime"
)

func Packer() {
	debian.New().Packer(constant.Bookworm, runtime.GOARCH)
}
