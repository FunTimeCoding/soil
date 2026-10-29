package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/refusal"
	"github.com/prometheus/alertmanager/api/v2/client/general"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) Status() (*models.AlertmanagerStatus, error) {
	r, e := c.client.General.GetStatus(&general.GetStatusParams{})

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return r.GetPayload(), nil
}
