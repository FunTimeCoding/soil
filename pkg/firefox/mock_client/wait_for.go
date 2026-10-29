package mock_client

import "github.com/funtimecoding/soil/pkg/firefox/tab"

func (c *Client) WaitFor(
	identifier int,
	_ string,
	_ int,
) (*tab.Tab, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if e := c.failures.Caller(); e != nil {
		return tab.Stub(), e
	}

	_, t, f := c.findTab(identifier)

	if f != nil {
		return tab.Stub(), f
	}

	return t, nil
}
