package claude

import (
	"bufio"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/message"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/generative/types/tool_context_result"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
	"strings"
)

func (c *Client) ToolContext(
	sessionIdentifier string,
	toolFilter string,
	surroundCount int,
) []tool_context_result.Result {
	path := filepath.Join(
		c.base,
		join.Empty(sessionIdentifier, constant.NotationLogExtension),
	)
	f, e := os.Open(path)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(f)
	var messages []message.Message
	var toolUses [][]string
	scanner := bufio.NewScanner(f)
	scanner.Buffer(nil, constant.NotationScanBuffer)

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

		var uses []string

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
						uses = append(uses, b.Name)
					}
				}
			}
		}

		messages = append(
			messages,
			message.Message{
				Role:      m.Role,
				Text:      text,
				Timestamp: line.Timestamp,
				IsMeta:    line.Meta || IsSystemNoise(text),
			},
		)
		toolUses = append(toolUses, uses)
	}

	var results []tool_context_result.Result

	for i, uses := range toolUses {
		if len(uses) == 0 {
			continue
		}

		for _, toolName := range uses {
			r := tool_context_result.Result{ToolName: toolName}
			start := i - surroundCount

			if start < 0 {
				start = 0
			}

			for j := start; j < i; j++ {
				if messages[j].IsMeta {
					continue
				}

				r.Before = append(r.Before, messages[j])
			}

			end := i + surroundCount + 1

			if end > len(messages) {
				end = len(messages)
			}

			for j := i + 1; j < end; j++ {
				if messages[j].IsMeta {
					continue
				}

				r.After = append(r.After, messages[j])
			}

			results = append(results, r)
		}
	}

	return results
}
