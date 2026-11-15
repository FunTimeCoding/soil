package permission_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/tool_call"
	"time"
)

func Call(
	name string,
	command string,
	source string,
	reason string,
	ago time.Duration,
) *tool_call.Call {
	result := tool_call.New(name, "", Now().Add(-ago).Format(time.RFC3339Nano))
	result.Detail = command
	result.PermissionSource = source
	result.PermissionReason = reason
	result.PermissionDecision = "accept"

	if source == "user_reject" {
		result.PermissionDecision = "reject"
	}

	return result
}
