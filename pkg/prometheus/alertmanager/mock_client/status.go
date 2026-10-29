package mock_client

import "github.com/prometheus/alertmanager/api/v2/models"

func (c *Client) Status() (*models.AlertmanagerStatus, error) {
	return nil, c.failure
}
