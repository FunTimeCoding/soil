package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) Labels() ([]*response.LabelResult, error) {
	body, e := c.basic.GetV2Path(constant.Label)

	if e != nil {
		return nil, e
	}

	var r *response.Labels
	notation.MustDecode(body, &r, false)

	return r.Results, nil
}
