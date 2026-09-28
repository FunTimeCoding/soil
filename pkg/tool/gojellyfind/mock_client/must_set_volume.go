package mock_client

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustSetVolume(
	sessionIdentifier string,
	level int,
) {
	errors.PanicOnError(c.SetVolume(sessionIdentifier, level))
}
