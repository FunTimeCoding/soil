package server

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
)

func checkSection(r *output.Results) server.CheckSection {
	entries := make([]server.CheckConcern, 0, len(r.Entries))

	for _, c := range r.Entries {
		entries = append(
			entries,
			server.CheckConcern{
				Key:      c.Key,
				Text:     c.Text,
				Path:     c.Path,
				Type:     c.Type,
				Line:     c.Line,
				LineText: c.LineText,
				Fixed:    c.Fixed,
				Planned:  c.Planned,
			},
		)
	}

	unchecked := make([]server.CheckUnchecked, 0, len(r.Unchecked))

	for _, u := range r.Unchecked {
		unchecked = append(
			unchecked,
			server.CheckUnchecked{
				Path:   u.Path,
				Line:   u.Line,
				Span:   u.Span,
				Reason: string(u.Reason),
			},
		)
	}

	return server.CheckSection{
		Header:    r.Header,
		Entries:   entries,
		Unchecked: unchecked,
	}
}
