package confluence

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/basic"

func (c *Client) Basic() *basic.Client {
	return c.basic
}
