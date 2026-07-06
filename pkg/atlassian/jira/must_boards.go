package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustBoards() []*jira.Board {
	result, e := c.Boards()
	errors.PanicOnError(e)

	return result
}
