package basic

import "github.com/funtimecoding/soil/pkg/web/stall_client"

func (c *Client) Untrusted() {
	c.requester.WithClient(stall_client.New().Insecure().Build().Client())
}
