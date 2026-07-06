package claude

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/content_block"
	"strings"
)

func ExtractText(raw json.RawMessage) string {
	var s string

	if json.Unmarshal(raw, &s) == nil {
		return s
	}

	var blocks []content_block.Block

	if json.Unmarshal(raw, &blocks) == nil {
		var parts []string

		for _, b := range blocks {
			if b.Type == "text" && b.Text != "" {
				parts = append(parts, b.Text)
			}
		}

		return strings.Join(parts, "\n\n")
	}

	return ""
}
