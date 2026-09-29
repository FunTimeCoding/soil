package linkace

import "fmt"

func (c *Client) DeleteTag(identifier int) error {
	return c.basic.Delete(fmt.Sprintf("tags/%d", identifier))
}
