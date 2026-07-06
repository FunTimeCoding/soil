package store

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"go.etcd.io/bbolt"
	"time"
)

func (s *Store) Prune(cutoff time.Time) (int, error) {
	var result int
	e := s.client.Update(
		func(t *bbolt.Tx) error {
			b := s.client.Bucket(t, Bucket)

			if b == nil {
				return nil
			}

			var keys [][]byte
			c := b.Cursor()

			for k, v := c.First(); k != nil; k, v = c.Next() {
				var r Record
				notation.MustDecodeBytes(v, &r, false)

				if r.Start.Before(cutoff) {
					keys = append(keys, k)
				}
			}

			for _, k := range keys {
				if f := b.Delete(k); f != nil {
					return f
				}
			}

			result = len(keys)

			return nil
		},
	)

	return result, e
}
