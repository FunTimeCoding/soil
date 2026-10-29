package mock_client

import "github.com/prometheus/client_golang/api/prometheus/v1"

func (c *Client) Series() (v1.TSDBResult, error) {
	return v1.TSDBResult{}, c.Failure
}
