package store_tester

import "maps"

func (s *Tester) WithExisting(values map[string]string) *Tester {
	s.existing = maps.Clone(values)

	return s
}
