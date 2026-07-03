package model_context

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chromium/constant"
	"strings"
)

// frameHint lists the cross-origin iframe targets rendered inside the
// given tab. Their content is invisible to the tab's accessibility
// tree because each frame is a separate target; appending this hint
// to snapshot output tells the model how to reach them.
func (s *Server) frameHint(identifier string) string {
	var lines []string

	for _, t := range s.client.Tabs() {
		if t.Type != constant.IframeTabType ||
			t.ParentIdentifier != identifier {
			continue
		}

		lines = append(
			lines,
			fmt.Sprintf("  %s %s", t.Identifier, t.Locator),
		)
	}

	if len(lines) == 0 {
		return ""
	}

	return fmt.Sprintf(
		"\n\nThis tab contains cross-origin iframe targets. Their content renders as bare Iframe nodes above. Snapshot, evaluate, click, and read_body reach inside them when targeted by tab_id or url substring:\n%s",
		strings.Join(lines, "\n"),
	)
}
