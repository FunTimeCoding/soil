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
		run.New().Launch("xdg-open", locator)
	case constant.Darwin:
		run.New().Launch("open", locator)
	default:
		unexpected.String(p)
	}
}
