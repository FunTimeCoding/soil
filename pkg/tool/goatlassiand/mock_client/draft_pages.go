package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (c *Client) DraftPages() ([]*page.Page, error) {
	var result []*page.Page

	for _, e := range c.pages {
		if e.Deleted || e.Page == nil {
			continue
		}

		if e.Page.Status == "draft" {
			result = append(result, toPage(e.Page))
		}
	}

	return result, nil
}
