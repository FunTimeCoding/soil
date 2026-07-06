package jira

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustTransition(
	key string,
	transitionIdentifier string,
) {
	errors.PanicOnError(c.Transition(key, transitionIdentifier))
}
