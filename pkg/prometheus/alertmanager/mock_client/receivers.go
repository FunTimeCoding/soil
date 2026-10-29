package mock_client

import "github.com/prometheus/alertmanager/api/v2/models"

func (c *Client) Receivers() ([]*models.Receiver, error) {
	return nil, c.failure
}
