package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (c *Client) LabelValues(
	_ string,
	_ []string,
	start time.Time,
	end time.Time,
) (*label_result.Result, error) {
	c.record(start, end)

	if c.Failure != nil {
		return nil, c.Failure
	}

	return label_result.New([]string{}, nil, nil), nil
}
