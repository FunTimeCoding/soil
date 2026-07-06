package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustLabels() []*response.LabelResult {
	result, e := c.Labels()
	errors.PanicOnError(e)

	return result
}
