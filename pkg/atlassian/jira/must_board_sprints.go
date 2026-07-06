package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustBoardSprints(identifier int) []*jira.Sprint {
	result, e := c.BoardSprints(identifier)
	errors.PanicOnError(e)

	return result
}
