package jenkins

import "github.com/funtimecoding/soil/pkg/jenkins/basic"

func (c *Client) Basic() *basic.Client {
	return c.basic
}
