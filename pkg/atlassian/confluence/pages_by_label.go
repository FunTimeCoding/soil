package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) PagesByLabel(labelIdentifier string) ([]*page.Page, error) {
	body, e := c.basic.GetV2(
		c.basic.Base().Copy().Path(
			"%s/%s%s",
			constant.ConfluenceLabel,
			labelIdentifier,
			constant.ConfluencePage,
		).Set(
			constant.ConfluenceBodyFormat,
			constant.ConfluenceStorageFormat,
		).String(),
	)

	if e != nil {
		return nil, e
	}

	var result *response.Pages
	notation.MustDecode(body, &result, false)

	return page.NewSlice(result.Results, c.host), nil
}
