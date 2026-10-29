package mock_client

import "github.com/prometheus/alertmanager/api/v2/models"

func (c *Client) Groups() (models.AlertGroups, error) {
	return nil, c.failure
}
