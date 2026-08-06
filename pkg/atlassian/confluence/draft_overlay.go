package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) DraftOverlay(identifier string) (*page.Page, error) {
	body, e := c.basic.GetV2(
		c.basic.Base().Copy().Path(
			"%s/%s",
			constant.ConfluencePage,
			identifier,
		).Set(
			constant.ConfluenceBodyFormat,
			constant.ConfluenceStorageFormat,
		).
			Set(constant.ConfluenceGetDraft, "true").String(),
	)

	if e != nil {
		return nil, e
	}

	var result *response.Page
	notation.MustDecode(body, &result, false)

	return page.New(result, c.host), nil
}
