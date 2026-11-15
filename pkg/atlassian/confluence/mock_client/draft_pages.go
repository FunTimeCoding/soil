package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/page"

func (c *Client) DraftPages() ([]*page.Page, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if f := c.failures.Caller(); f != nil {
		return nil, f
	}

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
