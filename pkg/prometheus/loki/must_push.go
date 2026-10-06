package loki

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustPush(
	labels map[string]string,
	lines ...string,
) {
	errors.PanicOnError(c.Push(labels, lines...))
}
