package ollama

import (
	"context"
	"time"
)

func (c *Client) Reachable() bool {
	x, cancel := context.WithTimeout(c.context, 5*time.Second)
	defer cancel()

	return c.client.Heartbeat(x) == nil
}
