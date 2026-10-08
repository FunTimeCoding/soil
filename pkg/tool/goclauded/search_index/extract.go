package search_index

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/entry"
	"slices"
)

func extract(
	b []byte,
	session string,
	turn string,
) ([]*entry.Entry, string) {
	var l notation.Line

	if json.Unmarshal(b, &l) != nil || l.Message == nil {
		return nil, turn
	}

	if l.Type != "user" && l.Type != "assistant" {
		return nil, turn
	}

	var m notation.Message

	if json.Unmarshal(l.Message, &m) != nil {
		return nil, turn
	}

	var result []*entry.Entry

	for _, c := range contentBlocks(m.Content) {
		switch c.Type {
		case "text":
			if c.Text == "" || l.Meta || claude.IsSystemNoise(c.Text) {
				continue
			}

			if l.Type == "user" {
				turn = l.Identifier
			}

			result = append(
				result,
				entry.New(
					l.Identifier,
					session,
					turn,
					l.Type,
					constant.BlockMessage,
					l.Timestamp,
					c.Text,
				),
			)
		case "tool_use":
			kind := constant.BlockCall

			if slices.Contains(constant.EditTools, c.Name) {
				kind = constant.BlockEdit
			}

			result = append(
				result,
				entry.New(
					l.Identifier,
					session,
					turn,
					l.Type,
					kind,
					l.Timestamp,
					inputText(c.Name, c.Input),
				),
			)
		}
	}

	return result, turn
}
