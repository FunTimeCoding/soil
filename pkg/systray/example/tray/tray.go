//go:build local

package tray

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/getlantern/systray"
)

func Tray() {
	console.Line("Start")
	systray.Run(onReady, onExit)
	console.Line("After run")
}
