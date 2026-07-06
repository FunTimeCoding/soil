package opsgenie

import "github.com/funtimecoding/soil/pkg/face"

func (c *Client) ShortAlert(f face.StringAlias) {
	c.shortAlert = f
}
