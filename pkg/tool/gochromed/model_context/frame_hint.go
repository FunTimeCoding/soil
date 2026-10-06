package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Server) frameHint(identifier string) string {
	tabs, e := s.client.Tabs()

	if e != nil {
		return ""
	}

	var lines []string

	for _, t := range tabs {
		if t.Type != constant.IframeTabType ||
			t.ParentIdentifier != identifier {
			continue
		}

		lines = append(lines, fmt.Sprintf("  %s %s", t.Identifier, t.Locator))
	}

	if len(lines) == 0 {
		return ""
	}

	return fmt.Sprintf(
		"\n\nThis tab contains cross-origin iframe targets. Their content renders as bare Iframe nodes above. Snapshot, evaluate, click, and read_body reach inside them when targeted by tab_id or url substring:\n%s",
		join.NewLine(lines),
	)
}
