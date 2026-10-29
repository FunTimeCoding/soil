package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (c *Client) LabelValues(
	_ string,
	_ []string,
	_ time.Time,
) (*label_result.Result, error) {
	return nil, c.Failure
}
