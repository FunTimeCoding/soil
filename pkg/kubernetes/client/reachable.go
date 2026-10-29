package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/refusal"
)

func (c *Client) Reachable(x context.Context) error {
	for _, l := range c.selectClients(nil) {
		_, e := l.client.Discovery().ServerVersionWithContext(x)

		if e != nil {
			return refusal.Classify(e)
		}
	}

	return nil
}
