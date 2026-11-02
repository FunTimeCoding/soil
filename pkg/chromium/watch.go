package chromium

import "github.com/funtimecoding/soil/pkg/chromium/event"

func (c *Client) Watch(observe func(*event.Event)) error {
	if e := c.ready(); e != nil {
		return e
	}

	return c.subscribe(observe)
}
