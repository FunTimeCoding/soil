package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/link"
)

func (c *Client) MustCreateLink(
	l string,
	title string,
	listIdentifier int,
	tags []string,
) *link.Link {
	result, e := c.CreateLink(l, title, listIdentifier, tags)
	errors.PanicOnError(e)

	return result
}
