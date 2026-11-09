package prometheus

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"time"
)

func (c *Client) AllLabels() ([]string, error) {
	result, e := c.LabelNames([]string{}, soil.Epoch, time.Time{})

	if e != nil {
		return nil, e
	}

	return result.Values, nil
}
