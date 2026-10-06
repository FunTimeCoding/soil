package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) pageWithStatus(
	identifier string,
	status string,
) (*page.Page, error) {
	u := c.basic.Base().Copy().Path(
		"%s/%s",
		constant.ConfluencePage,
		identifier,
	).Set(constant.ConfluenceBodyFormat, constant.ConfluenceStorageFormat)

	if status != "" {
		u = u.Set(constant.ConfluenceStatus, status)
	}

	var result *response.Page

	if e := c.basic.Get(u.String(), &result); e != nil {
		return nil, e
	}

	return page.New(result, c.host), nil
}
