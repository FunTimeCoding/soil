package ollama

import "fmt"

func (c *Client) Locator() string {
	return fmt.Sprintf("%s:%d", c.host, c.port)
}
