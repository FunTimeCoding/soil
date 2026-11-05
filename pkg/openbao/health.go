package openbao

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/openbao/refusal"
	"github.com/funtimecoding/soil/pkg/openbao/status"
)

func (c *Client) Health() (*status.Status, error) {
	h, e := c.client.Sys().Health()

	if e != nil {
		return nil, fmt.Errorf("read health: %w", refusal.Classify(e))
	}

	return status.New(h.Initialized, h.Sealed, h.Version), nil
}
