package mock_client

import "github.com/funtimecoding/soil/pkg/browser/chromium/event"

func (c *Client) Watch(_ func(*event.Event)) error {
	return c.Failure
}
