package store

import (
	"github.com/funtimecoding/soil/pkg/bolt"
	"github.com/funtimecoding/soil/pkg/notation"
	"go.etcd.io/bbolt"
	"sort"
)

func (s *Store) Unresolved() ([]UnresolvedRecord, error) {
	var result []UnresolvedRecord
	e := s.client.View(
		func(t *bbolt.Tx) error {
			b := s.client.Bucket(t, Bucket)

			if b == nil {
				return nil
			}

			return bolt.For(
				b,
				func(
					k string,
					v []byte,
				) {
					var r Record
					notation.MustDecodeBytes(v, &r, false)

					if r.End == nil {
						result = append(
							result,
							UnresolvedRecord{
								Fingerprint: r.Fingerprint,
								Key:         k,
							},
						)
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
			return result[i].Key < result[j].Key
		},
	)

	return result, nil
}
