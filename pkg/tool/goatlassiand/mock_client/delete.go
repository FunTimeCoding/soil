package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) Delete(pageIdentifier string) error {
	e, okay := c.pages[pageIdentifier]

	if !okay || e.Deleted || e.Page == nil || e.Page.Status != "current" {
		return not_found.New("page", pageIdentifier)
	}

	e.Deleted = true

	return nil
}
