package jellyfin

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustPlaybackCommand(
	sessionIdentifier string,
	command string,
	seekPositionTicks int64,
) {
	errors.PanicOnError(
		c.PlaybackCommand(sessionIdentifier, command, seekPositionTicks),
	)
}
