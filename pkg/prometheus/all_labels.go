package prometheus

import "github.com/funtimecoding/soil/pkg/constant"

func (c *Client) AllLabels() ([]string, error) {
	result, e := c.LabelNames([]string{}, constant.StartOfTime)

	if e != nil {
		return nil, e
	}

	return result.Values, nil
}
