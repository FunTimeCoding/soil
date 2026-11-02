package unit

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/report"
)

func reportLines(r *report.Report) []string {
	var result []string

	for _, e := range r.Entries {
		result = append(result, join.Space(e.Class, e.Key()))
	}

	return result
}
