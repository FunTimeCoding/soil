package store

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/notation"
	"go.etcd.io/bbolt"
)

func (s *Store) Save(r Record) (string, error) {
	key := fmt.Sprintf("%s|%d", r.Fingerprint, r.Start.UnixNano())
	e := s.client.Update(
		func(t *bbolt.Tx) error {
			return s.client.PutBytes(
				s.client.Bucket(t, Bucket),
				key,
				notation.Marshal(r),
			)
		},
	)

	return key, e
}
