package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/iterm/history"
)

func (c *Client) ReadHistory(
	identifier string,
	_ int,
) (*history.History, error) {
	if _, okay := c.screens[identifier]; !okay {
		return history.Stub(), fmt.Errorf(
			"session %s not found",
			identifier,
		)
	}

	result := history.Stub()
	result.Identifier = identifier

	return result, nil
}
