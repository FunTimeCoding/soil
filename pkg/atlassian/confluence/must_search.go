package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSearch(
	query string,
	a ...any,
) []*search_result.Result {
	result, e := c.Search(query, a...)
	errors.PanicOnError(e)

	return result
}
