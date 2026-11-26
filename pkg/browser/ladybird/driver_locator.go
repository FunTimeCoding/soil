package ladybird

import "fmt"

func (c *Client) driverLocator(route string) string {
	return fmt.Sprintf("http://%s%s", c.address, route)
}
