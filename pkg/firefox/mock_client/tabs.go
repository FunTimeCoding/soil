package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/tab"

func (c *Client) Tabs() ([]*tab.Tab, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return c.tabs, nil
}
