package mock_client

import "github.com/funtimecoding/soil/pkg/chromium/event"

func (c *Client) Watch(_ func(*event.Event)) error {
	return c.Failure
}
