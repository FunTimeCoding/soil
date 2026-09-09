package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/query_result"
	"time"
)

func (c *Client) Query(
	_ string,
	_ time.Time,
) (*query_result.Result, error) {
	return nil, nil
}
