package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) SetPageStatus(
	identifier string,
	status string,
) (*page.Page, error) {
	var current *page.Page
	var e error

	if status == constant.ConfluenceCurrentStatus {
		current, e = c.DraftOverlay(identifier)
	} else {
		current, e = c.Page(identifier)
	}

	if e != nil {
		return nil, e
	}

	version := 1

	if status == constant.ConfluenceCurrentStatus {
		published, f := c.Page(identifier)

		if f == nil && published.Raw.Status == constant.ConfluenceCurrentStatus {
			version = published.Raw.Version.Number + 1
		}
	}

	return c.PutPage(
		identifier,
		current.Name,
		current.Raw.Body.Storage.Value,
		version,
		"",
		status,
	)
}
