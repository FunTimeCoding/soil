package mock_client

import "github.com/funtimecoding/soil/pkg/jellyfin/library"

func (c *Client) Libraries() ([]*library.Library, error) {
	return c.libraries, nil
}
