package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) Delete(pageIdentifier string) error {
	e, okay := c.pages[pageIdentifier]

	if !okay || e.deleted || e.page == nil || e.page.Status != "current" {
		return not_found.New("page", pageIdentifier)
	}

	e.deleted = true

	return nil
}
