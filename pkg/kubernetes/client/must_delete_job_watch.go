package client

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteJobWatch(
	namespace string,
	name string,
) {
	errors.PanicOnError(c.DeleteJobWatch(namespace, name))
}
