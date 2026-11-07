package claude

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/notation"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"github.com/funtimecoding/soil/pkg/generative/constant"
)

func recordDecision(
	pending map[string]*tool_call.Call,
	line *notation.Line,
) {
	if line.Permission == nil {
		return
	}

	var m notation.Message

	if json.Unmarshal(line.Message, &m) != nil {
		return
	}

	for _, b := range Blocks(m.Content) {
		if b.Type != constant.ClaudeToolResultBlock {
			continue
		}

		call, okay := pending[b.ToolUseIdentifier]

		if !okay {
			continue
		}

		delete(pending, b.ToolUseIdentifier)
		call.PermissionDecision = line.Permission.Decision
		call.PermissionSource = line.Permission.Source
		call.PermissionReason = line.Permission.ReasonCode

		if call.PermissionReason == "" {
			call.PermissionReason = line.Permission.ReasonType
		}
	}
}
