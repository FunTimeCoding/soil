package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/iterm/history"
)

func (c *Client) ReadHistory(
	identifier string,
	count int,
) (*history.History, error) {
	if _, okay := c.screens[identifier]; !okay {
		return history.Stub(), not_found.Format(
			"session %s not found",
			identifier,
		)
	}

	result := history.Stub()
	result.Identifier = identifier
	result.TotalScrollback = count

	return result, nil
}
