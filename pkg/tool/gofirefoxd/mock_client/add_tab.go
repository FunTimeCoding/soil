package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/tab"

func (c *Client) AddTab(t *tab.Tab) {
	c.tabs = append(c.tabs, t)
}
