package bolt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"go.etcd.io/bbolt"
)

func (c *Client) MustView(f func(t *bbolt.Tx) error) {
	errors.PanicOnError(c.View(f))
}
