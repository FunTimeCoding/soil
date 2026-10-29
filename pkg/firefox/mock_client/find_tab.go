package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/firefox/tab"
)

func (c *Client) findTab(identifier int) (int, *tab.Tab, error) {
	for i, t := range c.tabs {
		if t.Identifier == identifier {
			return i, t, nil
		}
	}

	return -1, nil, not_found.New("tab", identifier)
}
