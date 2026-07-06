package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/iterm/screen"
)

func (c *Client) ReadScreen(identifier string) (*screen.Screen, error) {
	s, okay := c.screens[identifier]

	if !okay {
		return screen.Stub(), fmt.Errorf(
			"session %s not found",
			identifier,
		)
	}

	return s, nil
}
