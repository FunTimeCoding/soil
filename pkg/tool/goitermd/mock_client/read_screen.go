package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/iterm/screen"
)

func (c *Client) ReadScreen(identifier string) (*screen.Screen, error) {
	s, okay := c.screens[identifier]

	if !okay {
		return screen.Stub(), not_found.Format(
			"session %s not found",
			identifier,
		)
	}

	return s, nil
}
