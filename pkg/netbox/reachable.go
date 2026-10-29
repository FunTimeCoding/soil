package netbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) Reachable(x context.Context) error {
	_, _, e := c.client.StatusAPI.StatusRetrieve(x).Execute()

	return refusal.Classify(e)
}
