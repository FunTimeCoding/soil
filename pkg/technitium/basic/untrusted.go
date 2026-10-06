package basic

import "github.com/funtimecoding/soil/pkg/web"

func (c *Client) Untrusted() {
	c.requester.WithClient(web.InsecureStallClient())
}
