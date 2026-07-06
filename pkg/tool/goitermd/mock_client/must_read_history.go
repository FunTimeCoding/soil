package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/iterm/history"
)

func (c *Client) MustReadHistory(
	identifier string,
	count int,
) *history.History {
	result, e := c.ReadHistory(identifier, count)
	errors.PanicOnError(e)

	return result
}
