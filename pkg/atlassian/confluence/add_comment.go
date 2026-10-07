package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/request"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) AddComment(
	pageIdentifier string,
	body string,
) error {
	return c.basic.PostOldPath(
		"/content",
		notation.Encode(request.NewComment(pageIdentifier, body), false),
	)
}
