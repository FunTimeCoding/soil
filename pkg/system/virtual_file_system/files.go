package virtual_file_system

import (
	"slices"
	"sort"
)

func (s *System) Files() []string {
	if s.sorted == nil {
		s.sorted = make([]string, 0, len(s.files))

		for path := range s.files {
			s.sorted = append(s.sorted, path)
		}

		sort.Strings(s.sorted)
	}

	return slices.Clone(s.sorted)
}
