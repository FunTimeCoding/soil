package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustTransitions(key string) []jira.Transition {
	result, e := c.Transitions(key)
	errors.PanicOnError(e)

	return result
}
