package worker

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"slices"
	"strings"
)

func (w *Worker) Entries() []*board_entry.Entry {
	w.mutex.RLock()
	defer w.mutex.RUnlock()
	var result []*board_entry.Entry

	for _, i := range w.service.Inventory().Instances {
		result = append(result, w.entries[i.Name]...)
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
