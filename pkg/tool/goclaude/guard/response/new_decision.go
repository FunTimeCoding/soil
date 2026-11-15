package response

import "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"

func NewDecision(
	askReason string,
	hint string,
) *Decision {
	result := &SpecificOutput{
		HookEventName:     constant.PreToolUseEvent,
		AdditionalContext: hint,
	}

	if askReason != "" {
		result.PermissionDecision = constant.AskDecision
		result.PermissionDecisionReason = askReason
	}

	return &Decision{HookSpecificOutput: result}
}
