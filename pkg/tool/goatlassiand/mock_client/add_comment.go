package mock_client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) AddComment(
	pageIdentifier string,
	body string,
) error {
	e, okay := c.pages[pageIdentifier]

	if !okay || e.deleted {
		return not_found.New("page", pageIdentifier)
	}

	return nil
}
