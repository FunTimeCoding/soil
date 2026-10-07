package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) DeleteDraft(pageIdentifier string) error {
	e, okay := c.pages[pageIdentifier]

	if !okay || e.Deleted {
		return not_found.New("page", pageIdentifier)
	}

	if e.Page != nil && e.Page.Status == "draft" {
		e.Deleted = true

		return nil
	}

	return not_found.New("page", pageIdentifier)
}
