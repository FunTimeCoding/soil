package opsgenie

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/detail"

func (c *Client) ParseDescription(f detail.Parser) {
	c.parseDescription = f
}
