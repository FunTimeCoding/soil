package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) PagesByLabel(labelIdentifier string) ([]*page.Page, error) {
	var result *response.Pages

	if e := c.basic.Get(
		c.basic.Base().Copy().Path(
			"%s/%s%s",
			constant.ConfluenceLabel,
			labelIdentifier,
			constant.ConfluencePage,
		).Set(
			constant.ConfluenceBodyFormat,
			constant.ConfluenceStorageFormat,
		).String(),
		&result,
	); e != nil {
		return nil, e
	}

	return page.NewSlice(result.Results, c.host), nil
}
