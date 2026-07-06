package bolt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"go.etcd.io/bbolt"
)

func (c *Client) MustPutBytes(
	b *bbolt.Bucket,
	k string,
	v []byte,
) {
	errors.PanicOnError(c.PutBytes(b, k, v))
}
