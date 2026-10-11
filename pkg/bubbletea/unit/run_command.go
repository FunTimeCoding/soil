package unit

import "charm.land/bubbletea/v2"

func runCommand(c tea.Cmd) {
	if c == nil {
		return
	}

	if b, okay := c().(tea.BatchMsg); okay {
		for _, n := range b {
			runCommand(n)
		}
	}
}
