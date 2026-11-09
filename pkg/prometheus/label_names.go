package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (c *Client) LabelNames(
	matches []string,
	start time.Time,
	end time.Time,
) (*label_result.Result, error) {
	v, w, n, e := c.client.LabelNames(c.context, matches, start, end)

	if e != nil {
		return nil, e
	}

	names := make([]string, len(v))

	for i, name := range v {
		names[i] = string(name)
	}

	return label_result.New(names, w, n), nil
}
