package client

import "github.com/funtimecoding/soil/pkg/face"

func (c *Client) WithReporter(r face.Reporter) *Client {
	c.reporter = r

	return c
}
