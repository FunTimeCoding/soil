package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSpaces() []*space.Space {
	result, e := c.Spaces()
	errors.PanicOnError(e)

	return result
}
