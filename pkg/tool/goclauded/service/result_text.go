package service

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/content_block"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func resultText(raw string) string {
	if raw == "" {
		return ""
	}

	var blocks []content_block.Block

	if json.Unmarshal([]byte(raw), &blocks) == nil {
		var parts []string

		for _, b := range blocks {
			if b.Text != "" {
				parts = append(parts, b.Text)
			}
		}

		return join.Empty(parts...)
	}

	var plain string

	if json.Unmarshal([]byte(raw), &plain) == nil {
		return plain
	}

	return raw
}
