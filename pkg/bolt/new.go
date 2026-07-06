package bolt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"go.etcd.io/bbolt"
	"time"
)

func New(path string) *Client {
	result, e := bbolt.Open(
		path,
		0600,
		&bbolt.Options{Timeout: 1 * time.Second},
	)
	errors.PanicOnError(e)

	return &Client{client: result}
}
