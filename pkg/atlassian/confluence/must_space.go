package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSpace(name string) *space.Space {
	result, e := c.Space(name)
	errors.PanicOnError(e)

	return result
}
