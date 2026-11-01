package report

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"

type Report struct {
	Entries      []*entry.Entry
	FromUnparsed []string
	ToUnparsed   []string
}
