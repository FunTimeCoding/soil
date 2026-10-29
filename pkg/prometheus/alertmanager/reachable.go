package alertmanager

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/refusal"
	"github.com/prometheus/alertmanager/api/v2/client/general"
)

func (c *Client) Reachable(x context.Context) error {
	r, e := c.client.General.GetStatus(general.NewGetStatusParamsWithContext(x))

	if e != nil {
		return refusal.Classify(e)
	}

	if r.GetPayload() == nil {
		return unexpected.Format("alertmanager status without payload")
	}

	return nil
}
