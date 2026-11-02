package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"slices"
	"strings"
)

func (s *Service) Entries() []*board_entry.Entry {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	var result []*board_entry.Entry

	for _, i := range s.inventory.Instances {
		result = append(result, s.entries[i.Name]...)
	}

	slices.SortFunc(
		result,
		func(
			a *board_entry.Entry,
			b *board_entry.Entry,
		) int {
			if c := b.Updated.Compare(a.Updated); c != 0 {
				return c
			}

			if c := strings.Compare(a.Instance, b.Instance); c != 0 {
				return c
			}

			if c := strings.Compare(a.Project, b.Project); c != 0 {
				return c
			}

			return strings.Compare(a.Reference, b.Reference)
		},
	)

	return result
}
