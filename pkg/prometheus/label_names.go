package prometheus

import (
	"github.com/funtimecoding/soil/pkg/prometheus/label_result"
	"time"
)

func (c *Client) LabelNames(
	matches []string,
	since time.Time,
) (*label_result.Result, error) {
	v, w, e := c.client.LabelNames(
		c.context,
		matches,
		since,
		time.Now(),
	)

	if e != nil {
		return nil, e
	}

	return label_result.New(v, w), nil
}
