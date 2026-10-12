package mock_client

import (
	"github.com/funtimecoding/soil/pkg/band/response"
	"slices"
)

func (c *Client) PowerState() (*response.PowerState, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	result := response.NewPowerState()
	result.State = c.state
	result.Available = slices.Clone(c.available)

	return result, nil
}
