package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/board_limit"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustBoardLimits(identifier int) []*board_limit.Limit {
	result, e := c.BoardLimits(identifier)
	errors.PanicOnError(e)

	return result
}
