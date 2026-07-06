package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSpaceByName(name string) *space.Space {
	result, e := c.SpaceByName(name)
	errors.PanicOnError(e)

	return result
}
