package linkace

import "fmt"

func (c *Client) DeleteNote(identifier int) error {
	return c.basic.Delete(fmt.Sprintf("notes/%d", identifier))
}
