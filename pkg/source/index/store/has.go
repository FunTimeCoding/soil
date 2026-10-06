package store

import "os"

func (s *Store) Has(
	kind string,
	key string,
) bool {
	_, e := os.Stat(s.path(kind, key))

	return e == nil
}
