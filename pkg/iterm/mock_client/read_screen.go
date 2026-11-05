package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/iterm/screen"
)

func (c *Client) ReadScreen(identifier string) (*screen.Screen, error) {
	if c.Failure != nil {
		return screen.Stub(), c.Failure
	}

	s, okay := c.screens[identifier]

	if !okay {
		return screen.Stub(), not_found.Format(
			"session %s not found",
			identifier,
		)
	}

	return s, nil
}
