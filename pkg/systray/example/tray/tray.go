//go:build local

package tray

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/gogpu/systray"
	"os"
	"time"
)

func Tray() {
	t := systray.New()
	menu := systray.NewMenu()
	menu.Add(system.Hostname(), func() {}).SetDisabled(true)
	start := time.Now()
	menu.Add(start.Format("started 15:04:05"), func() {}).SetDisabled(true)
	uptime := menu.Add("up 0m", func() {})
	uptime.SetDisabled(true)
	go watchUptime(uptime, start)
	menu.AddSeparator()
	menu.Add(
		"Quit",
		func() {
			t.Remove()
			os.Exit(0)
		},
	)
	t.SetIcon(icon()).SetTooltip("soil tray example").SetMenu(menu)
	t.Show()
	errors.PanicOnError(t.Run())
}
