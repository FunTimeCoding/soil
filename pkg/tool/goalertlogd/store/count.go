package store

import (
	"github.com/funtimecoding/soil/pkg/bolt"
	"go.etcd.io/bbolt"
)

func (s *Store) Count() (int, error) {
	var result int
	e := s.client.View(
		func(t *bbolt.Tx) error {
			b := s.client.Bucket(t, Bucket)

			if b == nil {
				return nil
			}

			return bolt.For(
				b,
				func(
					_ string,
					_ []byte,
				) {
					result++
				},
			)
		},
	)

	return result, e
}
