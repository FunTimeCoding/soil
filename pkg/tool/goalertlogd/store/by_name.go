package store

import (
	"github.com/funtimecoding/soil/pkg/bolt"
	"github.com/funtimecoding/soil/pkg/notation"
	"go.etcd.io/bbolt"
	"sort"
)

func (s *Store) ByName(name string) ([]Record, error) {
	var result []Record
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
					v []byte,
				) {
					var r Record
					notation.MustDecodeBytes(v, &r, false)

					if r.Name == name {
						result = append(result, r)
					}
				},
			)
		},
	)

	if e != nil {
		return nil, e
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			return result[i].Start.After(result[j].Start)
		},
	)

	return result, nil
}
