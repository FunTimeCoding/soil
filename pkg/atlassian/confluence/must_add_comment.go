package confluence

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustAddComment(
	pageIdentifier string,
	body string,
) {
	errors.PanicOnError(c.AddComment(pageIdentifier, body))
}
