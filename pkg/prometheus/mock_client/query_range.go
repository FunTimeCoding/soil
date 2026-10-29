package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

func (c *Client) QueryRange(
	_ string,
	_ v1.Range,
) (*query_result.Result, error) {
	return nil, c.Failure
}
