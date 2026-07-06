package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) pageWithStatus(
	identifier string,
	status string,
) (*page.Page, error) {
	u := c.basic.Base().Copy().Path(
		"%s/%s",
		constant.Page,
		identifier,
	).Set(constant.BodyFormat, constant.StorageFormat)

	if status != "" {
		u = u.Set(constant.Status, status)
	}

	body, e := c.basic.GetV2(u.String())

	if e != nil {
		return nil, e
	}

	var result *response.Page
	notation.MustDecode(body, &result, false)

	return page.New(result, c.host), nil
}
