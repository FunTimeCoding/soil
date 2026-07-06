package bolt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"go.etcd.io/bbolt"
)

func (c *Client) MustCreateBucket(
	t *bbolt.Tx,
	name string,
) *bbolt.Bucket {
	result, e := c.CreateBucket(t, name)
	errors.PanicOnError(e)

	return result
}
