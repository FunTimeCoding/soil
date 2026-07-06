package opsgenie

import "github.com/funtimecoding/soil/pkg/face"

func (c *Client) DescriptionToName(f face.StringAlias) {
	c.descriptionToName = f
}
