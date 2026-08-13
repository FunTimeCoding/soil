package alertmanager

import (
	"github.com/prometheus/alertmanager/api/v2/client/receiver"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) Receivers() ([]*models.Receiver, error) {
	r, e := c.client.Receiver.GetReceivers(&receiver.GetReceiversParams{})

	if e != nil {
		return nil, e
	}

	return r.GetPayload(), nil
}
