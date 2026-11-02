package keepass

import "os"

func (c *Client) Changed() bool {
	i, e := os.Stat(c.path)

	return e != nil || !i.ModTime().Equal(c.loaded)
}
