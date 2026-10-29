package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/refusal"
	"github.com/prometheus/alertmanager/api/v2/client/receiver"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) Receivers() ([]*models.Receiver, error) {
	r, e := c.client.Receiver.GetReceivers(&receiver.GetReceiversParams{})

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return r.GetPayload(), nil
}
