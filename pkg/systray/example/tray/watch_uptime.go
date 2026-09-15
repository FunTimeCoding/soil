//go:build local

package tray

import (
	"fmt"
	"github.com/gogpu/systray"
	"time"
)

func watchUptime(
	m *systray.MenuItem,
	start time.Time,
) {
	for range time.Tick(time.Minute) {
		m.SetLabel(fmt.Sprintf("up %s", time.Since(start).Round(time.Minute)))
	}
}
