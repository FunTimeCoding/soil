package example

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/debian"
	"runtime"
)

func Packer() {
	debian.New().Packer(constant.Bookworm, runtime.GOARCH)
}
