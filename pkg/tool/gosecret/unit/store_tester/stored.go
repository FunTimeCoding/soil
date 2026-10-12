package store_tester

import "maps"

func (s *Tester) Stored(path string) map[string]string {
	return maps.Clone(s.stored[path])
}
