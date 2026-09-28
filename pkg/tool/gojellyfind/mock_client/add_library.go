package mock_client

import "github.com/funtimecoding/soil/pkg/jellyfin/library"

func (c *Client) AddLibrary(l *library.Library) {
	c.libraries = append(c.libraries, l)
}
