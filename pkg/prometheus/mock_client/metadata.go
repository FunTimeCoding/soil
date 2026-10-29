package mock_client

import "github.com/prometheus/client_golang/api/prometheus/v1"

func (c *Client) Metadata(_ string) (map[string][]v1.Metadata, error) {
	return nil, c.Failure
}
