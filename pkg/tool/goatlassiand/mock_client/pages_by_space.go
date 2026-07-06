package mock_client

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/page"
)

func (c *Client) PagesBySpace(
	identifier string,
	status string,
) ([]*page.Page, error) {
	if status == "" {
		status = constant.CurrentStatus
	}

	var result []*page.Page

	for _, e := range c.pages {
		if e.deleted || e.page == nil {
			continue
		}

		if e.page.SpaceIdentifier != identifier {
			continue
		}

		if e.page.Status != status {
			continue
		}

		result = append(result, toPage(e.page))
	}

	return result, nil
}
