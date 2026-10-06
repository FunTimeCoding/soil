package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) Labels() ([]*response.LabelResult, error) {
	var r *response.Labels

	if e := c.basic.GetV2Path(constant.ConfluenceLabel, &r); e != nil {
		return nil, e
	}

	return r.Results, nil
}
