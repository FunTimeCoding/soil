package monitor_tester

import "charm.land/bubbletea/v2"

func Run(c tea.Cmd) {
	if c == nil {
		return
	}

	if b, okay := c().(tea.BatchMsg); okay {
		for _, n := range b {
			Run(n)
		}
	}
}
