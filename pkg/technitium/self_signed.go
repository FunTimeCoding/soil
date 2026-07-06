package technitium

import "github.com/funtimecoding/soil/pkg/web"

func (c *Client) SelfSigned() *Client {
	c.client = web.InsecureClient()

	return c
}
