package directory

import "github.com/funtimecoding/soil/pkg/directory/constant"

func (c *Client) WithInsecure() *Client {
	c.insecure = true

	if c.port == constant.SecurePort {
		c.port = constant.InsecurePort
	}

	return c
}
