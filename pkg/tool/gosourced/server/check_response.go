package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gocheck/types/check_result"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/generated/server"
)

func checkResponse(c *check_result.Result) server.CheckResponse {
	sections := make([]server.CheckSection, 0, len(c.Sections))

	for _, r := range c.Sections {
		sections = append(sections, checkSection(r))
	}

	changes := make([]server.CheckChange, 0, len(c.Changes))

	for _, h := range c.Changes {
		changes = append(
			changes,
			server.CheckChange{
				Path:        h.Path,
				Destination: h.Destination,
				Original:    h.Original,
				Current:     h.Current,
			},
		)
	}

	return server.CheckResponse{
		Root:         c.Root,
		Announcement: c.Announcement,
		Unchanged:    c.Unchanged,
		Sections:     sections,
		Changes:      changes,
		Checked:      listed(c.Checked),
		Stale:        c.Stale,
	}
}
