package jellyfin

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustPlay(
	sessionIdentifier string,
	itemIDs []string,
	playCommand string,
	startPositionTicks int64,
) {
	errors.PanicOnError(
		c.Play(sessionIdentifier, itemIDs, playCommand, startPositionTicks),
	)
}
