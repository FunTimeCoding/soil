package mock_client

import (
	"slices"
	"time"
)

func (c *Client) Moments() []time.Time {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return slices.Clone(c.moments)
}
