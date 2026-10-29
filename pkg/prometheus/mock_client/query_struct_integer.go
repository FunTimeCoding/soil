package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/result/generic/integer"
	"time"
)

func (c *Client) QueryStructInteger(
	_ string,
	_ int,
	_ time.Time,
) (*integer.Result, error) {
	return nil, c.Failure
}
