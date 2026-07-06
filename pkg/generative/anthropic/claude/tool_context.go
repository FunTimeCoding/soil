package claude

import (
	"bufio"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
	"strings"
)

func (c *Client) ToolContext(
	sessionIdentifier string,
	toolFilter string,
	surroundCount int,
) []ToolContextResult {
	path := filepath.Join(
		c.base,
		join.Empty(sessionIdentifier, constant.NotationLogExtension),
	)
	f, e := os.Open(path)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(f)
	type indexedMessage struct {
		Message  message.Message
		ToolUses []string
	}
	var messages []indexedMessage
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	for scanner.Scan() {
		var line notation.Line

		if json.Unmarshal(scanner.Bytes(), &line) != nil {
			continue
		}

		if line.Type != "user" && line.Type != "assistant" {
			continue
		}

		if line.Message == nil {
			continue
		}

		var m notation.Message

		if json.Unmarshal(line.Message, &m) != nil {
			continue
		}

		text := ExtractText(m.Content)

		if text == "" && line.Type == "user" {
			continue
		}

		im := indexedMessage{
			Message: message.Message{
				Role:      m.Role,
				Text:      text,
				Timestamp: line.Timestamp,
				IsMeta:    line.Meta || IsSystemNoise(text),
			},
		}

		if line.Type == "assistant" {
			var blocks []json.RawMessage

			if json.Unmarshal(m.Content, &blocks) == nil {
				for _, raw := range blocks {
					var b notation.ContentBlock

					if json.Unmarshal(raw, &b) != nil {
						continue
					}

					if b.Type != "tool_use" {
						continue
					}

					if strings.Contains(b.Name, toolFilter) {
						im.ToolUses = append(
							im.ToolUses,
							b.Name,
						)
					}
				}
			}
		}

		messages = append(messages, im)
	}

	var results []ToolContextResult

	for i, im := range messages {
		if len(im.ToolUses) == 0 {
			continue
		}

		for _, toolName := range im.ToolUses {
			r := ToolContextResult{
				ToolName: toolName,
			}
			start := i - surroundCount

			if start < 0 {
				start = 0
			}

			for j := start; j < i; j++ {
				if messages[j].Message.IsMeta {
					continue
				}

				r.Before = append(
					r.Before,
					messages[j].Message,
				)
			}

			end := i + surroundCount + 1

			if end > len(messages) {
				end = len(messages)
			}

			for j := i + 1; j < end; j++ {
				if messages[j].Message.IsMeta {
					continue
				}

				r.After = append(
					r.After,
					messages[j].Message,
				)
			}

			results = append(results, r)
		}
	}

	return results
}
