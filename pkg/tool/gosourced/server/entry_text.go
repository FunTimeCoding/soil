package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/report"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/entry"
)

func entryText(e *entry.Entry) *[]string {
	lines := report.EntryText(e)

	if len(lines) == 0 {
		return nil
	}

	return &lines
}
