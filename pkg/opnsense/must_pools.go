package opnsense

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/opnsense/pool"
)

func (c *Client) MustPools(phrase string) []*pool.Pool {
	result, e := c.Pools(phrase)
	errors.PanicOnError(e)

	return result
}
