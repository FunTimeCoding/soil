package mock_client

import "github.com/funtimecoding/soil/pkg/chromium/tab"

func (c *Client) Tabs() ([]*tab.Tab, error) {
	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	return append([]*tab.Tab{}, c.tabs...), nil
}
