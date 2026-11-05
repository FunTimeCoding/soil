package mock_client

import "github.com/funtimecoding/soil/pkg/jellyfin/library"

func (c *Client) Libraries() ([]*library.Library, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.libraries, nil
}
