//go:build local

package tray

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func onReady() {
	console.Line("onReady")
	systray.SetIcon(icon.Data)
	systray.SetTitle("Example Title")
	systray.SetTooltip("Example Tooltip")
	mQuit := systray.AddMenuItem("Quit", "Quit application")
	mQuit.SetIcon(icon.Data)
	go func() {
		<-mQuit.ClickedCh
		console.Line("Requesting quit")
		systray.Quit()
		console.Line("Finished quitting")
	}()
}
