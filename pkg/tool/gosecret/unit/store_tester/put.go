package store_tester

import (
	"maps"
	"slices"
)

func (s *Tester) Put(
	path string,
	values map[string]string,
) ([]string, error) {
	merged := maps.Clone(s.existing)
	maps.Copy(merged, values)
	s.stored[path] = merged

	return slices.Sorted(maps.Keys(merged)), nil
}
