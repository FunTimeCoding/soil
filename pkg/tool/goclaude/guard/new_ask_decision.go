package guard

import "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"

func NewAskDecision(reason string) *Decision {
	return &Decision{
		HookSpecificOutput: &SpecificOutput{
			HookEventName:            constant.PreToolUseEvent,
			PermissionDecision:       constant.AskDecision,
			PermissionDecisionReason: reason,
		},
	}
}
