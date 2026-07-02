//go:build local

package tray

import (
	"fmt"
	"github.com/getlantern/systray"
)

func Tray() {
	fmt.Println("Start")
	systray.Run(onReady, onExit)
	fmt.Println("After run")
}
