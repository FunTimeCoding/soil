package keepass

import "github.com/funtimecoding/soil/pkg/system"

func (c *Client) Changed() bool {
	return !system.Stat(c.path).ModTime().Equal(c.loaded)
}
