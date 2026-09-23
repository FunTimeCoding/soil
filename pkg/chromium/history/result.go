package history

import "github.com/funtimecoding/soil/pkg/chromium/history/entry"

type Result struct {
	CurrentIndex int64
	Entries      []*entry.Entry
}
