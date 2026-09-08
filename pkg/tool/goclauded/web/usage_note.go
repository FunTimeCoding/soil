package web

import (
	"fmt"
	library "github.com/funtimecoding/soil/pkg/time"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strings"
	"time"
)

func (s *Server) usageNote() gomponents.Node {
	var items []string

	if result := s.service.Usage(); result != nil {
		items = append(
			items,
			fmt.Sprintf("Updated %s", library.FormatCompact(result.LastUpdated)),
		)
	}

	c := s.service.UtilizationCredential()

	if c != nil && c.Expired(time.Now()) {
		items = append(
			items,
			fmt.Sprintf(
				"credential expired %s, start a Claude Code session to refresh",
				library.FormatCompact(c.ExpiresAt),
			),
		)
	}

	if len(items) == 0 {
		return nil
	}

	return html.P(html.Em(gomponents.Text(strings.Join(items, " - "))))
}
