package claude

import (
	"bufio"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"path/filepath"
)

func (c *Client) ToolCalls(sessionIdentifier string) []tool_call.Call {
	path := filepath.Join(
		c.base,
		join.Empty(sessionIdentifier, constant.NotationLogExtension),
	)
	f, e := os.Open(path)

	if e != nil {
		return nil
	}

	defer errors.PanicClose(f)
	var result []tool_call.Call
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	for scanner.Scan() {
		var line notation.Line

		if json.Unmarshal(scanner.Bytes(), &line) != nil {
			continue
		}

		if line.Type != "assistant" {
			continue
		}

		if line.Message == nil {
			continue
		}

		var m notation.Message

		if json.Unmarshal(line.Message, &m) != nil {
			continue
		}

		var blocks []json.RawMessage

		if json.Unmarshal(m.Content, &blocks) != nil {
			continue
		}

		for _, raw := range blocks {
			var b notation.ContentBlock

			if json.Unmarshal(raw, &b) != nil {
				continue
			}

			if b.Type != "tool_use" {
				continue
			}

			var detail string

			if b.Name == "Bash" {
				var input notation.BashInput

				if json.Unmarshal(b.Input, &input) == nil {
					detail = input.Command
				}
			}

			result = append(
				result,
				tool_call.Call{
					Name:       b.Name,
					Identifier: b.Identifier,
					Timestamp:  line.Timestamp,
					Detail:     detail,
				},
			)
		}
	}

	return result
}
