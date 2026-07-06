package github

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteRun(
	owner string,
	name string,
	run int64,
) {
	errors.PanicOnError(c.DeleteRun(owner, name, run))
}
