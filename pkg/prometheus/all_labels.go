package prometheus

import soil "github.com/funtimecoding/soil/pkg/constant"

func (c *Client) AllLabels() ([]string, error) {
	result, e := c.LabelNames([]string{}, soil.StartOfTime)

	if e != nil {
		return nil, e
	}

	return result.Values, nil
}
