package board

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board/layout/entry"

func (b *Board) Entries() []*entry.Entry {
	var result []*entry.Entry

	for _, c := range b.Top {
		for _, s := range c.Sections {
			result = append(result, s.Entries...)
		}
	}

	for _, s := range b.Tail.Sections {
		result = append(result, s.Entries...)
	}

	return result
}
