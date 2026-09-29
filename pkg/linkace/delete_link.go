package linkace

import "fmt"

func (c *Client) DeleteLink(identifier int) error {
	return c.basic.Delete(fmt.Sprintf("links/%d", identifier))
}
