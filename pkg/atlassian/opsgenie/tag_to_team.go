package opsgenie

import "github.com/funtimecoding/soil/pkg/face"

func (c *Client) TagToTeam(f face.SliceAlias) {
	c.tagToTeam = f
}
