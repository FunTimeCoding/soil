package board

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/board/layout"

func (b *Board) Entries() []*layout.Entry {
	var result []*layout.Entry

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
