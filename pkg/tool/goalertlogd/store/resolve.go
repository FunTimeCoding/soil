package store

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"go.etcd.io/bbolt"
	"time"
)

func (s *Store) Resolve(key string) error {
	return s.client.Update(
		func(t *bbolt.Tx) error {
			b := s.client.Bucket(t, Bucket)
			v := b.Get([]byte(key))

			if v == nil {
				return nil
			}

			var r Record
			notation.MustDecodeBytes(v, &r, false)
			r.End = new(time.Now())

			return s.client.PutBytes(b, key, notation.Marshal(r))
		},
	)
}
