package opsgenie

import "github.com/funtimecoding/soil/pkg/face"

func (c *Client) ShortUser(f face.StringAlias) {
	c.shortUser = f
}
