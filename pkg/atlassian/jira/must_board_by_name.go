package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustBoardByName(name string) *jira.Board {
	result, e := c.BoardByName(name)
	errors.PanicOnError(e)

	return result
}
