package receive

import (
	"charm.land/bubbletea/v2"
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/client"
)

func Command(c *client.Client) tea.Cmd {
	return func() tea.Msg {
		if l := c.Consume(); len(l) > 0 {
			return Message{Text: l}
		}

		return nil
	}
}
