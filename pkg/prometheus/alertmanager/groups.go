package alertmanager

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/refusal"
	"github.com/prometheus/alertmanager/api/v2/models"
)

func (c *Client) Groups() (models.AlertGroups, error) {
	response, e := c.client.Alertgroup.GetAlertGroups(nil, nil)

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return response.GetPayload(), nil
}
