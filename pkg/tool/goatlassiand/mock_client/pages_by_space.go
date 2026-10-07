package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) PagesBySpace(
	identifier string,
	status string,
) ([]*page.Page, error) {
	if status == "" {
		status = constant.ConfluenceCurrentStatus
	}

	var result []*page.Page

	for _, e := range c.pages {
		if e.Deleted || e.Page == nil {
			continue
		}

		if e.Page.SpaceIdentifier != identifier {
			continue
		}

		if e.Page.Status != status {
			continue
		}

		result = append(result, toPage(e.Page))
	}

	return result, nil
}
