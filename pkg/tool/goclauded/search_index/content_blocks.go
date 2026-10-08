package search_index

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
)

func contentBlocks(raw json.RawMessage) []notation.ContentBlock {
	var text string

	if json.Unmarshal(raw, &text) == nil {
		return []notation.ContentBlock{{Type: "text", Text: text}}
	}

	var result []notation.ContentBlock

	if json.Unmarshal(raw, &result) != nil {
		return nil
	}

	return result
}
