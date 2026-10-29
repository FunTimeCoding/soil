package mock_client

import "github.com/prometheus/client_golang/api/prometheus/v1"

func (c *Client) Status() (v1.RuntimeinfoResult, error) {
	return v1.RuntimeinfoResult{}, c.Failure
}
