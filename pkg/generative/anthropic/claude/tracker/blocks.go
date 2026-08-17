package tracker

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
)

func blocks(content json.RawMessage) []notation.ContentBlock {
	var raw []json.RawMessage

	if json.Unmarshal(content, &raw) != nil {
		return nil
	}

	var result []notation.ContentBlock

	for _, one := range raw {
		var b notation.ContentBlock

		if json.Unmarshal(one, &b) != nil {
			continue
		}

		result = append(result, b)
	}

	return result
}
