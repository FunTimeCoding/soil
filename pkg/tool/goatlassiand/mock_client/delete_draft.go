package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) DeleteDraft(pageIdentifier string) error {
	e, okay := c.pages[pageIdentifier]

	if !okay || e.deleted {
		return not_found.New("page", pageIdentifier)
	}

	if e.page != nil && e.page.Status == "draft" {
		e.deleted = true

		return nil
	}

	return not_found.New("page", pageIdentifier)
}
