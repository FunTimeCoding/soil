package github

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/github/code"
)

func (c *Client) MustSearchCode(
	query string,
	a ...any,
) []*code.Code {
	result, e := c.SearchCode(query, a...)
	errors.PanicOnError(e)

	return result
}
