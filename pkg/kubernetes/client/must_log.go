package client

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustLog(
	namespace string,
	pod string,
	container string,
) string {
	result, e := c.Log(namespace, pod, container)
	errors.PanicOnError(e)

	return result
}
