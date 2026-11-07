package claude

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func lineCalls(line *notation.Line) []*tool_call.Call {
	var m notation.Message

	if json.Unmarshal(line.Message, &m) != nil {
		return nil
	}

	var result []*tool_call.Call

	for _, b := range Blocks(m.Content) {
		if b.Type != constant.ClaudeToolUseBlock {
			continue
		}

		call := tool_call.New(b.Name, b.Identifier, line.Timestamp)
		call.Directory = line.WorkDirectory

		if b.Name == "Bash" {
			var input notation.BashInput

			if json.Unmarshal(b.Input, &input) == nil {
				call.Detail = input.Command
			}
		}

		result = append(result, call)
	}

	return result
}
