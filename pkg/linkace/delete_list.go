package linkace

import "fmt"

func (c *Client) DeleteList(identifier int) error {
	return c.basic.Delete(fmt.Sprintf("lists/%d", identifier))
}
