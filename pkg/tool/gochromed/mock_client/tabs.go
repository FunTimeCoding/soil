package mock_client

import "github.com/funtimecoding/soil/pkg/chromium/tab"

func (c *Client) Tabs() []*tab.Tab {
	return c.tabs
}
