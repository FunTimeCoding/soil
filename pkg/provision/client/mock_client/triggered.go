package mock_client

import "github.com/funtimecoding/soil/pkg/provision/types/change"

func (c *Client) Triggered() []*change.Change {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return append([]*change.Change{}, c.triggered...)
}
