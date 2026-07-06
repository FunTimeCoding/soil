package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (c *Client) DraftPages() ([]*page.Page, error) {
	var result []*page.Page

	for _, e := range c.pages {
		if e.deleted || e.page == nil {
			continue
		}

		if e.page.Status == "draft" {
			result = append(result, toPage(e.page))
		}
	}

	return result, nil
}
