package mock_client

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustPlay(
	sessionIdentifier string,
	itemIdentifiers []string,
	playCommand string,
	startPositionTicks int64,
) {
	errors.PanicOnError(
		c.Play(
			sessionIdentifier,
			itemIdentifiers,
			playCommand,
			startPositionTicks,
		),
	)
}
