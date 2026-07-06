package store

import (
	"github.com/funtimecoding/soil/pkg/bolt"
	"go.etcd.io/bbolt"
)

func New(path string) *Store {
	c := bolt.New(path)
	c.MustUpdate(
		func(t *bbolt.Tx) error {
			_, e := c.CreateBucket(t, Bucket)

			return e
		},
	)

	return &Store{client: c}
}
