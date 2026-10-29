package mock_client

import "github.com/prometheus/client_golang/api/prometheus/v1"

func (c *Client) Targets() (v1.TargetsResult, error) {
	return v1.TargetsResult{}, c.Failure
}
