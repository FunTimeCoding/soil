package firefox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/firefox/content"
)

func (c *Client) MustReadTab(
	identifier int,
	raw bool,
) *content.Content {
	result, e := c.ReadTab(identifier, raw)
	errors.PanicOnError(e)

	return result
}
