package claude

import (
	"bufio"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/peek"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
	"sort"
)

func (c *Client) Peek(sessionIdentifier string) *peek.Peek {
	path := filepath.Join(
		c.base,
		join.Empty(sessionIdentifier, constant.NotationLogExtension),
	)
	f, e := os.Open(path)

	if e != nil {
		return peek.Stub()
	}

	defer errors.PanicClose(f)
	result := peek.New()
	s := bufio.NewScanner(f)
	s.Buffer(make([]byte, 1024*1024), 1024*1024)
	toolCounts := map[string]int{}
	var lastAssistantText string

	for s.Scan() {
		result.LineCount++
		var line notation.Line

		if json.Unmarshal(s.Bytes(), &line) != nil {
			continue
		}

		if line.Message == nil {
			continue
		}

		var m notation.Message

		if json.Unmarshal(line.Message, &m) != nil {
			continue
		}

		if line.Type == "assistant" {
			var blocks []json.RawMessage

			if json.Unmarshal(m.Content, &blocks) == nil {
				for _, raw := range blocks {
					var b notation.ContentBlock

					if json.Unmarshal(raw, &b) != nil {
						continue
					}

					if b.Type == "tool_use" {
						toolCounts[b.Name]++
						result.TotalToolCalls++
					}
				}
			}

			text := ExtractText(m.Content)

			if text != "" && !IsSystemNoise(text) {
				clean := CleanContent(text)

				if len(clean) > 200 {
					clean = clean[:200]
				}

				lastAssistantText = clean
			}

			continue
		}

		if line.Type != "user" || line.Meta {
			continue
		}

		text := ExtractText(m.Content)

		if IsSystemNoise(text) {
			continue
		}

		clean := CleanContent(text)

		if len(clean) <= 20 {
			continue
		}

		if len(clean) > 100 {
			clean = clean[:100]
		}

		if last := len(result.Entries) - 1; last >= 0 {
			result.Entries[last].AssistantContext = lastAssistantText
		}

		result.Entries = append(
			result.Entries,
			peek.Entry{UserText: clean},
		)
		result.UserMessageCount++
		lastAssistantText = ""
	}

	if last := len(result.Entries) - 1; last >= 0 && lastAssistantText != "" {
		result.Entries[last].AssistantContext = lastAssistantText
	}

	var sorted []peek.ToolCount

	for name, count := range toolCounts {
		sorted = append(
			sorted,
			peek.ToolCount{
				Name:  name,
				Count: count,
			},
		)
	}

	sort.Slice(
		sorted,
		func(i, j int) bool {
			return sorted[i].Count > sorted[j].Count
		},
	)
	result.ToolCounts = sorted

	return result
}
