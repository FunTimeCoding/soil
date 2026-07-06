package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/tab"

func (c *Client) Tabs() ([]*tab.Tab, error) {
	return c.tabs, nil
}
