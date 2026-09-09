package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/tree"
)

func (c *Client) MustTree(project int64) []*tree.Node {
	result, e := c.Tree(project, "", "", false, 0)
	errors.PanicOnError(e)

	return result
}
