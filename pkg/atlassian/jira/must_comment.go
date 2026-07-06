package jira

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustComment(
	key string,
	body string,
) {
	errors.PanicOnError(c.Comment(key, body))
}
