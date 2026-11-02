package report

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/module_range"
)

type Report struct {
	Entries      []*entry.Entry
	FromUnparsed []string
	ToUnparsed   []string
	Modules      []*module_range.Range
}
