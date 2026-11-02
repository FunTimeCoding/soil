package mock_client

import "github.com/funtimecoding/soil/pkg/provision/types/change"

func (c *Client) PostTrigger(v *change.Change) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.calls++

	if e := c.failures.Caller(); e != nil {
		return e
	}

	c.triggered = append(c.triggered, v)

	return nil
}
