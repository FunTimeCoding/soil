package opsgenie

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/constant"

func (c *Client) ParseDescription(f constant.ParseDescription) {
	c.parseDescription = f
}
