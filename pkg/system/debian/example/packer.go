package example

import (
	"github.com/funtimecoding/soil/pkg/system/debian"
	"runtime"
)

func Packer() {
	debian.New().Packer(debian.Bookworm, runtime.GOARCH)
}
