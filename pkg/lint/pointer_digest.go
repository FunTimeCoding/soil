package lint

import (
	"github.com/funtimecoding/soil/pkg/lint/file_report"
	"github.com/funtimecoding/soil/pkg/lint/types/unchecked"
	"github.com/funtimecoding/soil/pkg/system/writer"
	"strings"
)

func pointerDigest(
	r *file_report.Report,
	spans []*unchecked.Unchecked,
) string {
	var b strings.Builder

	for _, c := range r.Concerns {
		writer.Print(
			&b,
			"%s|%s|%d|%s|%v\n",
			c.Key,
			c.Text,
			c.Line,
			c.LineText,
			c.Fixed,
		)
	}

	for _, u := range spans {
		writer.Print(&b, "unchecked|%d|%s|%s\n", u.Line, u.Span, u.Reason)
	}

	writer.Print(&b, "fixed|%d|%s\n", len(r.Fixed), r.Fixed)

	return b.String()
}
