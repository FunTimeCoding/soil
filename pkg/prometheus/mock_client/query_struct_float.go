package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic/float"
	"time"
)

func (c *Client) QueryStructFloat(
	_ string,
	_ float64,
	_ time.Time,
) (*float.Result, error) {
	return nil, c.Failure
}
