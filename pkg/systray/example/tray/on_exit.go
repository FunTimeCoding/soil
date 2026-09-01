//go:build local

package tray

import "github.com/funtimecoding/soil/pkg/console"

func onExit() {
	console.Line("onExit")
}
