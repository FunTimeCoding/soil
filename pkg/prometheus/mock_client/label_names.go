package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (c *Client) LabelNames(
	_ []string,
	since time.Time,
) (*label_result.Result, error) {
	c.record(since)

	if c.Failure != nil {
		return nil, c.Failure
	}

	return label_result.New([]string{}, nil, nil), nil
}
