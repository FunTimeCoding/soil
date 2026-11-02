package guard

import "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"

func NewDecision(
	question string,
	hint string,
) *Decision {
	result := &SpecificOutput{
		HookEventName:     constant.PreToolUseEvent,
		AdditionalContext: hint,
	}

	if question != "" {
		result.PermissionDecision = constant.AskDecision
		result.PermissionDecisionReason = question
	}

	return &Decision{HookSpecificOutput: result}
}
