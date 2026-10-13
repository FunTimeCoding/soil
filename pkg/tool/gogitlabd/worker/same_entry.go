package worker

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"

func sameEntry(
	a *board_entry.Entry,
	b *board_entry.Entry,
) bool {
	return *a == *b
}
