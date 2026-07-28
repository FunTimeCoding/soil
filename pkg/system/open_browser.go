package system

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"runtime"
)

func OpenBrowser(locator string) {
	switch p := runtime.GOOS; p {
	case constant.Linux:
		run.New().Launch(constant.XdgOpenCommand, locator)
	case constant.Darwin:
		run.New().Launch(constant.OpenCommand, locator)
	default:
		unexpected.String(p)
	}
}
