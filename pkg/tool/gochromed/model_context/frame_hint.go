package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func frameHint(frames []*tab.Tab) string {
	if len(frames) == 0 {
		return ""
	}

	var lines []string

	for _, t := range frames {
		lines = append(lines, fmt.Sprintf("  %s %s", t.Identifier, t.Locator))
	}

	return fmt.Sprintf(
		"\n\nThis tab contains cross-origin iframe targets. Their content renders as bare Iframe nodes above. Snapshot, evaluate, click, and read_body reach inside them when targeted by tab_id or url substring:\n%s",
		join.NewLine(lines),
	)
}
