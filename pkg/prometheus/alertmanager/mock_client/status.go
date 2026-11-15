package mock_client

import "github.com/prometheus/alertmanager/api/v2/models"

func (c *Client) Status() (*models.AlertmanagerStatus, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return &models.AlertmanagerStatus{}, nil
}
