package prometheus

import "github.com/funtimecoding/soil/pkg/constant"

func (c *Client) AllLabels() []string {
	return c.MustLabelNames([]string{}, constant.StartOfTime).Values
}
