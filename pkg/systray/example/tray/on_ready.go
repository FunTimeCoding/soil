//go:build local

package tray

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func onReady() {
	fmt.Println("onReady")
	systray.SetIcon(icon.Data)
	systray.SetTitle("Example Title")
	systray.SetTooltip("Example Tooltip")
	mQuit := systray.AddMenuItem("Quit", "Quit application")
	mQuit.SetIcon(icon.Data)
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}
