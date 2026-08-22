package basic

import "github.com/funtimecoding/soil/pkg/web"

func (c *Client) SelfSigned() {
	c.client = web.InsecureClient()
}
