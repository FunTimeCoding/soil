package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) DraftOverlay(identifier string) (*page.Page, error) {
	var result *response.Page

	if e := c.basic.Get(
		c.basic.Base().Copy().Path(
			"%s/%s",
			constant.ConfluencePage,
			identifier,
		).Set(
			constant.ConfluenceBodyFormat,
			constant.ConfluenceStorageFormat,
		).
			Set(constant.ConfluenceGetDraft, "true").String(),
		&result,
	); e != nil {
		return nil, e
	}

	return page.New(result, c.host), nil
}
